// Package logger จัดการ structured logging ด้วย slog (Go 1.21+)
// Package logger provides structured logging using slog (Go 1.21+)
package logger

import (
	"log/slog"
	"os"
)

// InitLogger เริ่มต้น logger ระดับ global
// InitLogger initializes the global logger
// ไทย: รองรับ environment (dev=text, prod=json)
// English: Supports env (dev=text, prod=json)
func InitLogger(env string) {
	var handler slog.Handler
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	if env == "production" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)
	slog.Info("Logger initialized", "environment", env)
}