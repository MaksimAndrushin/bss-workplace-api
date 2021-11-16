package middleware

import (
	"context"
	"github.com/ozonmp/bss-workplace-api/internal/infra/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func ResponseLoggerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		response, err := handler(ctx, req)
		if err == nil {
			meta, ok := metadata.FromIncomingContext(ctx)
			if !ok {
				meta = nil
			}

			logger.DebugKV(ctx, "Raw response",
				"Server full method name", info.FullMethod,
				"Response Body", response,
				"Metadata", meta)
		} else {
			logger.ErrorKV(ctx, "Handler calling error", "err", err)
		}

		return response, err
	}
}
