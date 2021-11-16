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

func (o *workplaceAPI) RemoveWorkplaceV1(
	ctx context.Context,
	req *pb.RemoveWorkplaceV1Request,
) (*pb.RemoveWorkplaceV1Response, error) {

	logger.DebugKV(ctx, "RemoveWorkplaceV1 in", "req", req)

	span := tracer.CreateSpan(ctx, "API RemoveWorkplaceV1")
	defer span.Close()

	if err := req.Validate(); err != nil {
		logger.WarnKV(ctx, "RemoveWorkplaceV1 - invalid argument", "req", req)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	metrics.IncCudCount(metrics.Removed)

	ok, err := o.WorkplaceService.RemoveWorkplace(ctx, req.WorkplaceId)
	if err != nil {
		logger.ErrorKV(ctx, "RemoveWorkplaceV1 - failed", "err", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	if ok == false {
		logger.WarnKV(ctx, "RemoveWorkplaceV1 - workplace not removed", "workplaceId", req.WorkplaceId)
		metrics.IncNotFoundErrors()

		return nil, status.Error(codes.NotFound, "workplace not removed")
	}

	logger.DebugKV(ctx, "RemoveWorkplaceV1 out")

	return &pb.RemoveWorkplaceV1Response{
		Found: ok,
	}, nil
}
