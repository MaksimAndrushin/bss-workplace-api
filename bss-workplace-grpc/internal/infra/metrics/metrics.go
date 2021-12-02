package metrics

import (
	"github.com/ozonmp/bss-workplace-api/internal/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var notFoundErrors prometheus.Counter
var cudCount *prometheus.CounterVec
var eventsCount prometheus.Counter

func InitMetrics(cfg *config.Config) {
	notFoundErrors = promauto.NewCounter(prometheus.CounterOpts{
		Subsystem: "bss_workplace_api",
		Name:      "errors_not_found_count",
		Help:      "Number of not found errors",
	})

	cudCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Subsystem: "bss_workplace_api",
		Name:      "cud_operations_count",
		Help:      "Number of CUD operations",
	}, []string{"type"})

	eventsCount = promauto.NewCounter(prometheus.CounterOpts{
		Subsystem: "bss_workplace_api",
		Name:      "retranslator_events_count",
		Help:      "Number of events processed by retranslator",
	})
}

//go:generate stringer -type=eventType
type eventType uint

const (
	_ = eventType(iota)
	Created
	Updated
	Removed
)

func IncNotFoundErrors() {
	if notFoundErrors == nil {
		return
	}

	notFoundErrors.Inc()
}

func IncCudCount(eventType eventType) {
	if cudCount == nil {
		return
	}

	cudCount.WithLabelValues(eventType.String()).Inc()
}

func AddRetranslatorEvents(eventsCnt int) {
	if eventsCount == nil {
		return
	}

	eventsCount.Add(float64(eventsCnt))
}
