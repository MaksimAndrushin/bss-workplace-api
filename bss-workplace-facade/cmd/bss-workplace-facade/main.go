package main

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
	"github.com/ozonmp/bss-workplace-facade/internal/config"
	"github.com/ozonmp/bss-workplace-facade/internal/database"
	"github.com/ozonmp/bss-workplace-facade/internal/server"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx := context.Background()

	var configFileName = "facade-config.yml"
	//var configFileName = "facade-config_local.yml"

	if err := config.ReadConfigYML(configFileName); err != nil {
		log.Fatal().Msgf("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	log.Info().
		Str("version", cfg.Project.Version).
		Str("commitHash", cfg.Project.CommitHash).
		Bool("debug", cfg.Project.Debug).
		Str("environment", cfg.Project.Environment).
		Msgf("Starting service: %s", cfg.Project.Name)

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SslMode,
	)

	db, err := database.NewPostgres(ctx, dsn, cfg.Database.Driver, cfg.Database.ReconnectCount, cfg.Database.ReconnectDelay)
	if err != nil {
		log.Fatal().Msgf("Failed init postgres, err - %v", err)
		return
	}
	defer db.Close()

	if err := server.NewGrpcServer(db).Start(ctx, &cfg); err != nil {
		log.Fatal().Msgf("Failed creating gRPC server, err - %v", err)
		return
	}

}
