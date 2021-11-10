package api

import (
	pb "github.com/ozonmp/bss-workplace-api/pkg/bss-workplace-api"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *workplaceAPI) CreateWorkplaceV1(
	ctx context.Context,
	req *pb.CreateWorkplaceV1Request,
) (*pb.CreateWorkplaceV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("CreateWorkplaceV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	workplaceId, err := o.WorkplaceService.CreateWorkplace(ctx, req.GetName(), req.GetSize())
	if err != nil {
		log.Error().Err(err).Msg("CreateWorkplaceV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Debug().Uint64("workplaceId", workplaceId).Msg("Workplace was created")

	return &pb.CreateWorkplaceV1Response{
		WorkplaceId: workplaceId,
	}, nil
}
