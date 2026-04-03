package repository

import (
	"gorestapi//internal/models"
	"gorestapi//internal/repository"
	"gorestapi//internal/users"

	"github.com/redis/go-redis/v9"
)

type UserRedisRepo struct {
	repository.RedisRepo[models.User]
}

func CreateUserRedisRepository(redisClient *redis.Client) users.UserRedisRepository {
	return &UserRedisRepo{
		RedisRepo: repository.RedisRepo[models.User](repository.CreateRedisRepo[models.User](redisClient)),
	}
}
