package mappers

import (
	"github.com/ozonmp/bss-workplace-api/internal/model"
	pb "github.com/ozonmp/bss-workplace-api/pkg/bss-workplace-api"
)

func WorkplaceToListItem(workplace model.Workplace) pb.Workplace {
	return pb.Workplace{
		Id:  workplace.ID,
		Name: workplace.Name,
		Size: workplace.Size,
	}
}

func WorkplacesToListItems(workplaces []model.Workplace) []*pb.Workplace {
	var items []*pb.Workplace
	for _, workplace := range workplaces {
		item := WorkplaceToListItem(workplace)
		items = append(items, &item)
	}

	return items
}

func WorkplacePBToWorkplaceModel(workplace pb.Workplace) *model.Workplace {
	return &model.Workplace{
		ID: workplace.GetId(),
		Name: workplace.GetName(),
		Size: workplace.GetSize(),
		Created: workplace.GetCreated().AsTime(),
	}
}
