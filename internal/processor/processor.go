package processor

import (
	"gorestapi/config"
	"gorestapi/pkg/logger"

	"github.com/hibiken/asynq"
)

// RedisTaskDistributor is a base distributor for async tasks using Redis.
type RedisTaskDistributor struct {
	RedisClient *asynq.Client
	Cfg         *config.Config
	Logger      logger.Logger
}

// NewRedisTaskDistributor creates a new base distributor.
func NewRedisTaskDistributor(redisClient *asynq.Client, cfg *config.Config, logger logger.Logger) RedisTaskDistributor {
	return RedisTaskDistributor{
		RedisClient: redisClient,
		Cfg:         cfg,
		Logger:      logger,
	}
}

// RedisTaskProcessor is a base processor for async tasks.
type RedisTaskProcessor struct {
	Server *asynq.Server
	Cfg    *config.Config
	Logger logger.Logger
}

// NewRedisTaskProcessor creates a new base processor.
func NewRedisTaskProcessor(server *asynq.Server, cfg *config.Config, logger logger.Logger) RedisTaskProcessor {
	return RedisTaskProcessor{
		Server: server,
		Cfg:    cfg,
		Logger: logger,
	}
}