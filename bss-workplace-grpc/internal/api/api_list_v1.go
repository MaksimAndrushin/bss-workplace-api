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

func (o *workplaceAPI) ListWorkplacesV1(
	ctx context.Context,
	req *pb.ListWorkplacesV1Request,
) (*pb.ListWorkplacesV1Response, error) {

	logger.DebugKV(ctx, "ListWorkplacesV1 in", "req", req)

	span := tracer.CreateSpan(ctx, "API ListWorkplacesV1")
	defer span.Close()

	if err := req.Validate(); err != nil {
		logger.WarnKV(ctx, "ListWorkplacesV1 - invalid argument", "req", req)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	workplaces, total, err := o.WorkplaceService.ListWorkplaces(ctx, req.GetOffset(), req.GetLimit())
	if err != nil {
		logger.ErrorKV(ctx, "ListWorkplacesV1 - failed", "err", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	if workplaces == nil {
		logger.WarnKV(ctx, "ListWorkplacesV1 - workplaces not found")
		metrics.IncNotFoundErrors()

		return nil, status.Error(codes.NotFound, "workplaces not found")
	}

	logger.DebugKV(ctx, "ListWorkplacesV1 out")

	return &pb.ListWorkplacesV1Response{
		Items: mappers.WorkplacesToListItems(workplaces),
		Total: total,
	}, nil
}
