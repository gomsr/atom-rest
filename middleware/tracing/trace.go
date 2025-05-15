package tracing

import (
	"github.com/gin-gonic/gin"
	"github.com/gomsr/atom-log/rich"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"

	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	oteltrace "go.opentelemetry.io/otel/trace"
)

var tracer oteltrace.Tracer

func init() {
	tracer = sdktrace.NewTracerProvider().Tracer("go-zero")
}

// TraceHandler trace request
func TraceHandler(l *zap.Logger) gin.HandlerFunc {

	return func(c *gin.Context) {
		propagator := otel.GetTextMapPropagator()

		spanName := c.Request.URL.Path
		ctx := propagator.Extract(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))
		spanCtx, span := tracer.Start(
			ctx,
			spanName,
			oteltrace.WithSpanKind(oteltrace.SpanKindServer),
			oteltrace.WithAttributes(semconv.HTTPServerAttributesFromHTTPRequest("wpp", spanName, c.Request)...),
		)
		defer span.End()

		c.Request = c.Request.WithContext(spanCtx)
		c.Set("tlog", rich.WithLogger(l).WithContext(spanCtx))
		// convenient for tracking error messages
		propagator.Inject(c, propagation.HeaderCarrier(c.Writer.Header()))

		c.Next()
	}
}
