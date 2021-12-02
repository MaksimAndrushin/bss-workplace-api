package main

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
	"github.com/ozonmp/bss-workplace-api/internal/app/retranslator"
	"github.com/ozonmp/bss-workplace-api/internal/app/sender"
	"github.com/ozonmp/bss-workplace-api/internal/config"
	"github.com/ozonmp/bss-workplace-api/internal/database"
	"github.com/ozonmp/bss-workplace-api/internal/infra/logger"
	"github.com/ozonmp/bss-workplace-api/internal/repo"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	batchSize uint = 2
)

func main() {
	ctx := context.Background()

	var configFileName = "retranslator-config.yml"
	//var configFileName = "retranslator-config_local.yml"

	if err := config.ReadConfigYML(configFileName); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	log.Info().
		Str("version", cfg.Project.Version).
		Str("commitHash", cfg.Project.CommitHash).
		Bool("debug", cfg.Project.Debug).
		Str("environment", cfg.Project.Environment).
		Msgf("Starting service: %s", cfg.Project.Name)

	syncLogger := logger.InitLogger(ctx, cfg)
	defer syncLogger()

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SslMode,
	)

	logger.InfoKV(ctx, "DSN", "dsn = ", dsn)

	db, err := database.NewPostgres(ctx, dsn, cfg.Database.Driver, cfg.Database.ReconnectCount, cfg.Database.ReconnectDelay)
	if err != nil {
		logger.FatalKV(ctx, "Failed init postgres", "err", err)
		return
	}
	defer db.Close()

	workplaceEventRepo := repo.NewWorkplaceEventRepo(db, batchSize)

	kafkaEventSender, err := sender.NewKafkaSender(cfg.Kafka.Brokers, cfg.Kafka.Topic, cfg.Kafka.ResendCount, cfg.Kafka.ResendDelay)
	if err != nil {
		logger.FatalKV(ctx, "Failed init kafka sender", "err", err)
	}

	var retranslatorConfig = retranslator.RetranslatorConfig{
		ChannelSize:    cfg.Retranslator.ChannelSize,
		ConsumerCount:  cfg.Retranslator.ConsumerCount,
		ConsumeSize:    cfg.Retranslator.ConsumeSize,
		ProducerCount:  cfg.Retranslator.ProducerCount,
		WorkerCount:    cfg.Retranslator.WorkerCount,
		ConsumeTimeout: time.Duration(cfg.Retranslator.ConsumeTimeout),
		Repo:           workplaceEventRepo,
		Sender:         kafkaEventSender,
	}

	var retranslator = retranslator.NewRetranslator(retranslatorConfig)
	retranslator.Start(ctx)

	logger.InfoKV(ctx,"The retranslator is ready to work")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}
