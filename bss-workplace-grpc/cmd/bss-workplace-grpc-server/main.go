package main

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
	"github.com/ozonmp/bss-workplace-api/internal/infra/logger"
	"github.com/ozonmp/bss-workplace-api/internal/infra/metrics"
	"github.com/ozonmp/bss-workplace-api/internal/infra/tracer"
	"github.com/rs/zerolog/log"

	"github.com/ozonmp/bss-workplace-api/internal/config"
	"github.com/ozonmp/bss-workplace-api/internal/database"
	"github.com/ozonmp/bss-workplace-api/internal/server"
)

var (
	batchSize uint = 3
)

func main() {

	ctx := context.Background()

	var configFileName = "grpc-server-config.yml"
	//var configFileName = "grpc-server-config_local.yml"

	if err := config.ReadConfigYML(configFileName); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	//migration := flag.Bool("migration", true, "Defines the migration start option")
	//flag.Parse()

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

	metrics.InitMetrics(&cfg)

	tracing, err := tracer.NewTracer(ctx, &cfg)
	if err != nil {
		logger.FatalKV(ctx, "Failed init tracing", "err", err)
		return
	}
	defer tracing.Close()

	if err := server.NewGrpcServer(db, batchSize).Start(ctx, &cfg); err != nil {
		logger.FatalKV(ctx, "Failed creating gRPC server", "err", err)
		return
	}

}
