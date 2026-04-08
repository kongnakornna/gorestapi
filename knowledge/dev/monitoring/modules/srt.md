icmongolang/
├── internal/
│   ├── monitoring/                     # ✅ โมดูลใหม่ (เพิ่มทั้งหมด)
│   │   ├── config/
│   │   │   └── monitoring_config.go    # อ่าน env และ flag เปิด-ปิด
│   │   ├── alert/
│   │   │   └── email_alert.go          # ส่งอีเมล alert
│   │   ├── metrics/
│   │   │   ├── prometheus.go           # Prometheus metrics registry
│   │   │   └── system_metrics.go       # CPU/RAM/Network collector
│   │   ├── tracing/
│   │   │   └── otel.go                 # OpenTelemetry init (Jaeger)
│   │   ├── logger/
│   │   │   └── slog_logger.go          # ตั้งค่า slog
│   │   ├── errors/
│   │   │   └── sentry.go               # Sentry init และ capture
│   │   ├── middleware/
│   │   │   └── monitoring_middleware.go # Middleware รวม (metrics/tracing/sentry)
│   │   └── handler/
│   │       └── monitoring_handler.go   # REST endpoints (/monitoring/*)
│   ├── ...                             # (โฟลเดอร์เดิม unchanged)
├── cmd/
│   └── api/
│       └── main.go                     # ✅ แก้ไข: import monitoring และใช้ config
├── .env.example                        # ✅ ตัวอย่าง environment variables
├── go.mod                              # ✅ เพิ่ม dependencies
└── ...