package repo

import (
	"database/sql"
	"encoding/json"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-workplace-api/internal/infra/metrics"
	"github.com/ozonmp/bss-workplace-api/internal/infra/tracer"
	"github.com/ozonmp/bss-workplace-api/internal/model"
	"golang.org/x/net/context"
	"time"
)

type WorkplaceEventRepo interface {
	Lock(ctx context.Context, recsCount uint64) ([]model.WorkplaceEvent, error)
	Unlock(ctx context.Context, eventIDs []uint64) error

	Add(ctx context.Context, event model.WorkplaceEvent, tx *sqlx.Tx) error
	Remove(ctx context.Context, eventIDs []uint64) error
}

type workplaceEventRepo struct {
	db        *sqlx.DB
	batchSize uint
}

type WorkplaceEntity struct {
	ID      uint64    `db:"id"`
	Name    string    `db:"name"`
	Size    uint32    `db:"size"`
	Removed bool      `db:"removed"`
	Created time.Time `db:"created"`
	Updated time.Time `db:"updated"`
}

type workplaceEventDb struct {
	ID          uint64            `db:"id"`
	WorkplaceId uint64            `db:"workplace_id"`
	Type        model.EventType   `db:"type"`
	Status      model.EventStatus `db:"status"`
	Entity      WorkplaceEntity   `db:"payload"`
	Updated     time.Time         `db:"updated"`
}

const WORKPLACES_EVENTS_TAB = "workplaces_events"
const WORKPLACES_EVENTS_ID = "id"
const WORKPLACES_EVENTS_PAYLOAD = "payload"
const WORKPLACES_EVENTS_TYPE = "type"
const WORKPLACES_EVENTS_WORKPLACE_ID = "workplace_id"
const WORKPLACES_EVENTS_STATUS = "status"
const WORKPLACES_EVENTS_UPDATED = "updated"

func NewWorkplaceEventRepo(db *sqlx.DB, batchSize uint) WorkplaceEventRepo {
	return &workplaceEventRepo{db: db, batchSize: batchSize}
}

func (r *workplaceEventRepo) Add(ctx context.Context, event model.WorkplaceEvent, tx *sqlx.Tx) error {

	span := tracer.CreateSpan(ctx, "Workplace Event Repo - Add")
	defer span.Close()

	query := sq.Insert(WORKPLACES_EVENTS_TAB).PlaceholderFormat(sq.Dollar).
		Columns(WORKPLACES_EVENTS_WORKPLACE_ID, WORKPLACES_EVENTS_TYPE, WORKPLACES_EVENTS_STATUS, WORKPLACES_EVENTS_UPDATED, WORKPLACES_EVENTS_PAYLOAD).
		Values(event.Entity.ID, event.Type, event.Status, "NOW()", event.Entity).
		Suffix("RETURNING id")

	s, args, err := query.ToSql()
	if err != nil {
		return err
	}

	rows, err := r.getQueryerContext(tx).QueryContext(ctx, s, args...)
	if err != nil {
		return err
	}

	defer rows.Close()

	if !rows.Next() {
		return sql.ErrNoRows
	}

	return nil
}

func (r *workplaceEventRepo) Remove(ctx context.Context, eventIDs []uint64) error {

	query := sq.Delete(WORKPLACES_EVENTS_TAB).PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{WORKPLACES_EVENTS_ID: eventIDs})

	s, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, s, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *workplaceEventRepo) Lock(ctx context.Context, recsCount uint64) ([]model.WorkplaceEvent, error) {

	subSelSql := sq.Select("we1.id").PlaceholderFormat(sq.Dollar).
		From("workplaces_events we1").
		Where(sq.Eq{"we1.status": model.Deferred}).
		OrderBy("we1.id ASC").
		Limit(recsCount).Prefix("we0.id IN(").Suffix(")")

	query := sq.Update("workplaces_events we0").PlaceholderFormat(sq.Dollar).
		Set(WORKPLACES_EVENTS_STATUS, model.Locked).
		Set(WORKPLACES_EVENTS_UPDATED, "NOW()").
		Where(subSelSql).
		Suffix("RETURNING we0.id, we0.workplace_id, weo.type, we0.status, we0.payload, we0.updated")

	s, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, s, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	workplacesEventsDb := make([]*workplaceEventDb, 0)
	err = sqlx.StructScan(rows, &workplacesEventsDb)
	if err != nil {
		return nil, err
	}

	workplacesEvents := convertToWorkplaceEvents(workplacesEventsDb)

	metrics.AddRetranslatorEvents(len(workplacesEvents))

	return workplacesEvents, nil
}

func (r *workplaceEventRepo) Unlock(ctx context.Context, eventIDs []uint64) error {

	query := sq.Update(WORKPLACES_EVENTS_TAB).PlaceholderFormat(sq.Dollar).
		Set(WORKPLACES_EVENTS_STATUS, model.Deferred).
		Set(WORKPLACES_EVENTS_UPDATED, "NOW()").
		Where(sq.Eq{WORKPLACES_EVENTS_ID: eventIDs})

	s, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, s, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *workplaceEventRepo) getExecerContext(tx *sqlx.Tx) sqlx.ExecerContext {
	if tx == nil {
		return r.db
	}
	return tx
}

func (r *workplaceEventRepo) getQueryerContext(tx *sqlx.Tx) sqlx.QueryerContext {
	if tx == nil {
		return r.db
	}
	return tx
}

func convertToWorkplaceEvents(workplacesEventsDb []*workplaceEventDb) []model.WorkplaceEvent {

	workplaceEvents := make([]model.WorkplaceEvent, 0)

	for _, wpEventDb := range workplacesEventsDb {
		workplace := model.Workplace{
			ID:      wpEventDb.Entity.ID,
			Name:    wpEventDb.Entity.Name,
			Size:    wpEventDb.Entity.Size,
			Removed: wpEventDb.Entity.Removed,
			Created: wpEventDb.Entity.Created,
			Updated: wpEventDb.Entity.Updated,
		}
		workplaceEvents = append(workplaceEvents, model.WorkplaceEvent{ID: wpEventDb.ID, Type: wpEventDb.Type, Status: wpEventDb.Status, Entity: &workplace})
	}

	return workplaceEvents
}

func (w *WorkplaceEntity) Scan(src interface{}) error {
	var source []byte
	switch src.(type) {
	case string:
		source = []byte(src.(string))
	case []byte:
		source = src.([]byte)
	default:
		return errors.New("incompatible type for workplace")
	}

	res := &model.Workplace{}

	err := json.Unmarshal(source, res)

	if err != nil {
		return err
	}

	w = &WorkplaceEntity{
		ID:   res.ID,
		Name: res.Name,
		Size: res.Size,
	}

	return nil
}
