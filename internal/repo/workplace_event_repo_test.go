package repo

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-workplace-api/internal/model"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestLock(t *testing.T) {
	r, dbMock := setupEventRepo()

	rows := sqlmock.NewRows([]string{"id", "workplace_id", "type", "status", "payload", "updated"}).
		AddRow(1, 2, model.Created, model.Deferred, "{}", time.Now())

	dbMock.ExpectQuery(`UPDATE workplaces_events we0 SET status = $1, updated = $2 
                                   WHERE we0.id IN( 
                                      SELECT we1.id FROM workplaces_events we1 WHERE we1.status = $3 ORDER BY we1.id ASC LIMIT 10 ) 
                                   RETURNING we0.id, we0.workplace_id, weo.type, we0.status, we0.payload, we0.updated`).
		WithArgs(model.Locked, "NOW()", model.Deferred).
		WillReturnRows(rows)

	_, err := r.Lock(context.Background(), 10, nil)

	require.NoError(t, err)

}

func TestUnlock(t *testing.T) {
	r, dbMock := setupEventRepo()

	dbMock.ExpectExec(`UPDATE workplaces_events SET status = $1, updated_at = $2 WHERE id IN ($3)`).
		WithArgs(model.Deferred, "NOW()", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := r.Unlock(context.Background(), []uint64{1}, nil)

	require.NoError(t, err)
}

func TestRemove(t *testing.T) {
	r, dbMock := setupEventRepo()

	dbMock.ExpectExec(`DELETE FROM workplaces_events WHERE id IN ($1)`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := r.Remove(context.Background(), []uint64{1}, nil)

	require.NoError(t, err)
}

func setupEventRepo() (*workplaceEventRepo, sqlmock.Sqlmock) {
	mockDB, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	repo := workplaceEventRepo{
		db:        sqlxDB,
		batchSize: 10,
	}

	return &repo, mock
}
