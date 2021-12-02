package kafka

import (
	"github.com/Shopify/sarama"
	"time"
)

func NewSyncProducer(brokers []string, numRetries uint, retryDelaySec uint) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	config.Producer.Retry.Max = int(numRetries)
	config.Producer.Retry.Backoff = time.Duration(retryDelaySec) * time.Second

	producer, err := sarama.NewSyncProducer(brokers, config)

	return producer, err
}

func SendMessage(producer sarama.SyncProducer, topic string, message []byte) error {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.ByteEncoder(message),
	}
	_, _, err := producer.SendMessage(msg)
	return err
}