package alert

import (
	"crypto/tls"
	"fmt"
	"log/slog"
	"net/smtp"

	monConfig "icmongolang/internal/monitoring/config"
)

// SendAlert ส่งอีเมล alert (ไม่ block การทำงานหลัก)
func SendAlert(subject, body string) {
	cfg := monConfig.LoadMonitoringConfig()
	if cfg.AlertSMTPHost == "" || !cfg.AlertEnabled {
		slog.Warn("SMTP not configured or alert disabled, skipping email alert")
		return
	}

	go func() {
		to := []string{cfg.AlertToEmail}
		msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
			cfg.AlertFromEmail, cfg.AlertToEmail, subject, body)

		auth := smtp.PlainAuth("", cfg.AlertSMTPUser, cfg.AlertSMTPPass, cfg.AlertSMTPHost)
		addr := fmt.Sprintf("%s:%s", cfg.AlertSMTPHost, cfg.AlertSMTPPort)

		conn, err := tls.Dial("tcp", addr, nil)
		if err != nil {
			slog.Error("Alert email failed (TLS dial)", "error", err)
			return
		}
		client, err := smtp.NewClient(conn, cfg.AlertSMTPHost)
		if err != nil {
			slog.Error("Alert email failed (client)", "error", err)
			return
		}
		defer client.Close()
		if err = client.Auth(auth); err != nil {
			slog.Error("Alert email auth failed", "error", err)
			return
		}
		if err = client.Mail(cfg.AlertFromEmail); err != nil {
			slog.Error("Alert email MAIL FROM failed", "error", err)
			return
		}
		if err = client.Rcpt(cfg.AlertToEmail); err != nil {
			slog.Error("Alert email RCPT TO failed", "error", err)
			return
		}
		w, err := client.Data()
		if err != nil {
			slog.Error("Alert email DATA failed", "error", err)
			return
		}
		_, err = w.Write([]byte(msg))
		if err != nil {
			slog.Error("Alert email write failed", "error", err)
			return
		}
		err = w.Close()
		if err != nil {
			slog.Error("Alert email close failed", "error", err)
			return
		}
		slog.Info("Alert email sent", "subject", subject, "to", cfg.AlertToEmail)
	}()
}