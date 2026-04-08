package postgres

import (
	"fmt"

	"icmongolang/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPsqlDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		cfg.Postgres.Host, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Dbname, cfg.Postgres.Port,
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		// ✅ บรรทัดนี้สำคัญที่สุด – ปิด FK อัตโนมัติ
		DisableForeignKeyConstraintWhenMigrating: true,
	})
}
