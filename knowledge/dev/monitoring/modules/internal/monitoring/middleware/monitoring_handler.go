// Package handler จัดการ REST endpoints สำหรับ monitoring
// Package handler serves REST endpoints for monitoring
package handler

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"

	"icmongolang/internal/monitoring/alert"
	monmetrics "icmongolang/internal/monitoring/metrics"
)

// HealthResponse โครงสร้างตอบกลับ health check
// HealthResponse struct for health check response
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Uptime    string    `json:"uptime"`
}

// MetricsHandler ส่ง Prometheus metrics (ใช้ promhttp)
// MetricsHandler serves Prometheus metrics
func MetricsHandler() http.Handler {
	return promhttp.Handler()
}

// HealthHandler ตรวจสอบว่า service ทำงานปกติ
// HealthHandler checks if service is alive
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	resp := HealthResponse{
		Status:    "ok",
		Timestamp: time.Now(),
		Uptime:    time.Since(startTime).String(),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// SystemStatsHandler คืนค่า CPU, RAM, Goroutines
// SystemStatsHandler returns CPU, RAM, Goroutines
func SystemStatsHandler(w http.ResponseWriter, r *http.Request) {
	cpuPercent, _ := cpu.Percent(0, false)
	memStat, _ := mem.VirtualMemory()

	stats := map[string]interface{}{
		"cpu_percent":    cpuPercent[0],
		"ram_percent":    memStat.UsedPercent,
		"ram_used_mb":    memStat.Used / 1024 / 1024,
		"ram_total_mb":   memStat.Total / 1024 / 1024,
		"goroutines":     runtime.NumGoroutine(),
		"num_cpu":        runtime.NumCPU(),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

var startTime = time.Now()