package processor

import (
	"gorestapi//config"
	"gorestapi//pkg/logger"

	"github.com/hibiken/asynq"
)

type RedisTaskProcessor struct {
	Server *asynq.Server
	Cfg    *config.Config
	Logger logger.Logger
}

func NewRedisTaskProcessor(server *asynq.Server, cfg *config.Config, logger logger.Logger) RedisTaskProcessor {
	return RedisTaskProcessor{Server: server, Cfg: cfg, Logger: logger}
}
