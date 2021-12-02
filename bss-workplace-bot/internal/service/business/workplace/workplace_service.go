package workplace

import (
	"context"
	bss_workplace_api "github.com/ozonmp/bss-workplace-bot/pkg/bss-workplace-api"
	bss_workplace_facade "github.com/ozonmp/bss-workplace-bot/pkg/bss-workplace-facade"
	"github.com/ozonmp/omp-bot/internal/model/business"
	//"google.golang.org/grpc"
)

type WorkplaceService interface {
	Describe(ctx context.Context, workplaceID uint64) (*business.Workplace, error)
	List(ctx context.Context, cursor uint64, limit uint64) ([]business.Workplace, uint64, error)
	Create(ctx context.Context, workplace business.Workplace) (uint64, error)
	Update(ctx context.Context, workplaceID uint64, workplace business.Workplace) error
	Remove(ctx context.Context, workplaceID uint64) (bool, error)
	ListEvents(ctx context.Context, cursor uint64, limit uint64) ([]business.WorkplaceEvent, uint64, error)
}

type GrpcWorkplaceService struct {
	GrpcApiClient bss_workplace_api.BssWorkplaceApiServiceClient
	GrpcFacadeClient bss_workplace_facade.BssFacadeEventsApiServiceClient
}

func NewGrpcWorkplaceService(grpcApiClient bss_workplace_api.BssWorkplaceApiServiceClient, grpcFacadeClient bss_workplace_facade.BssFacadeEventsApiServiceClient) *GrpcWorkplaceService {
	return &GrpcWorkplaceService{
		GrpcApiClient: grpcApiClient,
		GrpcFacadeClient: grpcFacadeClient,
	}
}

func (s *GrpcWorkplaceService) Describe(ctx context.Context, workplaceId uint64) (*business.Workplace, error) {

	req := bss_workplace_api.DescribeWorkplaceV1Request{
		WorkplaceId: workplaceId,
	}

	resp, err := s.GrpcApiClient.DescribeWorkplaceV1(ctx, &req)
	if err != nil {
		return nil, err
	}

	res := business.Workplace{
		ID:   resp.GetValue().GetId(),
		Name: resp.GetValue().GetName(),
		Size: resp.GetValue().GetSize(),
	}

	return &res, nil
}

func (s *GrpcWorkplaceService) List(ctx context.Context, cursor uint64, limit uint64) ([]business.Workplace, uint64, error) {

	req := bss_workplace_api.ListWorkplacesV1Request{
		Offset: cursor,
		Limit:  limit,
	}

	resp, err := s.GrpcApiClient.ListWorkplacesV1(ctx, &req)
	if err != nil {
		return nil, 0, err
	}

	workplaces := MapGrpcWorkplacesToModelWorkplaces(resp.GetItems())
	return workplaces, resp.GetTotal(), nil
}

func (s *GrpcWorkplaceService) Create(ctx context.Context, workplace business.Workplace) (uint64, error) {
	req := bss_workplace_api.CreateWorkplaceV1Request{
		Name: workplace.Name,
		Size: workplace.Size,
	}

	resp, err := s.GrpcApiClient.CreateWorkplaceV1(ctx, &req)
	if err != nil {
		return 0, err
	}

	return resp.GetWorkplaceId(), nil
}

func (s *GrpcWorkplaceService) Update(ctx context.Context, workplaceID uint64, workplace business.Workplace) error {
	req := bss_workplace_api.UpdateWorkplaceV1Request{
		Value: &bss_workplace_api.Workplace{
			Id:   workplaceID,
			Name: workplace.Name,
			Size: workplace.Size,
		},
	}

	_, err := s.GrpcApiClient.UpdateWorkplaceV1(ctx, &req)
	if err != nil {
		return err
	}

	return nil
}

func (s *GrpcWorkplaceService) Remove(ctx context.Context, workplaceID uint64) (bool, error) {
	req := bss_workplace_api.RemoveWorkplaceV1Request{
		WorkplaceId: workplaceID,
	}

	resp, err := s.GrpcApiClient.RemoveWorkplaceV1(ctx, &req)
	if err != nil {
		return false, err
	}

	return resp.GetFound(), nil
}

func (s *GrpcWorkplaceService) ListEvents(ctx context.Context, cursor uint64, limit uint64) ([]business.WorkplaceEvent, uint64, error) {

	req := bss_workplace_facade.ListEventsV1Request{
		Offset: cursor,
		Limit:  limit,
	}

	resp, err := s.GrpcFacadeClient.ListEventsV1(ctx, &req)
	if err != nil {
		return nil, 0, err
	}

	workplaces := MapGrpcEventsToModelEvents(resp.GetItems())
	return workplaces, resp.GetTotal(), nil
}