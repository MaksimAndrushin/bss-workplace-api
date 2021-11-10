package api

import (
	"github.com/ozonmp/bss-workplace-api/internal/api/mappers"
	pb "github.com/ozonmp/bss-workplace-api/pkg/bss-workplace-api"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *workplaceAPI) DescribeWorkplaceV1(
	ctx context.Context,
	req *pb.DescribeWorkplaceV1Request,
) (*pb.DescribeWorkplaceV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("DescribeWorkplaceV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	workplace, err := o.WorkplaceService.DescribeWorkplace(ctx, req.WorkplaceId)
	if err != nil {
		log.Error().Err(err).Msg("DescribeWorkplaceV1 -- failed")
		return nil, status.Error(codes.Internal, err.Error())
	}

	if workplace == nil {
		log.Debug().Uint64("workplaceId", req.WorkplaceId).Msg("workplace not found")
		totalWorkplaceNotFound.Inc()

		return nil, status.Error(codes.NotFound, "workplace not found")
	}

	log.Debug().Msg("DescribeWorkplaceV1 - success")

	item := mappers.WorkplaceToListItem(*workplace)
	return &pb.DescribeWorkplaceV1Response{
		Value: &item,
	}, nil
}
