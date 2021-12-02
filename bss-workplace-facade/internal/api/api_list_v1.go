package api

import (
	"context"
	"github.com/ozonmp/bss-workplace-facade/internal/api/mappers"
	pb "github.com/ozonmp/bss-workplace-facade/pkg/bss-workplace-facade"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *workplaceEventsFacadeAPI) ListEventsV1(
	ctx context.Context,
	req *pb.ListEventsV1Request,
) (*pb.ListEventsV1Response, error) {

	log.Debug().Msgf("ListEventsV1 in. Req = %v", req)

	if err := req.Validate(); err != nil {
		log.Warn().Msgf("ListEventsV1 - invalid argument. Req = %v", req)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	events, total, err := o.WorkplaceEventsService.ListWorkplacesEvents(ctx, req.GetOffset(), req.GetLimit())
	if err != nil {
		log.Error().Msgf( "ListEventsV1 - failed. Err - %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	if events == nil {
		log.Warn().Msg( "ListEventsV1 - events not found")
		return nil, status.Error(codes.NotFound, "events not found")
	}

	log.Debug().Msgf("ListEventsV1 out")

	return &pb.ListEventsV1Response{
		Items: mappers.WorkplacesEventsToListItems(events),
		Total: total,
	}, nil
}