package users

import (
	"gorestapi//internal"
	"gorestapi//internal/models"
)

type UserRedisRepository interface {
	internal.RedisRepository[models.User]
}
