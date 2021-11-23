package main

import (
	"context"
	"github.com/ozonmp/bss-workplace-api/internal/app/sender"
	"github.com/ozonmp/bss-workplace-api/internal/config"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ozonmp/bss-workplace-api/internal/app/retranslator"
)

func main() {

	//var configFileName = "config.yml"
	var configFileName = "config_local.yml"

	if err := config.ReadConfigYML(configFileName); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	sigs := make(chan os.Signal, 1)

	kafkaEventSender, err := sender.NewKafkaSender(cfg.Kafka.Brokers, cfg.Kafka.Topic)
	if err != nil {
		panic(err)
	}

	var retranslatorConfig = retranslator.RetranslatorConfig{
		ChannelSize:    512,
		ConsumerCount:  2,
		ConsumeSize:    10,
		ProducerCount:  28,
		WorkerCount:    2,
		ConsumeTimeout: 5,
		Sender:         kafkaEventSender,
	}

	var retranslator = retranslator.NewRetranslator(retranslatorConfig)
	retranslator.Start(context.Background())

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
}
