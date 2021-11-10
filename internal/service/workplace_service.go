package service

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-workplace-api/internal/database"
	"github.com/ozonmp/bss-workplace-api/internal/model"
	"github.com/ozonmp/bss-workplace-api/internal/repo"
	"github.com/rs/zerolog/log"
)

type WorkplaceService interface {
	CreateWorkplace(ctx context.Context, name string, size uint32) (uint64, error)
	DescribeWorkplace(ctx context.Context, workplaceID uint64) (*model.Workplace, error)
	ListWorkplaces(ctx context.Context, offset uint64, limit uint64) ([]model.Workplace, error)
	RemoveWorkplace(ctx context.Context, workplaceID uint64) (bool, error)
}

type workplaceService struct {
	WorkplaceRepo      repo.WorkplaceRepo
	WorkplaceEventRepo repo.WorkplaceEventRepo
	DbHolder           *sqlx.DB
}

func NewWorkplaceService(workplaceRepo repo.WorkplaceRepo, workplaceEventRepo repo.WorkplaceEventRepo, db *sqlx.DB) WorkplaceService {
	return &workplaceService{
		WorkplaceRepo:      workplaceRepo,
		WorkplaceEventRepo: workplaceEventRepo,
		DbHolder:           db,
	}
}

func (ws *workplaceService) CreateWorkplace(ctx context.Context, name string, size uint32) (uint64, error) {

	var workplaceId uint64
	err := database.WithTx(ctx, ws.DbHolder, func(ctx context.Context, tx *sqlx.Tx) error {

		workplaceIdInternal, err := ws.WorkplaceRepo.CreateWorkplace(ctx, name, size, tx)
		if err != nil {
			return err
		}

		workplaceEntity := model.CreateEventFromWorkplace(model.Created, model.Deferred, model.Workplace{ID: workplaceIdInternal, Name: name, Size: size})

		err = ws.WorkplaceEventRepo.Add(ctx, *workplaceEntity, tx)
		if err != nil {
			return err
		}

		workplaceId = workplaceIdInternal
		return nil
	})

	if err != nil {
		return 0, err
	}

	return workplaceId, nil
}

func (ws *workplaceService) DescribeWorkplace(ctx context.Context, workplaceID uint64) (*model.Workplace, error) {
	workplace, err := ws.WorkplaceRepo.DescribeWorkplace(ctx, workplaceID, nil)
	if err != nil {
		log.Error().Err(err).Msg("DescribeWorkplaceV1 -- failed")
		return nil, err
	}

	return workplace, nil
}

func (ws *workplaceService) ListWorkplaces(ctx context.Context, offset uint64, limit uint64) ([]model.Workplace, error) {
	workplaces, err := ws.WorkplaceRepo.ListWorkplaces(ctx, offset, limit, nil)
	if err != nil {
		log.Error().Err(err).Msg("ListWorkplacesV1 -- failed")
		return nil, err
	}

	return workplaces, nil

}

func (ws *workplaceService) RemoveWorkplace(ctx context.Context, workplaceID uint64) (bool, error) {
	ok, err := ws.WorkplaceRepo.RemoveWorkplace(ctx, workplaceID, nil)
	if err != nil {
		log.Error().Err(err).Msg("DescribeWorkplaceV1 -- failed")
		return false, err
	}

	return ok, nil
}
