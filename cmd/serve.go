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

	// ✅ Monitoring imports
	monConfig "icmongolang/internal/monitoring/config"
	monErrors "icmongolang/internal/monitoring/errors"
	monLogger "icmongolang/internal/monitoring/logger"
	monMetrics "icmongolang/internal/monitoring/metrics"
	monTracing "icmongolang/internal/monitoring/tracing"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start http server with configured api",
	Long:  `Starts a http server and serves the configured api`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetCfg()

		appLogger := logger.NewApiLogger(cfg)
		appLogger.InitLogger()
		appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode)

		// ========== ✅ Monitoring module initialization ==========
		// 1. Load monitoring configuration
		monCfg := monConfig.LoadMonitoringConfig()

		// 2. Initialize slog logger (always, but used only if monitoring enabled)
		env := os.Getenv("APP_ENV")
		if env == "" {
			env = "development"
		}
		monLogger.InitLogger(env)

		// 3. If monitoring is enabled, start components
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

			// 3.2 Jaeger Tracing (using OTLP exporter)
			if monCfg.TracingEnabled {
				jaegerEndpoint := monCfg.JaegerEndpoint
				if jaegerEndpoint == "" {
					jaegerEndpoint = "http://localhost:4317" // OTLP gRPC default
				}
				shutdown := monTracing.InitTracer("icmongolang-api", jaegerEndpoint)
				if shutdown != nil {
					defer shutdown()
				}
			}

			// 3.3 System Metrics Collector (background goroutine)
			if monCfg.SystemMetricsEnabled {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()
				monMetrics.StartSystemMetricsCollector(ctx)
			}
		} else {
			slog.Info("Monitoring module is DISABLED")
		}
		// ========== End of monitoring initialization ==========

		// Database connection
		psqlDB, err := postgres.NewPsqlDB(cfg)
		if err != nil {
			appLogger.Fatalf("Postgresql init: %s", err)
			appLogger.Error("เชื่อมต่อ Database Postgresql ล้มเหลว", "error", err)
		} else {
			appLogger.Infof("Postgres connected")
			appLogger.Info("เชื่อมต่อ Database Postgresql สำเร็จ")
		}

		// Migrate if needed
		if cfg.Server.MigrateOnStart {
			err = Migrate(psqlDB)
			if err != nil {
				appLogger.Info("Can not migrate data - ไม่สามารถ migrate data ได้")
			} else {
				appLogger.Info("Data migrated สำเร็จ")
			}
		}

		// Redis clients
		redisClient := redis.NewRedis(cfg)
		taskRedisClient := distributor.NewRedisClient(cfg)

		// ✅ Create server with monitoring config
		srv, err := server.NewServer(
			cfg,
			psqlDB,
			redisClient,
			taskRedisClient,
			appLogger,
			&monCfg, // pass monitoring config
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