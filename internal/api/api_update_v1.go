package api

import (
	"github.com/ozonmp/bss-workplace-api/internal/api/mappers"
	"github.com/ozonmp/bss-workplace-api/internal/infra/logger"
	"github.com/ozonmp/bss-workplace-api/internal/infra/metrics"
	"github.com/ozonmp/bss-workplace-api/internal/infra/tracer"
	pb "github.com/ozonmp/bss-workplace-api/pkg/bss-workplace-api"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *workplaceAPI) UpdateWorkplaceV1(
	ctx context.Context,
	req *pb.UpdateWorkplaceV1Request,
) (*pb.UpdateWorkplaceV1Response, error) {

	logger.DebugKV(ctx, "UpdateWorkplaceV1 in", "req", req)

	span := tracer.CreateSpan(ctx, "API UpdateWorkplaceV1")
	defer span.Close()

	if err := req.Validate(); err != nil {
		logger.WarnKV(ctx, "UpdateWorkplaceV1 - invalid argument", "req", req)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if req.Value == nil {
		logger.WarnKV(ctx, "UpdateWorkplaceV1 - invalid argument", "req", req)
		return nil, status.Error(codes.InvalidArgument, "value is nil")
	}

	metrics.IncCudCount(metrics.Updated)

	ok, err := o.WorkplaceService.UpdateWorkplace(ctx, *mappers.WorkplacePBToWorkplaceModel(*req.GetValue()))
	if err != nil {
		logger.ErrorKV(ctx, "UpdateWorkplaceV1 - failed", "err", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	if ok == false {
		logger.WarnKV(ctx, "UpdateWorkplaceV1 - workplace not updated", "workplaceId", req.GetValue().GetId())
		metrics.IncNotFoundErrors()

		return nil, status.Error(codes.NotFound, "workplace not updated")
	}

	logger.DebugKV(ctx, "UpdateWorkplaceV1 out")

	return &pb.UpdateWorkplaceV1Response{
		 Updated: ok,
	}, nil
}
