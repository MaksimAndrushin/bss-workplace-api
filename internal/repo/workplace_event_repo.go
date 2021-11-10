package repo

import (
	"database/sql"
	"encoding/json"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-workplace-api/internal/model"
	"golang.org/x/net/context"
	"time"
)

type WorkplaceEventRepo interface {
	Lock(ctx context.Context, recsCount uint64, tx *sqlx.Tx) ([]model.WorkplaceEvent, error)
	Unlock(ctx context.Context, eventIDs []uint64, tx *sqlx.Tx) error

	Add(ctx context.Context, event model.WorkplaceEvent, tx *sqlx.Tx) error
	Remove(ctx context.Context, eventIDs []uint64, tx *sqlx.Tx) error
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

func NewWorkplaceEventRepo(db *sqlx.DB, batchSize uint) WorkplaceEventRepo {
	return &workplaceEventRepo{db: db, batchSize: batchSize}
}

func (r *workplaceEventRepo) Add(ctx context.Context, event model.WorkplaceEvent, tx *sqlx.Tx) error {
	query := sq.Insert("workplaces_events").PlaceholderFormat(sq.Dollar).
		Columns("workplace_id", "type", "status", "updated", "payload").
		Values(event.Entity.ID, event.Type, event.Status, "NOW()", event.Entity).
		Suffix("RETURNING id")

	s, args, err := query.ToSql()
	if err != nil {
		return err
	}

	rows, err := r.getQueryerContext(tx).QueryContext(ctx, s, args...)
	defer rows.Close()

	if err != nil {
		return err
	}

	if !rows.Next() {
		return sql.ErrNoRows
	}

	return nil
}

func (r *workplaceEventRepo) Remove(ctx context.Context, eventIDs []uint64, tx *sqlx.Tx) error {
	query := sq.Delete("workplaces_events").PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": eventIDs})

	s, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.getExecerContext(tx).ExecContext(ctx, s, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *workplaceEventRepo) Lock(ctx context.Context, recsCount uint64, tx *sqlx.Tx) ([]model.WorkplaceEvent, error) {

	subSelSql := sq.Select("we1.id").PlaceholderFormat(sq.Dollar).
		From("workplaces_events we1").
		Where(sq.Eq{"we1.status": model.Deferred}).
		OrderBy("we1.id ASC").
		Limit(recsCount).Prefix("we0.id IN(").Suffix(")")

	query := sq.Update("workplaces_events we0").PlaceholderFormat(sq.Dollar).
		Set("status", model.Locked).
		Set("updated", "NOW()").
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
	return workplacesEvents, nil
}

func (r *workplaceEventRepo) Unlock(ctx context.Context, eventIDs []uint64, tx *sqlx.Tx) error {
	query := sq.Update("workplaces_events").PlaceholderFormat(sq.Dollar).
		Set("status", model.Deferred).
		Set("updated_at", "NOW()").
		Where(sq.Eq{"id": eventIDs})

	s, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.getExecerContext(tx).ExecContext(ctx, s, args...)
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
