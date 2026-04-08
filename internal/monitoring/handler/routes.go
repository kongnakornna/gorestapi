package handler

import (
	"github.com/go-chi/chi/v5"

	monConfig "icmongolang/internal/monitoring/config"
)

// MapMonitoringRoutes registers all monitoring endpoints on the given router.
func MapMonitoringRoutes(router *chi.Mux, cfg *monConfig.MonitoringConfig) {
	if cfg == nil || !cfg.Enabled {
		return
	}

	router.Route("/monitoring", func(r chi.Router) {
		r.Get("/", HealthHandler)
		if cfg.MetricsEnabled {
			r.Handle("/metrics", MetricsHandler())
		}
		r.Get("/health", HealthHandler) 
		if cfg.SystemMetricsEnabled {
			r.Get("/system", SystemStatsHandler)
		}
	})
}