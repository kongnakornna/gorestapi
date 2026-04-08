package config

import (
	icmConfig "icmongolang/config"
)

// MonitoringConfig โครงสร้างเก็บค่าการเปิด-ปิดของแต่ละ component
type MonitoringConfig struct {
	Enabled              bool
	MetricsEnabled       bool
	TracingEnabled       bool
	SentryEnabled        bool
	AlertEnabled         bool
	SystemMetricsEnabled bool
	SentryDSN            string
	JaegerEndpoint       string
	AlertSMTPHost        string
	AlertSMTPPort        string
	AlertSMTPUser        string
	AlertSMTPPass        string
	AlertFromEmail       string
	AlertToEmail         string
}

// LoadMonitoringConfig อ่านค่าจาก global config ของโปรเจกต์
func LoadMonitoringConfig() MonitoringConfig {
	cfg := icmConfig.GetCfg()
	mon := cfg.Monitoring

	if !mon.Enabled {
		return MonitoringConfig{Enabled: false}
	}

	return MonitoringConfig{
		Enabled:              mon.Enabled,
		MetricsEnabled:       mon.MetricsEnabled,
		TracingEnabled:       mon.TracingEnabled,
		SentryEnabled:        mon.SentryEnabled,
		AlertEnabled:         mon.AlertEnabled,
		SystemMetricsEnabled: mon.SystemMetricsEnabled,
		SentryDSN:            mon.SentryDSN,
		JaegerEndpoint:       mon.JaegerEndpoint,
		AlertSMTPHost:        mon.AlertSMTPHost,
		AlertSMTPPort:        mon.AlertSMTPPort,
		AlertSMTPUser:        mon.AlertSMTPUser,
		AlertSMTPPass:        mon.AlertSMTPPass,
		AlertFromEmail:       mon.AlertFromEmail,
		AlertToEmail:         mon.AlertToEmail,
	}
}