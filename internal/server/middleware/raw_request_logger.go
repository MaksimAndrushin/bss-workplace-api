package middleware

import (
	"context"
	"github.com/ozonmp/bss-workplace-api/internal/infra/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func RequestLoggerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {

		meta, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			meta = nil
		}

		logger.DebugKV(ctx, "Raw request",
			"Server full method name", info.FullMethod,
			"Request Body", req,
			"Metadata", meta)

		return handler(ctx, req)
	}
}
