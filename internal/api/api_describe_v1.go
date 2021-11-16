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

func (o *workplaceAPI) DescribeWorkplaceV1(
	ctx context.Context,
	req *pb.DescribeWorkplaceV1Request,
) (*pb.DescribeWorkplaceV1Response, error) {

	logger.DebugKV(ctx, "DescribeWorkplaceV1 in", "req", req)

	span := tracer.CreateSpan(ctx, "API DescribeWorkplaceV1")
	defer span.Close()

	if err := req.Validate(); err != nil {
		logger.WarnKV(ctx, "DescribeWorkplaceV1 - invalid argument", "req", req)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	workplace, err := o.WorkplaceService.DescribeWorkplace(ctx, req.WorkplaceId)
	if err != nil {
		logger.ErrorKV(ctx, "DescribeWorkplaceV1 - failed", "err", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	if workplace == nil {
		logger.WarnKV(ctx, "DescribeWorkplaceV1 - workplace not found", "workplaceId", req.WorkplaceId)
		metrics.IncNotFoundErrors()

		return nil, status.Error(codes.NotFound, "workplace not found")
	}

	logger.DebugKV(ctx, "DescribeWorkplaceV1 out")

	item := mappers.WorkplaceToListItem(*workplace)
	return &pb.DescribeWorkplaceV1Response{
		Value: &item,
	}, nil
}
