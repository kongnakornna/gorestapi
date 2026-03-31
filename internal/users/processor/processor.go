package processor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/kongnakornna/gorestapi/config"
	"github.com/kongnakornna/gorestapi/internal/processor"
	"github.com/kongnakornna/gorestapi/internal/users"
	"github.com/kongnakornna/gorestapi/pkg/logger"
	"github.com/kongnakornna/gorestapi/pkg/sendEmail"
)

type userRedisTaskProcessor struct {
	processor.RedisTaskProcessor
	emailSender sendEmail.EmailSender
}

func NewUserRedisTaskProcessor(server *asynq.Server, cfg *config.Config, logger logger.Logger, emailSender sendEmail.EmailSender) users.UserRedisTaskProcessor {
	return &userRedisTaskProcessor{
		RedisTaskProcessor: processor.NewRedisTaskProcessor(server, cfg, logger),
		emailSender:        emailSender,
	}
}

func (processor *userRedisTaskProcessor) ProcessTaskSendEmail(ctx context.Context, task *asynq.Task) error {
	var payload users.PayloadSendEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	if err := processor.emailSender.SendEmail(
		ctx,
		payload.From,
		payload.To,
		payload.Subject,
		payload.BodyHtml,
		payload.BodyPlain,
	); err != nil {
		return fmt.Errorf("failed to send verify email: %w", err)
	}

	processor.Logger.Infof("Type: %v, Msg: email sended", task.Type())

	return nil
}
