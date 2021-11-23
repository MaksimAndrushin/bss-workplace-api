package repo

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-workplace-api/internal/infra/tracer"

	"github.com/ozonmp/bss-workplace-api/internal/model"
)

// WorkplaceRepo is DAO for workplace
type WorkplaceRepo interface {
	CreateWorkplace(ctx context.Context, name string, size uint32, tx *sqlx.Tx) (uint64, error)
	DescribeWorkplace(ctx context.Context, workplaceID uint64) (*model.Workplace, error)
	ListWorkplaces(ctx context.Context, offset uint64, limit uint64) ([]model.Workplace, error)
	RemoveWorkplace(ctx context.Context, workplaceID uint64) (bool, error)
	UpdateWorkplace(ctx context.Context, workplace model.Workplace, tx *sqlx.Tx) (bool, error)
}

type workplaceRepo struct {
	db        *sqlx.DB
	batchSize uint
}

const WORKPLACES_TAB = "workplaces"
const WORKPLACES_ID = "id"
const WORKPLACES_NAME = "name"
const WORKPLACES_SIZE = "size"
const WORKPLACES_REMOVED = "removed"
const WORKPLACES_CREATED = "created"
const WORKPLACES_UPDATED = "updated"

// NewWorkplaceRepo returns WorkplaceRepo interface
func NewWorkplaceRepo(db *sqlx.DB, batchSize uint) WorkplaceRepo {
	return &workplaceRepo{db: db, batchSize: batchSize}
}

func (r *workplaceRepo) CreateWorkplace(ctx context.Context, name string, size uint32, tx *sqlx.Tx) (uint64, error) {

	span := tracer.CreateSpan(ctx, "Workplace Repo CreateWorkplace")
	defer span.Close()

	query := sq.Insert(WORKPLACES_TAB).PlaceholderFormat(sq.Dollar).
		Columns(WORKPLACES_NAME, WORKPLACES_SIZE, WORKPLACES_REMOVED, WORKPLACES_CREATED, WORKPLACES_UPDATED).
		Values(name, size, false, "NOW()", "NOW()").
		Suffix("RETURNING id")

	s, args, err := query.ToSql()
	if err != nil {
		return 0, err
	}

	rows, err := r.getQueryerContext(tx).QueryContext(ctx, s, args...)
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	var id uint64
	if rows.Next() {
		err = rows.Scan(&id)

		if err != nil {
			return 0, err
		}

		return id, nil
	} else {
		return 0, sql.ErrNoRows
	}
}

func (r *workplaceRepo) DescribeWorkplace(ctx context.Context, workplaceID uint64) (*model.Workplace, error) {
	span := tracer.CreateSpan(ctx, "Workplace Repo DescribeWorkplace")
	defer span.Close()

	query := sq.Select(WORKPLACES_ID, WORKPLACES_NAME, WORKPLACES_SIZE).PlaceholderFormat(sq.Dollar).
		From(WORKPLACES_TAB).
		Where(sq.Eq{WORKPLACES_ID: workplaceID})

	s, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var workplace model.Workplace
	err = r.db.QueryRowxContext(ctx, s, args...).Scan(&workplace.ID, &workplace.Name, &workplace.Size)
	if err != nil {
		return nil, err
	}

	return &workplace, err
}

func (r *workplaceRepo) ListWorkplaces(ctx context.Context, offset uint64, limit uint64) ([]model.Workplace, error) {
	span := tracer.CreateSpan(ctx, "Workplace Repo ListWorkplaces")
	defer span.Close()

	query := sq.Select("*").PlaceholderFormat(sq.Dollar).
		From(WORKPLACES_TAB).
		Where(sq.Eq{WORKPLACES_REMOVED: false}).
		OrderBy("id ASC").
		Offset(offset).
		Limit(limit)

	s, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, s, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	workplaces := make([]model.Workplace, 0)
	err = sqlx.StructScan(rows, &workplaces)

	if err != nil {
		return nil, err
	}

	return workplaces, err
}

func (r *workplaceRepo) RemoveWorkplace(ctx context.Context, workplaceID uint64) (bool, error) {
	span := tracer.CreateSpan(ctx, "Workplace Repo RemoveWorkplace")
	defer span.Close()

	query := sq.Delete(WORKPLACES_TAB).PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{WORKPLACES_ID: workplaceID})

	s, args, err := query.ToSql()
	if err != nil {
		return false, err
	}

	res, err := r.db.ExecContext(ctx, s, args...)
	if err != nil {
		return false, err
	}

	deletedRowsCount, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	if deletedRowsCount == 0 {
		return false, nil
	}

	return true, nil
}

func (r *workplaceRepo) UpdateWorkplace(ctx context.Context, workplace model.Workplace, tx *sqlx.Tx) (bool, error) {
	span := tracer.CreateSpan(ctx, "Workplace Repo UpdateWorkplace")
	defer span.Close()

	query := sq.Update(WORKPLACES_TAB).PlaceholderFormat(sq.Dollar).
		Set(WORKPLACES_NAME, workplace.Name).
		Set(WORKPLACES_SIZE, workplace.Size).
		Set(WORKPLACES_UPDATED, "NOW()").
		Where(sq.Eq{WORKPLACES_ID: workplace.ID})

	s, args, err := query.ToSql()
	if err != nil {
		return false, err
	}

	res, err := r.getExecerContext(tx).ExecContext(ctx, s, args...)
	if err != nil {
		return false, err
	}

	updatedRowsCount, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	if updatedRowsCount == 0 {
		return false, nil
	}

	return true, nil
}

func (r *workplaceRepo) getExecerContext(tx *sqlx.Tx) sqlx.ExecerContext {
	if tx == nil {
		return r.db
	}
	return tx
}

func (r *workplaceRepo) getQueryerContext(tx *sqlx.Tx) sqlx.QueryerContext {
	if tx == nil {
		return r.db
	}
	return tx
}
