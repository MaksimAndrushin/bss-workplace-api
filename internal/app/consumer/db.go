package consumer

import (
	"github.com/ozonmp/bss-workplace-api/internal/repo"
	"golang.org/x/net/context"
	"sync"
	"time"

	"github.com/ozonmp/bss-workplace-api/internal/model"
)

type Consumer interface {
	Start(ctx context.Context)
	Close()
}

type consumer struct {
	n      uint64
	events chan<- model.WorkplaceEvent

	repo repo.WorkplaceEventRepo

	batchSize uint64
	timeout   time.Duration

	done chan interface{}
	wg   *sync.WaitGroup
}

type Config struct {
	n         uint64
	events    chan<- model.WorkplaceEvent
	repo      repo.WorkplaceEventRepo
	batchSize uint64
	timeout   time.Duration
}

func NewDbConsumer(
	n uint64,
	batchSize uint64,
	consumeTimeout time.Duration,
	repo repo.WorkplaceEventRepo,
	events chan<- model.WorkplaceEvent) Consumer {

	var wg = &sync.WaitGroup{}
	done := make(chan interface{})

	return &consumer{
		n:         n,
		batchSize: batchSize,
		timeout:   consumeTimeout,
		repo:      repo,
		events:    events,
		wg:        wg,
		done:      done,
	}
}

func (c *consumer) Start(ctx context.Context) {
	for i := uint64(0); i < c.n; i++ {
		c.wg.Add(1)

		go func() {
			defer c.wg.Done()
			ticker := time.NewTicker(c.timeout)
			for {
				select {
				case <-ticker.C:
					events, err := c.repo.Lock(ctx, c.batchSize, nil)
					if err != nil {
						continue
					}
					for _, event := range events {
						c.events <- event
					}
				case <-c.done:
					return
				}
			}
		}()
	}
}

func (c *consumer) Close() {
	close(c.done)
	c.wg.Wait()
}
