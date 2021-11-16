package service

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-workplace-api/internal/database"
	"github.com/ozonmp/bss-workplace-api/internal/infra/logger"
	"github.com/ozonmp/bss-workplace-api/internal/infra/tracer"
	"github.com/ozonmp/bss-workplace-api/internal/model"
	"github.com/ozonmp/bss-workplace-api/internal/repo"
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

	span := tracer.CreateSpan(ctx, "Service CreateWorkplace")
	defer span.Close()

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
		logger.ErrorKV(ctx, "CreateWorkplace -- failed", "err", err)
		return 0, err
	}

	return workplaceId, nil
}

func (ws *workplaceService) DescribeWorkplace(ctx context.Context, workplaceID uint64) (*model.Workplace, error) {
	span := tracer.CreateSpan(ctx, "Service DescribeWorkplace")
	defer span.Close()

	workplace, err := ws.WorkplaceRepo.DescribeWorkplace(ctx, workplaceID)
	if err != nil {
		logger.ErrorKV(ctx, "DescribeWorkplace -- failed", "err", err)
		return nil, err
	}

	return workplace, nil
}

func (ws *workplaceService) ListWorkplaces(ctx context.Context, offset uint64, limit uint64) ([]model.Workplace, error) {
	span := tracer.CreateSpan(ctx, "Service ListWorkplaces")
	defer span.Close()

	workplaces, err := ws.WorkplaceRepo.ListWorkplaces(ctx, offset, limit)
	if err != nil {
		logger.ErrorKV(ctx, "ListWorkplaces -- failed", "err", err)
		return nil, err
	}

	return workplaces, nil

}

func (ws *workplaceService) RemoveWorkplace(ctx context.Context, workplaceID uint64) (bool, error) {
	span := tracer.CreateSpan(ctx, "Service RemoveWorkplace")
	defer span.Close()

	ok, err := ws.WorkplaceRepo.RemoveWorkplace(ctx, workplaceID)
	if err != nil {
		logger.ErrorKV(ctx, "RemoveWorkplace -- failed", "err", err)
		return false, err
	}

	return ok, nil
}
