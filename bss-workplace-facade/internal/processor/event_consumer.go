package processor

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
	"github.com/ozonmp/bss-workplace-facade/internal/config"
	"github.com/ozonmp/bss-workplace-facade/internal/infra/kafka"
	"github.com/ozonmp/bss-workplace-facade/internal/model"
	"github.com/ozonmp/bss-workplace-facade/internal/repo"
	bss_workplace_facade "github.com/ozonmp/bss-workplace-facade/pkg/bss-workplace-facade"
	"github.com/rs/zerolog/log"
)

type EventProcessor struct {
	EventsRepo repo.WorkplaceEventRepo
	Consumer   *kafka.Consumer
	topic      string
}

func NewEventsProcessor(cfg config.Config, eventsRepo repo.WorkplaceEventRepo) (*EventProcessor, error) {
	consumer, err := kafka.NewKafkaConsumer(cfg.Kafka.Brokers, cfg.Kafka.GroupID,
		func(ctx context.Context, message *sarama.ConsumerMessage) error {
			var workplaceEvent bss_workplace_facade.WorkplaceEvent

			err := proto.Unmarshal(message.Value, &workplaceEvent)
			if err != nil {
				log.Error().Msgf("Message unmarshall error %v", err)
				return err
			}

			err = eventsRepo.AddEvent(ctx, *model.CreateEventFromProtoEvent(workplaceEvent))
			if err != nil {
				log.Error().Msgf("Add message to repo error %v", err)
				return err
			}

			log.Debug().Msgf("New kafka message from topic %v", message.Topic)
			return nil
		})

	if err != nil {
		return nil, err
	}

	return &EventProcessor{
		EventsRepo: eventsRepo,
		Consumer:   consumer,
		topic:      cfg.Kafka.Topic,
	}, nil

}

func (e *EventProcessor) StartProcessor(ctx context.Context) {
	e.Consumer.StartConsuming(ctx, e.topic)
}
