package server

import (
	"context"
	"fmt"
	"github.com/ozonmp/bss-workplace-facade/internal/api"
	"github.com/ozonmp/bss-workplace-facade/internal/processor"
	"github.com/ozonmp/bss-workplace-facade/internal/repo"
	"github.com/rs/zerolog/log"

	"net"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	"github.com/ozonmp/bss-workplace-facade/internal/config"
	pb "github.com/ozonmp/bss-workplace-facade/pkg/bss-workplace-facade"
)

// GrpcServer is gRPC server
type GrpcServer struct {
	db *sqlx.DB
}

// NewGrpcServer returns gRPC server with supporting of batch listing
func NewGrpcServer(db *sqlx.DB) *GrpcServer {
	return &GrpcServer{
		db: db,
	}
}

// Start method runs server
func (s *GrpcServer) Start(ctx context.Context, cfg *config.Config) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcAddr := fmt.Sprintf("%s:%v", cfg.Grpc.Host, cfg.Grpc.Port)

	isReady := &atomic.Value{}
	isReady.Store(false)

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	defer l.Close()

	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: time.Duration(cfg.Grpc.MaxConnectionIdle) * time.Minute,
			Timeout:           time.Duration(cfg.Grpc.Timeout) * time.Second,
			MaxConnectionAge:  time.Duration(cfg.Grpc.MaxConnectionAge) * time.Minute,
			Time:              time.Duration(cfg.Grpc.Timeout) * time.Minute,
		}),
	)

	workplaceEventRepo := repo.NewWorkplaceEventRepo(s.db, 3)

	eventPprocessor, err := processor.NewEventsProcessor(*cfg, workplaceEventRepo)
	if err != nil {
		return err
	}
	eventPprocessor.StartProcessor(ctx)

	pb.RegisterBssFacadeEventsApiServiceServer(grpcServer, api.NewWorkplaceEventsFacadeAPI(workplaceEventRepo, s.db))

	go func() {
		log.Info().Msgf("GRPC Server started. Bind address: %s", grpcAddr)
		if err := grpcServer.Serve(l); err != nil {
			log.Fatal().Msgf("Failed running gRPC server: %v", err)
		}
	}()

	go func() {
		time.Sleep(2 * time.Second)
		isReady.Store(true)
		log.Info().Msg("The service is ready to accept requests")
	}()

	if cfg.Project.Debug {
		reflection.Register(grpcServer)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		log.Info().Msgf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		log.Info().Msgf("ctx.Done: %v ", done)
	}

	isReady.Store(false)

	grpcServer.GracefulStop()
	log.Info().Msg("grpcServer shut down correctly")

	return nil
}
