package workplace

import (
	pb "github.com/ozonmp/bss-workplace-bot/pkg/bss-workplace-api"
	pe "github.com/ozonmp/bss-workplace-bot/pkg/bss-workplace-facade"
	"github.com/ozonmp/omp-bot/internal/model/business"
)

func MapGrpcWorkplacesToModelWorkplaces(items []*pb.Workplace) []business.Workplace {
	res := make([]business.Workplace, 0, len(items))

	for _, item := range items {
		res = append(res, business.Workplace{
			ID:   item.GetId(),
			Name: item.GetName(),
			Size: item.GetSize(),
		})
	}

	return res
}

func MapGrpcEventsToModelEvents(items []*pe.WorkplaceEvent) []business.WorkplaceEvent {
	res := make([]business.WorkplaceEvent, 0, len(items))

	for _, item := range items {
		res = append(res, business.WorkplaceEvent{
			ID:     item.GetId(),
			Type:   business.EventType(item.GetEventType()),
			Status: business.EventStatus(item.GetEventStatus()),
			Entity: &business.Workplace{
				ID: item.GetWorkplace().GetId(),
				Name: item.GetWorkplace().GetName(),
				Size: item.GetWorkplace().GetSize(),
			},
		})
	}

	return res
}
