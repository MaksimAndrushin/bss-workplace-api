package main

import (
	"context"
	"fmt"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	bss_workplace_api "github.com/ozonmp/bss-workplace-bot/pkg/bss-workplace-api"
	bss_workplace_facade "github.com/ozonmp/bss-workplace-bot/pkg/bss-workplace-facade"
	"github.com/ozonmp/omp-bot/internal/config"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	routerPkg "github.com/ozonmp/omp-bot/internal/app/router"
)

func main() {
	_ = godotenv.Load()

	token, found := os.LookupEnv("TOKEN")
	if !found {
		log.Fatal().Msg("environment variable TOKEN not found in .env")
	}

	ctx := context.Background()

	var configFileName = "bot-config.yml"
	//var configFileName = "bot-config_local.yml"

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

	apiClient, err := createGrpcWorkplaceApiConnection(ctx, cfg)
	if err != nil {
		log.Fatal().Err(err)
	}

	facadeClient, err := createGrpcFacadeConnection(ctx, cfg)
	if err != nil {
		log.Fatal().Err(err)
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal().Err(err)
	}

	// Uncomment if you want debugging
	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal().Err(err)
	}

	routerHandler := routerPkg.NewRouter(bot, apiClient, facadeClient)

	for update := range updates {
		routerHandler.HandleUpdate(ctx, update)
	}
}

func createGrpcFacadeConnection(ctx context.Context, cfg config.Config) (bss_workplace_facade.BssFacadeEventsApiServiceClient, error) {
	dialContext, err := createGrpcDialContext(ctx, cfg.FacadeSvc.Host, cfg.FacadeSvc.Port, cfg.FacadeSvc.RetryCount, cfg.FacadeSvc.RetryDelaySec)
	if err != nil {
		return nil, err
	}

	facadeClient := bss_workplace_facade.NewBssFacadeEventsApiServiceClient(dialContext)
	return facadeClient, nil
}

func createGrpcWorkplaceApiConnection(ctx context.Context, cfg config.Config) (bss_workplace_api.BssWorkplaceApiServiceClient, error) {
	dialContext, err := createGrpcDialContext(ctx, cfg.WorkplaceSvc.Host, cfg.WorkplaceSvc.Port, cfg.WorkplaceSvc.RetryCount, cfg.WorkplaceSvc.RetryDelaySec)
	if err != nil {
		return nil, err
	}

	apiClient := bss_workplace_api.NewBssWorkplaceApiServiceClient(dialContext)
	return apiClient, err
}

func createGrpcDialContext(ctx context.Context, host string, port int, retryCount uint, retryDelay uint) (*grpc.ClientConn, error) {

	opts := []grpc_retry.CallOption{
		grpc_retry.WithMax(retryCount),
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(time.Duration(retryDelay) * time.Second)),
	}

	grpcConn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("%s:%d", host, port),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)),
	)

	return grpcConn, err
}
