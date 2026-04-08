// ไฟล์: cmd/serve.go
package cmd

import (
	"context"
	"log/slog"
	"os"

	"github.com/spf13/cobra"

	"icmongolang/config"
	"icmongolang/internal/distributor"
	"icmongolang/internal/server"
	"icmongolang/pkg/db/postgres"
	"icmongolang/pkg/db/redis"
	"icmongolang/pkg/logger"

	// ✅ Import monitoring packages
	monConfig "icmongolang/internal/monitoring/config"
	monErrors "icmongolang/internal/monitoring/errors"
	monLogger "icmongolang/internal/monitoring/logger"
	monMetrics "icmongolang/internal/monitoring/metrics"
	monTracing "icmongolang/internal/monitoring/tracing"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start http server with configured api",
	Long:  `Starts a http server and serves the configured api`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetCfg()

		// ---------- Logger ตัวเดิม (Zap) ----------
		appLogger := logger.NewApiLogger(cfg)
		appLogger.InitLogger()
		appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode)

		// ========== ✅ เพิ่ม Monitoring Module ==========
		// 1. โหลด config การเปิด-ปิด monitoring
		monCfg := monConfig.LoadMonitoringConfig()

		// 2. เริ่มต้น slog logger (standard library) - ทำงานเสมอ (แต่จะใช้สำหรับ monitoring โดยเฉพาะ)
		env := os.Getenv("APP_ENV")
		if env == "" {
			env = "development"
		}
		monLogger.InitLogger(env)

		// cmd/serve.go (ส่วนที่เรียก monitoring)
		// 3. ถ้าเปิด monitoring ให้เริ่ม components ต่างๆ
		if monCfg.Enabled {
			slog.Info("🔍 Monitoring module is ENABLED")

			// 3.1 Sentry (error tracking)
			if monCfg.SentryEnabled && monCfg.SentryDSN != "" {
				if err := monErrors.InitSentry(monCfg.SentryDSN, env); err != nil {
					slog.Warn("Sentry init failed", "error", err)
				} else {
					defer monErrors.RecoverPanic()
				}
			}

			// 3.2 Jaeger Tracing
			if monCfg.TracingEnabled {
				jaegerEndpoint := monCfg.JaegerEndpoint
				if jaegerEndpoint == "" {
					jaegerEndpoint = "http://localhost:14268/api/traces"
				}
				shutdown := monTracing.InitTracer("icmongolang-api", jaegerEndpoint)
				if shutdown != nil {
					defer shutdown()
				}
			}

			// 3.3 System Metrics Collector
			if monCfg.SystemMetricsEnabled {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()
				monMetrics.StartSystemMetricsCollector(ctx)
			}
		}
		// ========== สิ้นสุดส่วนเพิ่ม Monitoring ==========

		// ----- ต่อเนื่องโค้ดเดิม -----
		psqlDB, err := postgres.NewPsqlDB(cfg)
		if err != nil {
			appLogger.Fatalf("Postgresql init: %s", err)
			appLogger.Error("เชื่อมต่อ Databse Postgresql ล้มเหลว", "error", err)
		} else {
			appLogger.Infof("Postgres connected")
			appLogger.Info("เชื่อมต่อ Databse Postgresql สำเร็จ")
		}

		if cfg.Server.MigrateOnStart {
			err = Migrate(psqlDB)
			if err != nil {
				appLogger.Info("Can not migrate data - ไม่สามารถ migrate data ได้")
			} else {
				appLogger.Info("Data migrated สำเร็จ")
			}
		}

		redisClient := redis.NewRedis(cfg)
		taskRedisClient := distributor.NewRedisClient(cfg)

		// ✅ สร้าง Server โดยส่ง monitoring config เข้าไป (เพื่อให้ server.go ใช้ middleware และ routes)
		srv, err := server.NewServer(
			cfg,
			psqlDB,
			redisClient,
			taskRedisClient,
			appLogger,
			&monCfg, // ส่ง monitoring config ไปให้ server ใช้
		)
		if err != nil {
			appLogger.Fatal(err)
		}
		srv.Start()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}