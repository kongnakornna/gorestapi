package cmd

import (
	"context"

	"github.com/kongnakornna/gorestapi/config"
	"github.com/kongnakornna/gorestapi/internal/distributor"
	userDistributor "github.com/kongnakornna/gorestapi/internal/users/distributor"
	userRepository "github.com/kongnakornna/gorestapi/internal/users/repository"
	userUseCase "github.com/kongnakornna/gorestapi/internal/users/usecase"
	"github.com/kongnakornna/gorestapi/pkg/db/postgres"
	"github.com/kongnakornna/gorestapi/pkg/db/redis"
	"github.com/kongnakornna/gorestapi/pkg/logger"
	"github.com/spf13/cobra" 
)

var initDataCmd = &cobra.Command{
	Use:   "initdata",
	Short: "Init data",
	Long:  "Init data",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetCfg()

		appLogger := logger.NewApiLogger(cfg)
		appLogger.InitLogger()
		appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode)

		psqlDB, err := postgres.NewPsqlDB(cfg)
		if err != nil {
			appLogger.Fatalf("เชื่อมต่อไม่สำเร็จ - Postgresql init: %s", err)
		} else {
			appLogger.Infof("Postgres connected เชื่อมต่อสำเร็จ..")
		}

		redisClient := redis.NewRedis(cfg)

		taskRedisClient := distributor.NewRedisClient(cfg)

		// Repository
		userPgRepo := userRepository.CreateUserPgRepository(psqlDB)
		userRedisRepo := userRepository.CreateUserRedisRepository(redisClient)

		// Distributor
		userRedisTaskDistributor := userDistributor.NewUserRedisTaskDistributor(taskRedisClient, cfg, appLogger)

		// UseCase
		userUC := userUseCase.CreateUserUseCaseI(userPgRepo, userRedisRepo, userRedisTaskDistributor, cfg, appLogger)

		// Create super user if not exists
		isCreated, _ := userUC.CreateSuperUserIfNotExist(context.Background())

		if !isCreated {
			appLogger.Info("Super user is exists, skip create")
		} else {
			appLogger.Info("Created super user")
		}
	},
}

func init() {
	RootCmd.AddCommand(initDataCmd)
}