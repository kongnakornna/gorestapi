// Package tracing initializes OpenTelemetry tracing exporter to Jaeger via OTLP
package tracing

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

// InitTracer initializes the tracer and sets it as global
func InitTracer(serviceName, jaegerEndpoint string) func() {
	// Create OTLP gRPC exporter (Jaeger supports OTLP)
	exp, err := otlptracegrpc.New(context.Background(),
		otlptracegrpc.WithEndpoint(jaegerEndpoint),
		otlptracegrpc.WithInsecure(), // use only for dev, remove for production with TLS
	)
	if err != nil {
		slog.Error("Failed to create OTLP exporter", "error", err)
		return nil
	}

	res, err := resource.New(context.Background(),
		resource.WithAttributes(
			semconv.ServiceName(serviceName),
		),
	)
	if err != nil {
		slog.Error("Failed to create resource", "error", err)
		return nil
	}

	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(res),
		tracesdk.WithSampler(tracesdk.AlwaysSample()),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{}, propagation.Baggage{},
	))

	slog.Info("Tracing initialized", "jaeger_endpoint", jaegerEndpoint)
	return func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			slog.Error("Failed to shutdown tracer", "error", err)
		}
	}
}