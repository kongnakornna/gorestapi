// Package middleware combines all monitoring middlewares
package middleware

import (
	"log/slog"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"

	"icmongolang/internal/monitoring/errors"
	monmetrics "icmongolang/internal/monitoring/metrics"
)

// MonitoringMiddleware is the main middleware combining logging, metrics, tracing, sentry
func MonitoringMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// 1. Start tracing span
		tracer := otel.Tracer("icmongolang")
		ctx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))
		ctx, span := tracer.Start(ctx, r.URL.Path,
			trace.WithAttributes(
				semconv.HTTPMethod(r.Method),
				semconv.HTTPURL(r.URL.String()),
				semconv.HTTPTarget(r.URL.Path),
			),
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()

		r = r.WithContext(ctx)

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		defer func() {
			if rec := recover(); rec != nil {
				errors.CaptureError(nil, map[string]string{"panic": "true"})
				http.Error(ww, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(ww, r)

		duration := time.Since(start).Seconds()
		statusStr := strconv.Itoa(ww.Status())
		monmetrics.HttpRequestsTotal.WithLabelValues(r.Method, r.URL.Path, statusStr).Inc()
		monmetrics.HttpRequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)

		slog.Info("HTTP request",
			"method", r.Method,
			"path", r.URL.Path,
			"status", ww.Status(),
			"duration_ms", duration*1000,
			"remote_addr", r.RemoteAddr,
		)

		if ww.Status() >= 500 {
			errors.CaptureError(nil, map[string]string{
				"status": statusStr,
				"path":   r.URL.Path,
				"method": r.Method,
			})
			span.SetStatus(codes.Error, "HTTP "+statusStr+" error on "+r.URL.Path)
		}

		monmetrics.ActiveGoroutines.Set(float64(runtime.NumGoroutine()))
	})
}