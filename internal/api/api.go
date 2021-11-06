package api

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/ozonmp/bss-workplace-api/internal/repo"

	pb "github.com/ozonmp/bss-workplace-api/pkg/bss-workplace-api"
)

var (
	totalWorkplaceNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Name: "bss_workplace_api_workplace_not_found_total",
		Help: "Total number of workplaces that were not found",
	})
)

type workplaceAPI struct {
	pb.UnimplementedBssWorkplaceApiServiceServer
	repo repo.Repo
}

// NewWorkplaceAPI returns api of bss-workplace-api service
func NewWorkplaceAPI(r repo.Repo) pb.BssWorkplaceApiServiceServer {
	return &workplaceAPI{repo: r}
}