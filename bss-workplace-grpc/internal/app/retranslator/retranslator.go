package retranslator

import (
	"github.com/ozonmp/bss-workplace-api/internal/repo"
	"golang.org/x/net/context"
	"time"

	"github.com/ozonmp/bss-workplace-api/internal/app/consumer"
	"github.com/ozonmp/bss-workplace-api/internal/app/producer"
	"github.com/ozonmp/bss-workplace-api/internal/app/sender"
	"github.com/ozonmp/bss-workplace-api/internal/model"

	"github.com/gammazero/workerpool"
)

type Retranslator interface {
	Start(ctx context.Context)
	Close()
}

type RetranslatorConfig struct {
	ChannelSize uint64

	ConsumerCount  uint64
	ConsumeSize    uint64
	ConsumeTimeout time.Duration

	ProducerCount uint64
	WorkerCount   int

	Repo   repo.WorkplaceEventRepo
	Sender sender.EventSender
}

type retranslator struct {
	events     chan model.WorkplaceEvent
	consumer   consumer.Consumer
	producer   producer.Producer
	workerPool *workerpool.WorkerPool
}

func NewRetranslator(cfg RetranslatorConfig) Retranslator {
	events := make(chan model.WorkplaceEvent, cfg.ChannelSize)
	workerPool := workerpool.New(cfg.WorkerCount)

	consumer := consumer.NewDbConsumer(
		cfg.ConsumerCount,
		cfg.ConsumeSize,
		cfg.ConsumeTimeout,
		cfg.Repo,
		events)


	producer := producer.NewKafkaProducer(
		cfg.ProducerCount,
		cfg.Sender,
		events,
		workerPool,
		cfg.Repo,
		)

	return &retranslator{
		events:     events,
		consumer:   consumer,
		producer:   producer,
		workerPool: workerPool,
	}
}

func (r *retranslator) Start(ctx context.Context) {
	r.producer.Start(ctx)
	r.consumer.Start(ctx)
}

func (r *retranslator) Close() {
	r.consumer.Close()
	r.producer.Close()
	r.workerPool.StopWait()
}
