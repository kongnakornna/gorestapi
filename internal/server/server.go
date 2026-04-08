package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	// "github.com/go-chi/chi/v5"   // ❌ REMOVED – not used directly
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"icmongolang/config"
	"icmongolang/internal/middleware"
	"icmongolang/pkg/logger"

	monConfig "icmongolang/internal/monitoring/config"
	monHandler "icmongolang/internal/monitoring/handler"
	monMiddleware "icmongolang/internal/monitoring/middleware"
)

const (
	maxHeaderBytes = 1 << 20
	ctxTimeout     = 5
)

// Server provides an http.Server.
type Server struct {
	server *http.Server
	cfg    *config.Config
	db     *gorm.DB
	logger logger.Logger
	monCfg *monConfig.MonitoringConfig
}

// NewServer creates and configures an APIServer serving all application routes.
func NewServer(
	cfg *config.Config,
	db *gorm.DB,
	redisClient *redis.Client,
	taskRedisClient *asynq.Client,
	logger logger.Logger,
	monCfg *monConfig.MonitoringConfig,
) (*Server, error) {
	logger.Info("configuring server...")

	mainHandler, err := New(db, redisClient, taskRedisClient, cfg, logger)
	if err != nil {
		return nil, err
	}

	buildServer := func(finalHandler http.Handler) *Server {
		addr := cfg.Server.Port
		if !strings.Contains(addr, ":") {
			addr = ":" + addr
		}
		return &Server{
			server: &http.Server{
				Addr:           addr,
				Handler:        finalHandler,
				ReadTimeout:    time.Second * time.Duration(cfg.Server.ReadTimeout),
				WriteTimeout:   time.Second * time.Duration(cfg.Server.WriteTimeout),
				MaxHeaderBytes: maxHeaderBytes,
			},
			cfg:    cfg,
			db:     db,
			logger: logger,
			monCfg: monCfg,
		}
	}

	if monCfg != nil && monCfg.Enabled {
		logger.Info("Adding monitoring routes and middleware...")
		mux := mainHandler // *chi.Mux (type is known from New(), but no direct chi import needed)
		// ✅ Add this line – registers /monitoring endpoints 
		monHandler.MapMonitoringRoutes(mux, monCfg)
		handlerWithLogging := middleware.LoggingMiddleware(mux)
		handlerWithMonitoring := monMiddleware.MonitoringMiddleware(handlerWithLogging)
		return buildServer(handlerWithMonitoring), nil
	}

	handlerWithLogging := middleware.LoggingMiddleware(mainHandler)
	return buildServer(handlerWithLogging), nil
}

// Start runs ListenAndServe on the http.Server with graceful shutdown.
func (srv *Server) Start() {
	srv.logger.Info("starting server...")

	go func() {
		srv.logger.Infof("Listening on %s\n", srv.server.Addr)
		if err := srv.server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	sig := <-quit

	srv.logger.Infof("Shutting down server... Reason: %s", sig)

	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()

	if err := srv.server.Shutdown(ctx); err != nil {
		panic(err)
	}
	srv.logger.Info("Server gracefully stopped")
}