package sender

import (
	"github.com/Shopify/sarama"
	"github.com/ozonmp/bss-workplace-api/internal/infra/kafka"
	"github.com/ozonmp/bss-workplace-api/internal/model"
	bss_workplace_api "github.com/ozonmp/bss-workplace-api/pkg/bss-workplace-api"
	"google.golang.org/protobuf/proto"
)

type KafkaSender struct {
	Topic         string
	KafkaProducer sarama.SyncProducer
}

func NewKafkaSender(brokers []string, topic string) (*KafkaSender, error) {
	syncProducer, err := kafka.NewSyncProducer(brokers)
	if err != nil {
		return nil, err
	}

	return &KafkaSender{
		Topic:         topic,
		KafkaProducer: syncProducer,
	}, nil
}

func (k *KafkaSender) Send(subdomain *model.WorkplaceEvent) error {

	workplaceEventProto := bss_workplace_api.WorkplaceEvent {
		Id: subdomain.ID,
		EventType: uint32(subdomain.Type),
		EventStatus: uint32(subdomain.Status),
		Workplace: &bss_workplace_api.Workplace{
			Id: subdomain.Entity.ID,
			Name: subdomain.Entity.Name,
			Size: subdomain.Entity.Size,
		},
	}

	msg, err := proto.Marshal(&workplaceEventProto)
	if err != nil {
		//logger.ErrorKV(err)
		return err
	}

	err = kafka.SendMessage(k.KafkaProducer, k.Topic, msg)
	if err != nil {
		//logger.ErrorKV(err)
		return err
	}

	return nil
}
