package database

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"time"
)

// NewPostgres returns DB
func NewPostgres(ctx context.Context, dsn string, driver string, maxReconnectsCount int64, reconnectDelaySecs int64) (*sqlx.DB, error) {

	db, err := tryCreateDBConnection(ctx, dsn, driver)
	if err == nil {
		return db, nil
	}

	ticker := time.NewTicker(time.Duration(reconnectDelaySecs) * time.Second)
	defer ticker.Stop()

	for i := 1; i <= int(maxReconnectsCount); i++ {
		select {
		case <-ticker.C:
			db, err := tryCreateDBConnection(ctx, dsn, driver)
			if err == nil {
				return db, nil
			}

		case <-ctx.Done():
			return nil, errors.New("DB Connection create canceled")
		}
	}

	return nil, errors.New("Can't connect to DB")
}

func tryCreateDBConnection(ctx context.Context, dsn string, driver string) (*sqlx.DB, error) {
	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		log.Error().Msgf("failed to create database connection - %v", err)
		return nil, err
	}

	if err = db.PingContext(ctx); err != nil {
		log.Error().Msgf("failed ping the database - %v", err)
		return nil, err
	}

	return db, nil
}
