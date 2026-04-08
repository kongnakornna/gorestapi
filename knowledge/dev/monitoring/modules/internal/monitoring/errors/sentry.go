// Package errors จัดการ error tracking ด้วย Sentry
// Package errors handles error tracking with Sentry
package errors

import (
	"log/slog"
	"time"

	"github.com/getsentry/sentry-go"
)

// InitSentry เริ่มต้น Sentry client
// InitSentry initializes Sentry client
func InitSentry(dsn string, environment string) error {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		Environment:      environment,
		TracesSampleRate: 1.0,
		AttachStacktrace: true,
	})
	if err != nil {
		slog.Error("Sentry initialization failed", "error", err)
		return err
	}
	slog.Info("Sentry initialized", "environment", environment)
	return nil
}

// CaptureError ส่ง error ไปยัง Sentry (ไม่ block)
// CaptureError sends error to Sentry (non-blocking)
func CaptureError(err error, tags map[string]string) {
	if err == nil {
		return
	}
	eventID := sentry.CaptureException(err)
	if tags != nil {
		sentry.ConfigureScope(func(scope *sentry.Scope) {
			scope.SetTags(tags)
		})
	}
	slog.Debug("Error sent to Sentry", "event_id", eventID)
}

// RecoverPanic ใช้ใน defer เพื่อ catch panic แล้วส่งไป Sentry
// RecoverPanic used in defer to catch panic and send to Sentry
func RecoverPanic() {
	if r := recover(); r != nil {
		sentry.CurrentHub().Recover(r)
		sentry.Flush(time.Second * 2)
		panic(r) // re-panic ถ้าต้องการให้ process จบ
	}
}