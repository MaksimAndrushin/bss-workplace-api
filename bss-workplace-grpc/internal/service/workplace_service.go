package service

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-workplace-api/internal/database"
	"github.com/ozonmp/bss-workplace-api/internal/infra/logger"
	"github.com/ozonmp/bss-workplace-api/internal/infra/redis"
	"github.com/ozonmp/bss-workplace-api/internal/infra/tracer"
	"github.com/ozonmp/bss-workplace-api/internal/model"
	"github.com/ozonmp/bss-workplace-api/internal/repo"
)

type WorkplaceService interface {
	CreateWorkplace(ctx context.Context, name string, size uint32) (uint64, error)
	DescribeWorkplace(ctx context.Context, workplaceID uint64) (*model.Workplace, error)
	ListWorkplaces(ctx context.Context, offset uint64, limit uint64) ([]model.Workplace, uint64, error)
	RemoveWorkplace(ctx context.Context, workplaceID uint64) (bool, error)
	UpdateWorkplace(ctx context.Context, workplace model.Workplace) (bool, error)
}

type workplaceService struct {
	WorkplaceRepo      repo.WorkplaceRepo
	WorkplaceEventRepo repo.WorkplaceEventRepo
	RedisClient        redis.RedisClient
	DbHolder           *sqlx.DB
}

func NewWorkplaceService(workplaceRepo repo.WorkplaceRepo, workplaceEventRepo repo.WorkplaceEventRepo, db *sqlx.DB, redisClient *redis.RedisClient) WorkplaceService {
	return &workplaceService{
		WorkplaceRepo:      workplaceRepo,
		WorkplaceEventRepo: workplaceEventRepo,
		RedisClient:        *redisClient,
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

		workplace := model.Workplace{ID: workplaceIdInternal, Name: name, Size: size}
		workplaceEntity := model.CreateEventFromWorkplace(model.Created, model.Deferred, workplace)

		err = ws.WorkplaceEventRepo.Add(ctx, *workplaceEntity, tx)
		if err != nil {
			return err
		}

		workplaceId = workplaceIdInternal

		ws.RedisClient.CacheWorkplace(workplace)

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

	cachedWorkplace, founded, _ := ws.RedisClient.GetWorkplaceFromCache(workplaceID)
	if founded {
		logger.DebugKV(ctx, "DescribeWorkplace -- get workplace from Redis")
		return cachedWorkplace, nil
	}

	workplace, err := ws.WorkplaceRepo.DescribeWorkplace(ctx, workplaceID)
	if err != nil {
		logger.ErrorKV(ctx, "DescribeWorkplace -- failed", "err", err)
		return nil, err
	}

	ws.RedisClient.CacheWorkplace(*workplace)

	return workplace, nil
}

func (ws *workplaceService) ListWorkplaces(ctx context.Context, offset uint64, limit uint64) ([]model.Workplace, uint64, error) {
	span := tracer.CreateSpan(ctx, "Service ListWorkplaces")
	defer span.Close()

	workplaces, total, err := ws.WorkplaceRepo.ListWorkplaces(ctx, offset, limit)
	if err != nil {
		logger.ErrorKV(ctx, "ListWorkplaces -- failed", "err", err)
		return nil, 0, err
	}

	return workplaces, total, nil

}

func (ws *workplaceService) RemoveWorkplace(ctx context.Context, workplaceID uint64) (bool, error) {
	span := tracer.CreateSpan(ctx, "Service RemoveWorkplace")
	defer span.Close()

	ok, err := ws.WorkplaceRepo.RemoveWorkplace(ctx, workplaceID)
	if err != nil {
		logger.ErrorKV(ctx, "RemoveWorkplace -- failed", "err", err)
		return false, err
	}

	ws.RedisClient.DeleteCachedWorkplace(workplaceID)

	return ok, nil
}

func (ws *workplaceService) UpdateWorkplace(ctx context.Context, workplace model.Workplace) (bool, error) {

	err := database.WithTx(ctx, ws.DbHolder, func(ctx context.Context, tx *sqlx.Tx) error {
		span := tracer.CreateSpan(ctx, "Service UpdateWorkplace")
		defer span.Close()

		_, err := ws.WorkplaceRepo.UpdateWorkplace(ctx, workplace, tx)
		if err != nil {
			return err
		}

		workplaceEntity := model.CreateEventFromWorkplace(model.Updated, model.Deferred, model.Workplace{ID: workplace.ID, Name: workplace.Name, Size: workplace.Size})

		err = ws.WorkplaceEventRepo.Add(ctx, *workplaceEntity, tx)
		if err != nil {
			return err
		}

		ws.RedisClient.CacheWorkplace(workplace)

		return nil
	})

	if err != nil {
		logger.ErrorKV(ctx, "UpdateWorkplace -- failed", "err", err)
		return false, err
	}

	return true, nil
}
