package repo

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/ozonmp/bss-workplace-api/internal/model"
)

// WorkplaceRepo is DAO for workplace
type WorkplaceRepo interface {
	CreateWorkplace(ctx context.Context, name string, size uint32, tx *sqlx.Tx) (uint64, error)
	DescribeWorkplace(ctx context.Context, workplaceID uint64, tx *sqlx.Tx) (*model.Workplace, error)
	ListWorkplaces(ctx context.Context, offset uint64, limit uint64, tx *sqlx.Tx) ([]model.Workplace, error)
	RemoveWorkplace(ctx context.Context, workplaceID uint64, tx *sqlx.Tx) (bool, error)
}

type workplaceRepo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewWorkplaceRepo returns WorkplaceRepo interface
func NewWorkplaceRepo(db *sqlx.DB, batchSize uint) WorkplaceRepo {
	return &workplaceRepo{db: db, batchSize: batchSize}
}

func (r *workplaceRepo) CreateWorkplace(ctx context.Context, name string, size uint32, tx *sqlx.Tx) (uint64, error) {
	query := sq.Insert("workplaces").PlaceholderFormat(sq.Dollar).
		Columns("name", "size", "removed", "created", "updated").
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

func (r *workplaceRepo) DescribeWorkplace(ctx context.Context, workplaceID uint64, tx *sqlx.Tx) (*model.Workplace, error) {
	query := sq.Select("id, name, size").PlaceholderFormat(sq.Dollar).
		From("workplaces").
		Where(sq.Eq{"id": workplaceID})

	s, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var workplace []model.Workplace
	err = r.db.SelectContext(ctx, &workplace, s, args...)
	if err != nil {
		return nil, err
	}

	if len(workplace) == 0 {
		return nil, nil
	}

	return &workplace[0], err
}

func (r *workplaceRepo) ListWorkplaces(ctx context.Context, offset uint64, limit uint64, tx *sqlx.Tx) ([]model.Workplace, error) {
	query := sq.Select("*").PlaceholderFormat(sq.Dollar).
		From("workplaces").
		Where(sq.Eq{"removed": false}).
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

	workplaces := make([]model.Workplace, 0)
	err = sqlx.StructScan(rows, &workplaces)

	if err != nil {
		return nil, err
	}

	return workplaces, err
}

func (r *workplaceRepo) RemoveWorkplace(ctx context.Context, workplaceID uint64, tx *sqlx.Tx) (bool, error) {
	query := sq.Delete("workplaces").PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": workplaceID})

	s, args, err := query.ToSql()
	if err != nil {
		return false, err
	}

	_, err = r.getExecerContext(tx).ExecContext(ctx, s, args...)
	if err != nil {
		return false, err
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
