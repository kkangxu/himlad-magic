package tracer

import (
	"context"
	"github.com/asim/go-micro/v3/debug/trace"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)

type microTraceWrapper struct {
	tracer trace.Tracer
}

func WrapperTracer() {
	trace.DefaultTracer = &microTraceWrapper{
		tracer: trace.DefaultTracer,
	}
}

// Start a trace
func (t *microTraceWrapper) Start(ctx context.Context, name string) (context.Context, *trace.Span) {
	if ginctx, ok := ctx.(*gin.Context); ok {
		ctx = ginctx.Request.Context()
	}
	return t.tracer.Start(ctx, name)
}

// Finish the trace
func (t *microTraceWrapper) Finish(span *trace.Span) error {
	return t.tracer.Finish(span)
}

// Read the traces
func (t *microTraceWrapper) Read(opts ...trace.ReadOption) ([]*trace.Span, error) {
	return t.tracer.Read(opts...)

}

// TraceID 获取jager Trace ID
func TraceID(ctx context.Context) string {
	if span := opentracing.SpanFromContext(ctx); span != nil {
		if jaegerCtx, ok := span.Context().(jaeger.SpanContext); ok {
			return jaegerCtx.TraceID().String()
		}
	}
	return ""
}
