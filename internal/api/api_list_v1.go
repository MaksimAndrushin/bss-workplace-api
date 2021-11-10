package api

import (
	"github.com/ozonmp/bss-workplace-api/internal/api/mappers"
	pb "github.com/ozonmp/bss-workplace-api/pkg/bss-workplace-api"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *workplaceAPI) ListWorkplacesV1(
	ctx context.Context,
	req *pb.ListWorkplacesV1Request,
) (*pb.ListWorkplacesV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("ListWorkplaceV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	workplaces, err := o.WorkplaceService.ListWorkplaces(ctx, req.GetOffset(), req.GetLimit())
	if err != nil {
		log.Error().Err(err).Msg("ListWorkplacesV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if workplaces == nil {
		log.Debug().Msg("Workplaces not found")
		totalWorkplaceNotFound.Inc()

		return nil, status.Error(codes.NotFound, "workplaces not found")
	}

	log.Debug().Msg("ListWorkplacesV1 - success")

	return &pb.ListWorkplacesV1Response{
		Items: mappers.WorkplacesToListItems(workplaces),
	}, nil
}
