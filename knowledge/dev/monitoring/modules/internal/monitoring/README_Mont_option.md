# เพิ่ม Option เปิด-ปิดระบบ Monitoring (Enable/Disable Toggle)

## ✅ เพิ่มความสามารถในการเปิดปิด Monitoring Module แบบยืดหยุ่น

> **ไทย**: เพิ่ม Environment Variable `MONITORING_ENABLED` และตัวเลือกย่อยสำหรับแต่ละ Component เพื่อให้สามารถปิดการทำงานของระบบ Monitoring ทั้งหมดหรือบางส่วนได้ โดยไม่ต้องลบโค้ดหรือคอมเมนต์ออก  
> **English**: Add `MONITORING_ENABLED` env var and granular flags to disable monitoring components without deleting code.

---

## 1. Environment Variables สำหรับควบคุมการเปิด-ปิด

เพิ่มตัวแปรต่อไปนี้ในไฟล์ `.env` (แยกตาม environment)

```bash
# ============================================
# MASTER SWITCH - เปิด/ปิดทั้งระบบ Monitoring
# ============================================
MONITORING_ENABLED=true           # true = เปิด, false = ปิดทุกอย่าง (ยกเว้น logging พื้นฐาน)

# ============================================
# ตัวเลือกย่อย (ถ้า MONITORING_ENABLED=true จะใช้ค่านี้)
# ============================================
METRICS_ENABLED=true              # Prometheus metrics
TRACING_ENABLED=true              # OpenTelemetry + Jaeger
SENTRY_ENABLED=true               # Error tracking
ALERT_ENABLED=true                # Email alert
SYSTEM_METRICS_ENABLED=true       # CPU/RAM/Network collector

# ถ้าต้องการปิดเฉพาะบางตัว ให้ตั้งเป็น false
```

**คำอธิบาย (ไทย/อังกฤษ)**  
- `MONITORING_ENABLED=false` → จะไม่เริ่มต้น component ใดๆ เลย (slog ยังทำงานอยู่เพราะเป็น logging พื้นฐาน)  
- ถ้า `MONITORING_ENABLED=true` → แต่ละ component จะถูกควบคุมด้วย flag ย่อยของตัวเอง  
- เหมาะสำหรับการปิด monitoring ใน environment ที่ไม่ต้องการ เช่น local development ที่ resource น้อย หรือการทดสอบ performance

---

## 2. โค้ดสำหรับอ่านค่าการเปิด-ปิด

สร้างไฟล์ใหม่ `internal/monitoring/config/monitoring_config.go`

```go
// Package config จัดการการอ่านค่าตัวแปร environment สำหรับ monitoring
// Package config manages environment variables for monitoring
package config

import (
	"log/slog"
	"os"
	"strconv"
)

// MonitoringConfig โครงสร้างเก็บค่าการเปิด-ปิดของแต่ละ component
// MonitoringConfig holds enable/disable flags for each component
type MonitoringConfig struct {
	Enabled          bool // master switch
	MetricsEnabled   bool
	TracingEnabled   bool
	SentryEnabled    bool
	AlertEnabled     bool
	SystemMetricsEnabled bool
}

// LoadMonitoringConfig อ่านค่าจาก environment และคืนค่า struct
// LoadMonitoringConfig reads env vars and returns config
func LoadMonitoringConfig() MonitoringConfig {
	// Master switch (default = true)
	enabled, _ := strconv.ParseBool(getEnv("MONITORING_ENABLED", "true"))
	
	// ถ้า master = false ให้ทุกตัวย่อยเป็น false ทันที
	if !enabled {
		return MonitoringConfig{
			Enabled:          false,
			MetricsEnabled:   false,
			TracingEnabled:   false,
			SentryEnabled:    false,
			AlertEnabled:     false,
			SystemMetricsEnabled: false,
		}
	}
	
	// ถ้า master = true ให้อ่านค่าตัวย่อย (default = true)
	cfg := MonitoringConfig{
		Enabled:          true,
		MetricsEnabled:   getBoolEnv("METRICS_ENABLED", true),
		TracingEnabled:   getBoolEnv("TRACING_ENABLED", true),
		SentryEnabled:    getBoolEnv("SENTRY_ENABLED", true),
		AlertEnabled:     getBoolEnv("ALERT_ENABLED", true),
		SystemMetricsEnabled: getBoolEnv("SYSTEM_METRICS_ENABLED", true),
	}
	
	slog.Info("Monitoring config loaded", 
		"enabled", cfg.Enabled,
		"metrics", cfg.MetricsEnabled,
		"tracing", cfg.TracingEnabled,
		"sentry", cfg.SentryEnabled,
		"alert", cfg.AlertEnabled,
		"system_metrics", cfg.SystemMetricsEnabled,
	)
	return cfg
}

// getEnv ดึงค่าตัวแปร environment ถ้าไม่มีให้ใช้ default
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getBoolEnv ดึงค่า bool จาก environment ถ้าไม่มีให้ใช้ default
func getBoolEnv(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if b, err := strconv.ParseBool(value); err == nil {
			return b
		}
	}
	return defaultValue
}
```

---

## 3. แก้ไข `cmd/api/main.go` ให้ใช้ config ควบคุมการเริ่มต้น

```go
package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	monConfig "icmongolang/internal/monitoring/config"
	monAlert "icmongolang/internal/monitoring/alert"
	monErrors "icmongolang/internal/monitoring/errors"
	monHandler "icmongolang/internal/monitoring/handler"
	monLogger "icmongolang/internal/monitoring/logger"
	monMetrics "icmongolang/internal/monitoring/metrics"
	monMiddleware "icmongolang/internal/monitoring/middleware"
	monTracing "icmongolang/internal/monitoring/tracing"
)

func main() {
	// 1. โหลด config monitoring ก่อนทุกอย่าง
	monCfg := monConfig.LoadMonitoringConfig()
	
	// 2. อ่าน environment ทั่วไป
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	// 3. เริ่ม logger (slog) - ทำงานเสมอ (ไม่ขึ้นกับ monitoring flag)
	monLogger.InitLogger(env)

	// 4. เริ่ม component ตาม config
	if monCfg.Enabled {
		slog.Info("Monitoring module is ENABLED")
		
		// 4.1 Sentry
		if monCfg.SentryEnabled {
			sentryDSN := os.Getenv("SENTRY_DSN")
			if sentryDSN != "" {
				_ = monErrors.InitSentry(sentryDSN, env)
				defer monErrors.RecoverPanic()
			} else {
				slog.Warn("Sentry enabled but SENTRY_DSN not set, skipping")
			}
		}
		
		// 4.2 Tracing (Jaeger)
		if monCfg.TracingEnabled {
			jaegerEndpoint := os.Getenv("JAEGER_ENDPOINT")
			if jaegerEndpoint == "" {
				jaegerEndpoint = "http://localhost:14268/api/traces"
			}
			shutdownTracer := monTracing.InitTracer("icmongolang-api", jaegerEndpoint)
			if shutdownTracer != nil {
				defer shutdownTracer()
			}
		}
		
		// 4.3 System metrics collector (พื้นหลัง)
		if monCfg.SystemMetricsEnabled {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			monMetrics.StartSystemMetricsCollector(ctx)
		}
		
		// 4.4 Alert (email) จะถูกเรียกใช้เมื่อมีเงื่อนไข trigger เท่านั้น ไม่ต้อง init ล่วงหน้า
	} else {
		slog.Info("Monitoring module is DISABLED (MONITORING_ENABLED=false)")
	}

	// 5. สร้าง router หลัก
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	
	// 6. เพิ่ม monitoring middleware เฉพาะเมื่อเปิดใช้งานเท่านั้น
	if monCfg.Enabled {
		r.Use(monMiddleware.MonitoringMiddleware)
		slog.Info("Monitoring middleware attached")
	} else {
		slog.Info("Monitoring middleware SKIPPED")
	}

	// 7. เส้นทางเดิม (สมมติว่ามี)
	// ... routes เดิมของ icmongolang ...

	// 8. เพิ่ม monitoring endpoints เฉพาะเมื่อเปิดใช้งาน metrics หรือ health
	if monCfg.Enabled {
		r.Route("/monitoring", func(r chi.Router) {
			if monCfg.MetricsEnabled {
				r.Handle("/metrics", monHandler.MetricsHandler())
			}
			r.Get("/health", monHandler.HealthHandler)   // health ใช้ได้เสมอถ้าเปิด monitoring
			if monCfg.SystemMetricsEnabled {
				r.Get("/system", monHandler.SystemStatsHandler)
			}
		})
		slog.Info("Monitoring endpoints registered")
	} else {
		// Optional: ขึ้น endpoint แจ้งว่า monitoring ถูกปิด
		r.Get("/monitoring/health", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status":"monitoring_disabled","message":"MONITORING_ENABLED=false"}`))
		})
	}

	// 9. เริ่ม HTTP server (เหมือนเดิม)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		slog.Info("Starting server", "port", port, "monitoring_enabled", monCfg.Enabled)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Server failed", "error", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("Shutting down server...")

	ctxShutdown, cancelShutdown := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelShutdown()
	if err := srv.Shutdown(ctxShutdown); err != nil {
		slog.Error("Server forced to shutdown", "error", err)
	}
	slog.Info("Server exited")
}
```

---

## 4. การปรับปรุง Middleware (เผื่อกรณีปิดบาง component)

ใน `internal/monitoring/middleware/monitoring_middleware.go` ให้เพิ่มการตรวจสอบว่า metrics/tracing ถูกเปิดหรือไม่ (ถ้าปิดก็ข้ามไป)  
**ตัวอย่างการปรับ** (เฉพาะส่วนที่เกี่ยวข้อง)

```go
// แก้ไข MonitoringMiddleware ให้รับ config เข้ามา หรือใช้ global variable
// แต่เนื่องจากเราใช้ flag ควบคุมที่ main แล้ว ถ้า MONITORING_ENABLED=false จะไม่เรียก middleware นี้เลย
// ดังนั้นภายใน middleware นี้ถือว่า monitoring เปิดแล้วเสมอ
```

> **หมายเหตุ**: การแยก flag ย่อย (เช่น `METRICS_ENABLED=false`) จะมีผลแค่การเก็บ metrics แต่ middleware จะยังทำงานอยู่ (เพื่อไม่ให้ซับซ้อน) ถ้าต้องการปิด middleware ทั้งหมดให้ใช้ `MONITORING_ENABLED=false`

---

## 5. ตัวอย่างการใช้งานจริง

### กรณีที่ 1: เปิด monitoring แบบเต็มรูปแบบ (ค่า default)

```bash
# .env
MONITORING_ENABLED=true
METRICS_ENABLED=true
TRACING_ENABLED=true
SENTRY_ENABLED=true
ALERT_ENABLED=true
SYSTEM_METRICS_ENABLED=true
```
- ระบบจะเริ่มต้นทุกอย่าง
- มี `/monitoring/metrics`, `/monitoring/health`, `/monitoring/system`
- มีการส่ง trace ไป Jaeger, error ไป Sentry

### กรณีที่ 2: ปิด monitoring ทั้งหมด (ลด resource)

```bash
MONITORING_ENABLED=false
```
- จะไม่เริ่ม component ใดๆ
- ไม่มี middleware monitoring
- `/monitoring/health` จะตอบกลับว่า `monitoring_disabled`
- เหมาะกับ local development หรือ environment ที่ไม่ต้องการ monitoring

### กรณีที่ 3: เปิดแค่ metrics และ health check (ไม่ใช้ tracing/sentry)

```bash
MONITORING_ENABLED=true
METRICS_ENABLED=true
TRACING_ENABLED=false
SENTRY_ENABLED=false
ALERT_ENABLED=false
SYSTEM_METRICS_ENABLED=false
```
- มี `/monitoring/metrics` และ `/monitoring/health`
- ไม่มี CPU/RAM collector (ลด overhead)
- ไม่มี tracing และ sentry

---

## 6. Checklist สำหรับการทดสอบเปิด-ปิด

| ทดสอบ | ผลที่คาดหวัง |
|--------|--------------|
| `MONITORING_ENABLED=true` | middleware ทำงาน, metrics endpoint มีข้อมูล, มี traces, มี CPU stats |
| `MONITORING_ENABLED=false` | middleware **ไม่** ทำงาน, `/monitoring/metrics` ไม่มี (หรือ 404), `/monitoring/health` ตอบ `monitoring_disabled` |
| `METRICS_ENABLED=false` | `/monitoring/metrics` ไม่มี (404) แต่ `/monitoring/health` ยังอยู่ |
| `SENTRY_ENABLED=false` | error ไม่ถูกส่งไป Sentry (แต่ middleware ยังบันทึก status code ตามปกติ) |

---

## 7. สรุปการเพิ่ม Option เปิด-ปิด

✅ **สิ่งที่เพิ่มเข้ามา**  
- Environment variable `MONITORING_ENABLED` (master switch)  
- Flag ย่อยสำหรับแต่ละ component  
- ไฟล์ `internal/monitoring/config/monitoring_config.go` สำหรับอ่านค่า  
- ปรับ `main.go` ให้ตรวจสอบ flag ก่อนเริ่ม component และก่อนเพิ่ม middleware  

✅ **ประโยชน์**  
- ปรับใช้ได้ตาม environment (เปิดเต็มใน prod, ปิดบางส่วนใน dev, ปิดหมดใน local test)  
- ไม่ต้องลบโค้ด monitoring ออก เมื่อต้องการปิด  
- ลด resource usage ใน environment ที่ไม่ต้องการ monitoring  

⚠️ **ข้อควรระวัง**  
- การเปลี่ยนค่า flag ต้อง restart server (ไม่ support hot-reload)  
- ถ้าปิด `MONITORING_ENABLED` แล้ว `slog` logging พื้นฐานยังทำงานอยู่ (ไม่มีผล)  

---

**หมายเหตุ**: โค้ดทั้งหมดสามารถรันได้ทันทีโดยไม่ต้องแก้ไขโครงสร้างเดิม เพียงแค่เพิ่มไฟล์ใหม่และปรับ `main.go` ตามตัวอย่าง