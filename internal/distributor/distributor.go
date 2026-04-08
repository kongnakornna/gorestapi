package distributor

import (
	"icmongolang/config"
	"icmongolang/pkg/logger"

	"github.com/hibiken/asynq"
)

type RedisTaskDistributor struct {
	RedisClient *asynq.Client
	Cfg         *config.Config
	Logger      logger.Logger
}

func NewRedisClient(cfg *config.Config) *asynq.Client {
	redisOpt := asynq.RedisClientOpt{
		Addr: cfg.TaskRedis.Addr,
		DB:   cfg.TaskRedis.Db,
	}

	return asynq.NewClient(redisOpt)
}

func NewRedisTaskDistributor(redisClient *asynq.Client, cfg *config.Config, loggger logger.Logger) RedisTaskDistributor {
	return RedisTaskDistributor{
		RedisClient: redisClient,
		Cfg:         cfg,
		Logger:      loggger,
	}
}
