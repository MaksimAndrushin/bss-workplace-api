package database

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-workplace-api/internal/infra/logger"
)

// NewPostgres returns DB
func NewPostgres(ctx context.Context, dsn string, driver string) (*sqlx.DB, error) {
	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		logger.ErrorKV(ctx,"failed to create database connection", "err", err)
		return nil, err
	}

	if err = db.PingContext(ctx); err != nil {
		logger.ErrorKV(ctx,"failed ping the database", "err", err)
		return nil, err
	}

	return db, nil
}
