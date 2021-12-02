package api

import (
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-workplace-facade/internal/repo"
	"github.com/ozonmp/bss-workplace-facade/internal/service"

	//"github.com/ozonmp/bss-workplace-facade/internal/service"
	pb "github.com/ozonmp/bss-workplace-facade/pkg/bss-workplace-facade"
)

type workplaceEventsFacadeAPI struct {
	pb.UnimplementedBssFacadeEventsApiServiceServer
	WorkplaceEventsService service.WorkplaceEventsService
}

// NewWorkplaceEventsFacadeAPI returns api of bss-workplace-evets-facade api service
func NewWorkplaceEventsFacadeAPI(workplaceEventRepo repo.WorkplaceEventRepo, db *sqlx.DB) pb.BssFacadeEventsApiServiceServer {
	return &workplaceEventsFacadeAPI{
		WorkplaceEventsService: service.NewWorkplaceEventsService(workplaceEventRepo, db),
	}
}
