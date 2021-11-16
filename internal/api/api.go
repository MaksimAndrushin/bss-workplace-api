package api

import (
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-workplace-api/internal/repo"
	"github.com/ozonmp/bss-workplace-api/internal/service"
	pb "github.com/ozonmp/bss-workplace-api/pkg/bss-workplace-api"
)

//var (
//	totalWorkplaceNotFound = promauto.NewCounter(prometheus.CounterOpts{
//		Name: "bss_workplace_api_workplace_not_found_total",
//		Help: "Total number of workplaces that were not found",
//	})
//)

type workplaceAPI struct {
	pb.UnimplementedBssWorkplaceApiServiceServer
	WorkplaceService service.WorkplaceService
}

// NewWorkplaceAPI returns api of bss-workplace-api service
func NewWorkplaceAPI(workplaceRepo repo.WorkplaceRepo, workplaceEventRepo repo.WorkplaceEventRepo, db *sqlx.DB) pb.BssWorkplaceApiServiceServer {
	return &workplaceAPI{
		WorkplaceService: service.NewWorkplaceService(workplaceRepo, workplaceEventRepo, db),
	}
}


