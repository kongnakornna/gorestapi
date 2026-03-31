package middleware

import (
	"github.com/kongnakornna/gorestapi/config"
	"github.com/kongnakornna/gorestapi/internal/users"
	"github.com/kongnakornna/gorestapi/pkg/logger"
)

type MiddlewareManager struct {
	cfg     *config.Config
	logger  logger.Logger
	usersUC users.UserUseCaseI
}

func CreateMiddlewareManager(cfg *config.Config, logger logger.Logger, usersUC users.UserUseCaseI) *MiddlewareManager {
	return &MiddlewareManager{
		cfg:     cfg,
		logger:  logger,
		usersUC: usersUC,
	}
}
