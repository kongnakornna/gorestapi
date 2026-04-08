
// Package handler จัดการ REST endpoints สำหรับ monitoring
package handler

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

var startTime = time.Now()

// MetricsHandler ส่ง Prometheus metrics
// @Summary      Get Prometheus metrics
// @Description  ส่ง metrics ในรูปแบบที่ Prometheus เข้าใจ (plain text)
// @Tags         Monitoring
// @Produce      plain
// @Success      200 {string} string "Prometheus metrics"
// @Router       /monitoring/metrics [get]
func MetricsHandler() http.Handler {
	return promhttp.Handler()
}

// HealthResponse โครงสร้างตอบกลับ health check
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Uptime    string    `json:"uptime"`
}

// HealthHandler ตรวจสอบว่า service ทำงานปกติ
// @Summary      Health check
// @Description  ตรวจสอบสถานะของ service (alive)
// @Tags         Monitoring
// @Produce      json
// @Success      200 {object} HealthResponse "service is healthy"
// @Router       /monitoring/health [get]
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
// @Summary      System statistics
// @Description  ดึงข้อมูลการใช้ CPU, RAM, จำนวน goroutine
// @Tags         Monitoring
// @Produce      json
// @Success      200 {object} map[string]interface{} "system stats"
// @Router       /monitoring/system [get]
func SystemStatsHandler(w http.ResponseWriter, r *http.Request) {
	cpuPercent, _ := cpu.Percent(0, false)
	memStat, _ := mem.VirtualMemory()

	stats := map[string]interface{}{
		"cpu_percent":   cpuPercent[0],
		"ram_percent":   memStat.UsedPercent,
		"ram_used_mb":   memStat.Used / 1024 / 1024,
		"ram_total_mb":  memStat.Total / 1024 / 1024,
		"goroutines":    runtime.NumGoroutine(),
		"num_cpu":       runtime.NumCPU(),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}