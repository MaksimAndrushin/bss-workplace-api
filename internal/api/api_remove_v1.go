package api

import (
	pb "github.com/ozonmp/bss-workplace-api/pkg/bss-workplace-api"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *workplaceAPI) RemoveWorkplaceV1(
	ctx context.Context,
	req *pb.RemoveWorkplaceV1Request,
) (*pb.RemoveWorkplaceV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("RemoveWorkplaceV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	ok, err := o.repo.RemoveWorkplace(ctx, req.WorkplaceId)
	if err != nil {
		log.Error().Err(err).Msg("DescribeWorkplaceV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if ok == false {
		log.Debug().Uint64("workplaceId", req.WorkplaceId).Msg("workplace not removed")
		totalWorkplaceNotFound.Inc()

		return nil, status.Error(codes.NotFound, "workplace not removed")
	}

	log.Debug().Msg("RemoveWorkplaceV1 - success")

	return &pb.RemoveWorkplaceV1Response{
		Found: ok,
	}, nil
}