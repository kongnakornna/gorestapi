// Package metrics รวบรวม Prometheus metrics ทั้งหมด
// Package metrics collects all Prometheus metrics
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// ตัวแปร metrics แบบ global
// Global metric variables
var (
	// HttpRequestsTotal นับจำนวน request ทั้งหมด (method, path, status)
	// HttpRequestsTotal counts total HTTP requests by method, path, status
	HttpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	// HttpRequestDuration เวลาตอบสนองของ request
	// HttpRequestDuration measures request latency in seconds
	HttpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request latency in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)

	// ActiveGoroutines จำนวน goroutine ที่กำลังทำงาน
	// ActiveGoroutines number of active goroutines
	ActiveGoroutines = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "go_goroutines_active",
			Help: "Number of active goroutines",
		},
	)
)