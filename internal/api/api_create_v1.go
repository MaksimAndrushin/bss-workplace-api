package api

import (
	"github.com/ozonmp/bss-workplace-api/internal/infra/logger"
	"github.com/ozonmp/bss-workplace-api/internal/infra/metrics"
	"github.com/ozonmp/bss-workplace-api/internal/infra/tracer"
	pb "github.com/ozonmp/bss-workplace-api/pkg/bss-workplace-api"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *workplaceAPI) CreateWorkplaceV1(
	ctx context.Context,
	req *pb.CreateWorkplaceV1Request,
) (*pb.CreateWorkplaceV1Response, error) {

	logger.DebugKV(ctx, "CreateWorkplaceV1 in", "req", req)

	span := tracer.CreateSpan(ctx, "API CreateWorkplaceV1")
	defer span.Close()

	if err := req.Validate(); err != nil {
		logger.WarnKV(ctx, "CreateWorkplaceV1 - invalid argument", "req", req)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	metrics.IncCudCount(metrics.Created)

	workplaceId, err := o.WorkplaceService.CreateWorkplace(ctx, req.GetName(), req.GetSize())
	if err != nil {
		logger.ErrorKV(ctx, "CreateWorkplaceV1 - failed", "err", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	logger.DebugKV(ctx, "CreateWorkplaceV1 - successful", "workplaceId", workplaceId)
	logger.DebugKV(ctx, "CreateWorkplaceV1 out")

	return &pb.CreateWorkplaceV1Response{
		WorkplaceId: workplaceId,
	}, nil
}
