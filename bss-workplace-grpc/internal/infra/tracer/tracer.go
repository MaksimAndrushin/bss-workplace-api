package tracer

import (
	"context"
	"github.com/ozonmp/bss-workplace-api/internal/infra/logger"
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"

	"github.com/ozonmp/bss-workplace-api/internal/config"

	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// NewTracer - returns new tracer.
func NewTracer(ctx context.Context, cfg *config.Config) (io.Closer, error) {
	cfgTracer := &jaegercfg.Configuration{
		ServiceName: cfg.Jaeger.Service,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: cfg.Jaeger.Host + cfg.Jaeger.Port,
		},
	}
	tracer, closer, err := cfgTracer.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		logger.ErrorKV(ctx, "failed init jaeger", "err", err)
		return nil, err
	}

	opentracing.SetGlobalTracer(tracer)
	logger.InfoKV(ctx, "Traces started")

	return closer, nil
}

type Span struct {
	span opentracing.Span
}

func CreateSpan(ctx context.Context, operationName string) *Span {

	spanFromContext := opentracing.SpanFromContext(ctx)
	var span opentracing.Span

	if spanFromContext == nil {
		span = opentracing.StartSpan(operationName)
	} else {
		span = opentracing.StartSpan(operationName, opentracing.FollowsFrom(spanFromContext.Context()))
	}

	return &Span{
		span: span,
	}
}

func (s *Span) Close() {
	s.span.Finish()
}
