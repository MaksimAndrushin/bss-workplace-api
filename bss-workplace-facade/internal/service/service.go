package service

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-workplace-facade/internal/model"
	"github.com/ozonmp/bss-workplace-facade/internal/repo"
	"github.com/rs/zerolog/log"
)

type WorkplaceEventsService interface {
	ListWorkplacesEvents(ctx context.Context, offset uint64, limit uint64) ([]model.WorkplaceEvent, uint64, error)
}

type workplaceEventsService struct {
	WorkplaceEventRepo repo.WorkplaceEventRepo
	DbHolder           *sqlx.DB
}

func NewWorkplaceEventsService(workplaceEventRepo repo.WorkplaceEventRepo, db *sqlx.DB) WorkplaceEventsService {
	return &workplaceEventsService{
		WorkplaceEventRepo: workplaceEventRepo,
		DbHolder:           db,
	}
}

func (ws *workplaceEventsService) ListWorkplacesEvents(ctx context.Context, offset uint64, limit uint64) ([]model.WorkplaceEvent, uint64, error) {
	events, total, err := ws.WorkplaceEventRepo.ListEvents(ctx, offset, limit)
	if err != nil {
		log.Error().Msgf("List Workplaces Events - failed. Err - %v", err)
		return nil, 0, err
	}

	return events, total, nil
}
