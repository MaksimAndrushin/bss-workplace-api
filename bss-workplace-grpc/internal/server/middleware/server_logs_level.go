package middleware

import (
	"context"
	"github.com/ozonmp/bss-workplace-api/internal/infra/logger"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
)

func LogsLevelInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			levels := md.Get("log-level")
			logger.InfoKV(ctx, "Got log level", "levels", levels)
			if len(levels) > 0 {
				if parsedLevel, ok := parseLevel(levels[0]); ok {
					newLogger := logger.CloneWithLevel(ctx, parsedLevel)
					ctx = logger.AttachLogger(ctx, newLogger)
				}
			}
		}
		return handler(ctx, req)
	}
}

func parseLevel(str string) (zapcore.Level, bool) {
	switch strings.ToLower(str) {
	case "debug":
		return zapcore.DebugLevel, true
	case "info":
		return zapcore.DebugLevel, true
	case "warn":
		return zapcore.DebugLevel, true
	case "error":
		return zapcore.DebugLevel, true
	default:
		return zapcore.DebugLevel, false
	}

}
