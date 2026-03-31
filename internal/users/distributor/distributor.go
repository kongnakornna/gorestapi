package distributor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/kongnakornna/gorestapi/config"
	"github.com/kongnakornna/gorestapi/internal/distributor"
	"github.com/kongnakornna/gorestapi/internal/users"
	"github.com/kongnakornna/gorestapi/pkg/logger"
)

type userRedisTaskDistributor struct {
	distributor.RedisTaskDistributor
}

func NewUserRedisTaskDistributor(redisClient *asynq.Client, cfg *config.Config, loggger logger.Logger) users.UserRedisTaskDistributor {
	return &userRedisTaskDistributor{
		RedisTaskDistributor: distributor.NewRedisTaskDistributor(redisClient, cfg, loggger),
	}
}

func (distributor *userRedisTaskDistributor) DistributeTaskSendEmail(ctx context.Context, payload *users.PayloadSendEmail, opts ...asynq.Option) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload %w", err)
	}

	task := asynq.NewTask(users.TaskSendEmail, jsonPayload, opts...)

	info, err := distributor.RedisClient.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %w", err)
	}

	distributor.Logger.Infof("Type: %v, Queue: %v, Max-Retry: %v, Msg: queued task", task.Type(), info.Queue, info.MaxRetry)

	return nil
}
