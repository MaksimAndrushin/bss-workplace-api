package repo

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCreateWorkplace(t *testing.T) {
	r, dbMock := setupWorkplaceRepo()

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	dbMock.ExpectQuery(`INSERT INTO workplaces (name,size,removed,created,updated) VALUES ($1,$2,$3,$4,$5) RETURNING id`).
		WithArgs("NAME 1", 10, false, "NOW()", "NOW()").
		WillReturnRows(rows)

	_, err := r.CreateWorkplace(context.Background(), "NAME 1", 10, nil)

	require.NoError(t, err)
}

func TestDescribeWorkplace(t *testing.T) {
	r, dbMock := setupWorkplaceRepo()

	rows := sqlmock.NewRows([]string{"id", "name", "size", "removed", "created", "updated"}).
		AddRow(1, "NAME 1", 10, false, time.Now(), time.Now())

	dbMock.ExpectQuery(`SELECT id, name, size FROM workplaces WHERE id = $1`).
		WithArgs(1).
		WillReturnRows(rows)

	_, err := r.DescribeWorkplace(context.Background(), 1, nil)

	require.NoError(t, err)
}

func TestListWorkplace(t *testing.T) {
	r, dbMock := setupWorkplaceRepo()

	rows := sqlmock.NewRows([]string{"id", "name", "size", "removed", "created", "updated"}).
		AddRow(1, "NAME 1", 10, false, time.Now(), time.Now()).
		AddRow(2, "NAME 2", 10, false, time.Now(), time.Now()).
		AddRow(3, "NAME 3", 10, false, time.Now(), time.Now())

	dbMock.ExpectQuery(`SELECT * FROM workplaces WHERE removed = $1 ORDER BY id ASC LIMIT 3 OFFSET 0`).
		WillReturnRows(rows)

	_, err := r.ListWorkplaces(context.Background(), 0, 3, nil)

	require.NoError(t, err)
}

func TestRemoveWorkplace(t *testing.T) {
	r, dbMock := setupWorkplaceRepo()

	dbMock.ExpectExec(`DELETE FROM workplaces WHERE id = $1`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err := r.RemoveWorkplace(context.Background(), 1, nil)

	require.NoError(t, err)
}

func setupWorkplaceRepo() (*workplaceRepo, sqlmock.Sqlmock) {
	mockDB, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	repo := workplaceRepo{
		db:        sqlxDB,
		batchSize: 10,
	}

	return &repo, mock
}
