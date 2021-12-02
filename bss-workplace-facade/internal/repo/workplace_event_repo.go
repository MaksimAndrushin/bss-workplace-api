package repo

import (
	"database/sql"
	"encoding/json"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-workplace-facade/internal/model"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"time"
)

type WorkplaceEventRepo interface {
	AddEvent(ctx context.Context, event model.WorkplaceEvent) error
	ListEvents(ctx context.Context, offset uint64, limit uint64) ([]model.WorkplaceEvent, uint64, error)
}

type workplaceEventRepo struct {
	db        *sqlx.DB
	batchSize uint
}

type DbWorkplaceEvent struct {
	ID          uint64             `db:"id"`
	WorkplaceId uint64             `db:"workplace_id"`
	Type        model.EventType    `db:"type"`
	Status      model.EventStatus  `db:"status"`
	Updated     time.Time          `db:"updated"`
	Entity      *DbWorkplaceEntity `db:"payload"`
}

type DbWorkplaceEntity struct {
	ID      uint64    `db:"id"`
	Name    string    `db:"name"`
	Size    uint32    `db:"size"`
	Removed bool      `db:"removed"`
	Created time.Time `db:"created"`
	Updated time.Time `db:"updated"`
}

const WORKPLACES_EVENTS_TAB = "workplaces_events_facade"
const WORKPLACES_EVENTS_PAYLOAD = "payload"
const WORKPLACES_EVENTS_TYPE = "type"
const WORKPLACES_EVENTS_WORKPLACE_ID = "workplace_id"
const WORKPLACES_EVENTS_STATUS = "status"
const WORKPLACES_EVENTS_UPDATED = "updated"

func NewWorkplaceEventRepo(db *sqlx.DB, batchSize uint) WorkplaceEventRepo {
	return &workplaceEventRepo{db: db, batchSize: batchSize}
}

func (r *workplaceEventRepo) AddEvent(ctx context.Context, event model.WorkplaceEvent) error {

	query := sq.Insert(WORKPLACES_EVENTS_TAB).PlaceholderFormat(sq.Dollar).
		Columns(WORKPLACES_EVENTS_WORKPLACE_ID, WORKPLACES_EVENTS_TYPE, WORKPLACES_EVENTS_STATUS, WORKPLACES_EVENTS_UPDATED, WORKPLACES_EVENTS_PAYLOAD).
		Values(event.Entity.ID, event.Type, event.Status, "NOW()", event.Entity).
		Suffix("RETURNING id")

	s, args, err := query.ToSql()
	if err != nil {
		return err
	}

	rows, err := r.db.QueryContext(ctx, s, args...)
	if err != nil {
		return err
	}

	defer rows.Close()

	var id uint64
	if rows.Next() {
		err = rows.Scan(&id)

		if err != nil {
			return err
		}

		log.Debug().Msgf("Created event id - %v", id)

		return nil
	} else {
		return sql.ErrNoRows
	}
}

func (r *workplaceEventRepo) ListEvents(ctx context.Context, offset uint64, limit uint64) ([]model.WorkplaceEvent, uint64, error) {
	query := sq.Select("*").PlaceholderFormat(sq.Dollar).
		From(WORKPLACES_EVENTS_TAB).
		Offset(offset).
		Limit(limit)

	s, args, err := query.ToSql()
	if err != nil {
		return nil, 0, err
	}

	rows, err := r.db.QueryContext(ctx, s, args...)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	events := make([]DbWorkplaceEvent, 0)
	err = sqlx.StructScan(rows, &events)

	if err != nil {
		return nil, 0, err
	}

	count, err := r.CountWorkplacesEvents(ctx)
	if err != nil {
		return nil, 0, err
	}

	resEvents := DBWorkplacesEventsToModelWorkplacesEvents(events)
	return resEvents, count, err
}

func (r *workplaceEventRepo) CountWorkplacesEvents(ctx context.Context) (uint64, error) {
	query := sq.Select("COUNT(*)").PlaceholderFormat(sq.Dollar).
		From(WORKPLACES_EVENTS_TAB)

	s, args, err := query.ToSql()
	if err != nil {
		return 0, err
	}

	var count uint64
	err = r.db.QueryRowxContext(ctx, s, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (w *DbWorkplaceEntity) Scan(src interface{}) error {
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

	*w = DbWorkplaceEntity{
		ID:   res.ID,
		Name: res.Name,
		Size: res.Size,
	}

	return nil
}

func DBWorkplaceEventToModelWorkplaceEvent(workplaceEvent DbWorkplaceEvent) model.WorkplaceEvent {
	return model.WorkplaceEvent{
		ID:     workplaceEvent.ID,
		Type:   workplaceEvent.Type,
		Status: workplaceEvent.Status,
		Entity: &model.Workplace{
			ID:   workplaceEvent.Entity.ID,
			Name: workplaceEvent.Entity.Name,
			Size: workplaceEvent.Entity.Size,
		},
	}
}

func DBWorkplacesEventsToModelWorkplacesEvents(workplaceEvents []DbWorkplaceEvent) []model.WorkplaceEvent {
	items := make([]model.WorkplaceEvent, 0)
	for _, workplace := range workplaceEvents {
		item := DBWorkplaceEventToModelWorkplaceEvent(workplace)
		items = append(items, item)
	}

	return items
}
