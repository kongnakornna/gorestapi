package processor

import (
	"github.com/hibiken/asynq"
	"github.com/kongnakornna/gorestapi/config"
	"github.com/kongnakornna/gorestapi/pkg/logger"
)

type RedisTaskProcessor struct {
	Server *asynq.Server
	Cfg    *config.Config
	Logger logger.Logger
}

func NewRedisTaskProcessor(server *asynq.Server, cfg *config.Config, logger logger.Logger) RedisTaskProcessor {
	return RedisTaskProcessor{Server: server, Cfg: cfg, Logger: logger}
}
