package users

import (
	"github.com/kongnakornna/gorestapi/internal"
	"github.com/kongnakornna/gorestapi/internal/models"
)

type UserRedisRepository interface {
	internal.RedisRepository[models.User]
}
