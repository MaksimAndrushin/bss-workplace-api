package mappers

import (
	"github.com/ozonmp/bss-workplace-facade/internal/model"
	pb "github.com/ozonmp/bss-workplace-facade/pkg/bss-workplace-facade"
)

func WorkplaceEventToListItem(workplaceEvent model.WorkplaceEvent) pb.WorkplaceEvent {
	return pb.WorkplaceEvent{
		Id:          workplaceEvent.ID,
		EventType:   uint32(workplaceEvent.Type),
		EventStatus: uint32(workplaceEvent.Status),
		Workplace: &pb.Workplace{
			Id:   workplaceEvent.Entity.ID,
			Name: workplaceEvent.Entity.Name,
			Size: workplaceEvent.Entity.Size,
		},
	}
}

func WorkplacesEventsToListItems(workplaceEvents []model.WorkplaceEvent) []*pb.WorkplaceEvent {
	var items []*pb.WorkplaceEvent
	for _, workplace := range workplaceEvents {
		item := WorkplaceEventToListItem(workplace)
		items = append(items, &item)
	}

	return items
}

//func WorkplacePBToWorkplaceModel(workplace pb.Workplace) *model.Workplace {
//	return &model.Workplace{
//		ID:      workplace.GetId(),
//		Name:    workplace.GetName(),
//		Size:    workplace.GetSize(),
//		Created: workplace.GetCreated().AsTime(),
//	}
//}
