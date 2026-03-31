package cmd

import (
	"github.com/kongnakornna/gorestapi/config"
	"github.com/kongnakornna/gorestapi/internal/models"
	"github.com/kongnakornna/gorestapi/pkg/db/postgres"
	"github.com/kongnakornna/gorestapi/pkg/logger"
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
		} else {
			appLogger.Infof("Postgres connected successfully - เชื่อมต่อสำเร็จ..")
		}

		err = Migrate(psqlDB)

		if err != nil {
			appLogger.Info("Can not migrate data - ไม่สามารถ migrate ข้อมูลได้")
		} else {
			appLogger.Info("Data migrated - Data migrated สำเร็จ")
		}
	},
}

func Migrate(db *gorm.DB) error {
	var migrationModels = []interface{}{&models.User{}, &models.Item{}}

	err := db.AutoMigrate(migrationModels...)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	RootCmd.AddCommand(migrateCmd)
}
