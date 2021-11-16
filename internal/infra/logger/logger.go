package logger

import (
	"context"
	"github.com/ozonmp/bss-workplace-api/internal/config"
	gelf "github.com/snovichkov/zap-gelf"
	"log"
	"os"

	"go.uber.org/zap/zapcore"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"go.uber.org/zap"
)

type ctxKey struct{}

var attachedLoggerKey = &ctxKey{}

var globalLogger *zap.SugaredLogger

func fromContext(ctx context.Context) *zap.SugaredLogger {
	var result *zap.SugaredLogger
	if attachedLogger, ok := ctx.Value(attachedLoggerKey).(*zap.SugaredLogger); ok {
		result = attachedLogger
	} else {
		result = globalLogger
	}

	jaegerSpan := opentracing.SpanFromContext(ctx)
	if jaegerSpan != nil {
		if spanCtx, ok := opentracing.SpanFromContext(ctx).Context().(jaeger.SpanContext); ok {
			result = result.With("trace-id", spanCtx.TraceID())
		}
	}

	return result
}

// ErrorKV write error log with additional metadata
func ErrorKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Errorw(message, kvs...)
}

// WarnKV write warning log with additional metadata
func WarnKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Warnw(message, kvs...)
}

// InfoKV write information log with additional metadata
func InfoKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Infow(message, kvs...)
}

// DebugKV write debug log with additional metadata
func DebugKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Debugw(message, kvs...)
}

// FatalKV write fatal log with additional metadata
func FatalKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Fatalw(message, kvs...)
}

// AttachLogger attach logger to the context
func AttachLogger(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, attachedLoggerKey, logger)
}

// CloneWithLevel clone the current context logger to the new one with changed log level
func CloneWithLevel(ctx context.Context, newLevel zapcore.Level) *zap.SugaredLogger {
	return fromContext(ctx).
		Desugar().
		WithOptions(WithLevel(newLevel)).
		Sugar()
}

// SetLogger set default logger
func SetLogger(newLogger *zap.SugaredLogger) {
	globalLogger = newLogger
}

func init() {
	notSugaredLogger, err := zap.NewProduction()
	if err != nil {
		log.Panic(err)
	}

	globalLogger = notSugaredLogger.Sugar()
}

func InitLogger(ctx context.Context, cfg config.Config) (syncFn func()) {
	loggingLevel := zap.InfoLevel
	if cfg.Project.Debug {
		loggingLevel = zap.DebugLevel
	}

	consoleCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		os.Stderr,
		zap.NewAtomicLevelAt(loggingLevel),
	)

	gelfCore, err := gelf.NewCore(
		gelf.Addr(cfg.Telemetry.GraylogPath),
		gelf.Level(loggingLevel),
	)
	if err != nil {
		FatalKV(ctx, "initLogger() error", "err", err)
	}

	notSugaredLogger := zap.New(zapcore.NewTee(consoleCore, gelfCore))

	sugaredLogger := notSugaredLogger.Sugar()
	SetLogger(sugaredLogger.With(
		"service", cfg.Project.Name,
	))

	return func() {
		err := notSugaredLogger.Sync()
		if err != nil {
			FatalKV(ctx, "initLogger() error", "err", err)
		}
	}
}