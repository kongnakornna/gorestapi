// Package metrics เก็บ system stats (CPU, RAM, Network)
// Package metrics collects system stats (CPU, RAM, Network)
package metrics

import (
	"context"
	"log/slog"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

var (
	// cpuUsagePercent เปอร์เซ็นต์การใช้ CPU
	// cpuUsagePercent CPU usage percentage
	cpuUsagePercent = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "system_cpu_usage_percent",
			Help: "CPU usage percentage",
		},
	)
	// memUsagePercent เปอร์เซ็นต์การใช้ RAM
	// memUsagePercent RAM usage percentage
	memUsagePercent = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "system_memory_usage_percent",
			Help: "Memory usage percentage",
		},
	)
	// netBytesRecv จำนวน bytes ที่รับจาก network
	// netBytesRecv total bytes received
	netBytesRecv = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "system_network_receive_bytes_total",
			Help: "Total network bytes received",
		},
	)
)

// StartSystemMetricsCollector เริ่ม goroutine พื้นหลังสำหรับเก็บ system stats ทุก 30 วินาที
// StartSystemMetricsCollector starts a background goroutine to collect system stats every 30s
func StartSystemMetricsCollector(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)
	go func() {
		for {
			select {
			case <-ctx.Done():
				slog.Info("Stopping system metrics collector")
				return
			case <-ticker.C:
				collect()
			}
		}
	}()
	slog.Info("System metrics collector started")
}

func collect() {
	// CPU
	percent, err := cpu.Percent(0, false)
	if err == nil && len(percent) > 0 {
		cpuUsagePercent.Set(percent[0])
	}

	// Memory
	memStat, err := mem.VirtualMemory()
	if err == nil {
		memUsagePercent.Set(memStat.UsedPercent)
	}

	// Network (since last call - ใช้ counters แบบ cumulative)
	netIO, err := net.IOCounters(false)
	if err == nil && len(netIO) > 0 {
		// ไทย: ใช้ Add เพื่อสะสมค่า (เพราะเป็น cumulative counter)
		// English: Use Add to accumulate (cumulative counter)
		netBytesRecv.Add(float64(netIO[0].BytesRecv))
	}
}