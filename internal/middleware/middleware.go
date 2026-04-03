package middleware

import (
	"icmongolang/config"
	"icmongolang/internal/users"
	"icmongolang/pkg/logger"
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
