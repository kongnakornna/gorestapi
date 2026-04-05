package cmd

import (
	"icmongolang/config"
	"icmongolang/internal/models"
	"icmongolang/pkg/db/postgres"
	"icmongolang/pkg/logger"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate data",
	Long:  "Migrate data",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetCfg()

		appLogger := logger.NewApiLogger(cfg)
		appLogger.InitLogger()
		appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode)

		psqlDB, err := postgres.NewPsqlDB(cfg)
		if err != nil {
			appLogger.Fatalf("เชื่อมต่อไม่สำเร็จ - Postgresql init: %s", err)
		}
		appLogger.Infof("Postgres connected successfully - เชื่อมต่อสำเร็จ..")

		// Create necessary extensions
		if err := createExtensions(psqlDB); err != nil {
			appLogger.Warnf("Failed to create extensions: %v", err)
		}

		// Run migration
		if err := Migrate(psqlDB); err != nil {
			appLogger.Fatal("Migrate ข้อมูลไม่สำเร็จ: ", err)
		}
		appLogger.Info("Migrate ข้อมูลสำเร็จ")
	},
}

// createExtensions creates required PostgreSQL extensions
func createExtensions(db *gorm.DB) error {
	extensions := []string{
		"CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"",
		"CREATE EXTENSION IF NOT EXISTS \"pgcrypto\"",
	}
	for _, ext := range extensions {
		if err := db.Exec(ext).Error; err != nil {
			return err
		}
	}
	return nil
}

// Migrate performs auto migration for all models
// cmd/migrate.go
func Migrate(db *gorm.DB) error {
    // Remove the session block – it’s not needed and caused the error
    migrationModels := []interface{}{
        &models.ActivityLog{},
        &models.CommandLog{},
        &models.DeviceAlert{},
        &models.DeviceConfig{},
        &models.DeviceStatus{},
        &models.IotData{},
        &models.NotiNotificationLog{},
        &models.NotiNotificationRule{},
        &models.NotiNotificationType{},
        &models.NotiNotification{},
        &models.NotificationDevice{},
        &models.NotificationGroup{},
        &models.NotificationGroupsDevicesNotificationDevice{},
        &models.NotificationLog{},
        &models.NotificationType{},
        &models.SdActivityLog{},
        &models.SdActivityTypeLog{},
        &models.SdAdminAccessMenu{},
        &models.SdAirControl{},
        &models.SdAirControlDeviceMap{},
        &models.SdAirControlLog{},
        &models.SdAirMod{},
        &models.SdAirModDeviceMap{},
        &models.SdAirPeriod{},
        &models.SdAirPeriodDeviceMap{},
        &models.SdAirSettingWarning{},
        &models.SdAirSettingWarningDeviceMap{},
        &models.SdAirWarning{},
        &models.SdAirWarningDeviceMap{},
        &models.SdAlarmProcessLog{},
        &models.SdAlarmProcessLogEmail{},
        &models.SdAlarmProcessLogLine{},
        &models.SdAlarmProcessLogMqtt{},
        &models.SdAlarmProcessLogSms{},
        &models.SdAlarmProcessLogTelegram{},
        &models.SdAlarmProcessLogTemp{},
        &models.SdApiKey{},
        &models.SdAuditLog{},
        &models.SdChannelTemplate{},
        &models.SdDashboardConfig{},
        &models.SdDeviceCategory{},
        &models.SdDeviceGroup{},
        &models.SdDeviceLog{},
        &models.SdDeviceMember{},
        &models.SdDeviceNotificationConfig{},
        &models.SdDeviceSchedule{},
        &models.SdDeviceStatusHistory{},
        &models.SdGroupNotificationConfig{},
        &models.SdIotAlarmDevice{},
        &models.SdIotAlarmDeviceEvent{},
        &models.SdIotApi{},
        &models.SdIotDevice{},
        &models.SdIotDeviceAction{},
        &models.SdIotDeviceActionLog{},
        &models.SdIotDeviceActionUser{},
        &models.SdIotDeviceAlarmAction{},
        &models.SdIotDeviceType{},
        &models.SdIotEmail{},
        &models.SdIotGroup{},
        &models.SdIotHost{},
        &models.SdIotInfluxdb{},
        &models.SdIotLine{},
        &models.SdIotLocation{},
        &models.SdIotMqtt{},
        &models.SdIotNodered{},
        &models.SdIotSchedule{},
        &models.SdIotScheduleDevice{},
        &models.SdIotSensor{},
        &models.SdIotSetting{},
        &models.SdIotSms{},
        &models.SdIotTelegram{},
        &models.SdIotToken{},
        &models.SdIotType{},
        &models.SdModuleLog{},
        &models.SdMqttHost{},
        &models.SdMqttLog{},
        &models.SdNotificationChannel{},
        &models.SdNotificationCondition{},
        &models.SdNotificationLog{},
        &models.SdNotificationType{},
        &models.SdReportData{},
        &models.SdScheduleProcessLog{},
        &models.SdSensorData{},
        &models.SdSystemSetting{},
        &models.SdUser{},
        &models.SdUserAccessMenu{},
        &models.SdUserFile{},
        &models.SdUserLog{},
        &models.SdUserLogType{},
        &models.SdUserRole{},
        &models.SdUserRolesAccess{},
        &models.SdUserRolesPermision{},
        &models.Tnb{},
    }

    // Run AutoMigrate directly – foreign keys are disabled by the config
    if err := db.AutoMigrate(migrationModels...); err != nil {
        return err
    }

    return nil
}

// Optional: Add foreign keys manually after migration to avoid ordering issues
func addForeignKeys(db *gorm.DB) error {
	// Example: fix foreign key for sd_activity_log -> sd_activity_type_log
	// This is optional and depends on your schema requirements
	sqls := []string{
		// "ALTER TABLE sd_activity_log ADD CONSTRAINT fk_sd_activity_log_type FOREIGN KEY (type_id) REFERENCES sd_activity_type_log(typeId) ON DELETE SET NULL",
		// "ALTER TABLE sd_activity_log ADD CONSTRAINT fk_sd_activity_log_module FOREIGN KEY (modules_id) REFERENCES sd_module_log(id) ON DELETE SET NULL",
	}
	for _, sql := range sqls {
		if err := db.Exec(sql).Error; err != nil {
			// Ignore errors if constraint already exists
			continue
		}
	}
	return nil
}

func init() {
	RootCmd.AddCommand(migrateCmd)
}

// go run cmd/api/main.go migrate