internal/monitoring/
├── config/
│   └── monitoring_config.go       ✅ package config
├── alert/
│   └── email_alert.go             ✅ package alert
├── metrics/
│   ├── prometheus.go              ✅ package metrics
│   └── system_metrics.go          ✅ package metrics
├── tracing/
│   └── otel.go                    ✅ package tracing
├── logger/
│   └── slog_logger.go             ✅ package logger
├── errors/
│   └── sentry.go                  ✅ package errors
├── middleware/
│   └── monitoring_middleware.go   ✅ package middleware (ต้องมี import slog)
└── handler/
    └── monitoring_handler.go      ✅ package handler
	
	
	
	
# ลบไฟล์ handler ที่อยู่ใน middleware folder (ถ้ายังมี)
rm -f internal/monitoring/middleware/monitoring_handler.go

# ลบ vendor และ rebuild
rm -rf vendor/

go mod vendor
go clean -cache

# รัน server
go run main.go serve
	
	
	
##
rm -rf vendor
go mod vendor
go clean -cache
go mod tidy
go mod download
go mod verify
rm -rf docs/# สร้าง docs ใหม่
swag init
go run  main.go serve
	
	
	
	
	
 # Health endpoint
curl http://localhost:8088/monitoring/health

# Metrics endpoint (ถ้าเปิด metrics)
curl http://localhost:8088/monitoring/metrics

# System stats
curl http://localhost:8088/monitoring/system