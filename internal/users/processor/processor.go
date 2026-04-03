package processor

import (
	"context"
	"encoding/json"
	"fmt"

	"gorestapi/config"
	"gorestapi/internal/processor"
	"gorestapi/internal/users"
	"gorestapi/pkg/logger"
	"gorestapi/pkg/sendEmail"

	"github.com/hibiken/asynq"
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
