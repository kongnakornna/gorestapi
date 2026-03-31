package repository

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/kongnakornna/gorestapi/internal"
	"github.com/kongnakornna/gorestapi/pkg/httpErrors"
	"github.com/redis/go-redis/v9"
)

type RedisRepo[M any] struct {
	RedisClient *redis.Client
}

/*************  ✨ Windsurf Command ⭐  *************/
// CreateRedisRepo returns a new RedisRepo instance with the given Redis client.
//
// This function is a generic constructor for RedisRepo instances. It
// is intended to be used with any type that satisfies the
// constraints of the RedisRepo type parameter.
//
// The returned RedisRepo instance is ready to use, with the given
// Redis client set as its RedisClient field.
//
// Note that the returned RedisRepo instance is not thread-safe. If
// you need to use the same RedisRepo instance from multiple goroutines,
// you will need to ensure that access to the instance is properly
// synchronized.
// CreateRedisRepo สร้าง instance ใหม่ของ RedisRepo พร้อมกับ Redis client ที่กำหนด
// ฟังก์ชันนี้เป็น constructor ทั่วไปสำหรับสร้าง instance ของ RedisRepo
// โดยสามารถใช้งานได้กับ type ใดๆ ที่เป็นไปตามข้อจำกัด (constraints) 
// ของ type parameter ใน RedisRepo
//
// instance RedisRepo ที่คืนกลับมาพร้อมสำหรับการใช้งาน โดยมี Redis client 
// ที่กำหนดถูกตั้งค่าในฟิลด์ RedisClient แล้ว
//
// หมายเหตุ: instance RedisRepo ที่คืนกลับมานี้ไม่ปลอดภัยสำหรับการใช้งานแบบ concurrent
// หากคุณต้องการใช้ RedisRepo instance เดียวกันจากหลาย goroutine
// คุณจะต้องตรวจสอบให้แน่ใจว่าการเข้าถึง instance นั้นมีการจัดการการซิงโครไนซ์อย่างเหมาะสม
/*******  c111f467-a265-443f-969f-86eba31cbe94  *******/
func CreateRedisRepo[M any](redisClient *redis.Client) RedisRepo[M] {
	return RedisRepo[M]{RedisClient: redisClient}
}

/*************  ✨ Windsurf Command ⭐  *************/
// CreateRedisRepository returns a new RedisRepository instance with the given Redis client.
//
// This function is a generic constructor for RedisRepository instances. It
// is intended to be used with any type that satisfies the
// constraints of the RedisRepository type parameter.
//
// The returned RedisRepository instance is ready to use, with the given
// Redis client set as its RedisClient field.
//
// Note that the returned RedisRepository instance is not thread-safe. If
// you need to use the same RedisRepository instance from multiple goroutines,
// you will need to ensure that access to the instance is properly
// synchronized.
/*******  7a2d5453-9a92-4129-a7ac-8d83e9d34994  *******/
func CreateRedisRepository[M any](redisClient *redis.Client) internal.RedisRepository[M] {
	return &RedisRepo[M]{RedisClient: redisClient}
}

/*************  ✨ Windsurf Command ⭐  *************/
// Create sets the value of the given key to the JSON encoded value of
// exp in the Redis database with the given TTL in seconds.
//
// The function returns an error if the JSON encoding of exp fails,
// or if the Redis SET command fails.
//
// Note that the returned error may be wrapped with an
// httpErrors.Error, if the error is related to JSON encoding or
// Redis communication.
/*******  3338ab83-d62e-4cf6-ae12-22ee6a87ef27  *******/
func (r *RedisRepo[M]) Create(ctx context.Context, key string, exp *M, seconds int) error {
	objBytes, err := json.Marshal(exp)
	if err != nil {
		return httpErrors.ErrJson(err)
	}

	if err = r.RedisClient.Set(ctx, key, objBytes, time.Second*time.Duration(seconds)).Err(); err != nil {
		// TODO: Using httpErrors
		return err
	}
	return nil
}

/*************  ✨ Windsurf Command ⭐  *************/
// Get retrieves the value of the given key from the Redis database and
// unmarshals it into obj of type M.
//
// If the key does not exist in the Redis database, the function returns
// nil for obj and nil for err.
//
// If the JSON unmarshaling of objBytes fails, the function returns an error
// wrapped with an httpErrors.Error.
//
// Note that the returned error may be wrapped with an
// httpErrors.Error, if the error is related to JSON encoding or
// Redis communication.
/*******  367846a6-dcaf-4696-8f8a-5dcce38c96f0  *******/
func (r *RedisRepo[M]) Get(ctx context.Context, key string) (*M, error) {
	objBytes, err := r.RedisClient.Get(ctx, key).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, err
	}

	var obj M

	if err = json.Unmarshal(objBytes, &obj); err != nil {
		return nil, httpErrors.ErrJson(err)
	}

	return &obj, nil
}

/*************  ✨ Windsurf Command ⭐  *************/
// Delete removes the key from the Redis database.
//
// If the key does not exist in the Redis database, the function returns
// nil for err.
//
// If the Redis DEL command fails, the function returns an error
// wrapped with an httpErrors.Error.
//
// Note that the returned error may be wrapped with an
// httpErrors.Error, if the error is related to Redis communication.
/*******  0eb6315f-4ddc-478e-9e3b-d215d5243850  *******/
func (r *RedisRepo[M]) Delete(ctx context.Context, key string) error {
	if err := r.RedisClient.Del(ctx, key).Err(); err != nil {
		if errors.Is(err, redis.Nil) {
			return nil
		}
		// TODO: Using httpErrors
		return err
	}
	return nil
}

/*************  ✨ Windsurf Command ⭐  *************/
// Sadd adds the given value to the Redis set with the given key.
//
// If the value is already in the set, the function returns nil for err.
//
// If the Redis SADD command fails, the function returns an error
// wrapped with an httpErrors.Error.
//
// Note that the returned error may be wrapped with an
// httpErrors.Error, if the error is related to Redis communication.
/*******  460b09a0-018c-41a6-9645-a6b1f395d74a  *******/
func (r *RedisRepo[M]) Sadd(ctx context.Context, key string, value string) error {
	if err := r.RedisClient.SAdd(ctx, key, value).Err(); err != nil {
		return err
	}
	return nil
}

/*************  ✨ Windsurf Command ⭐  *************/
// Sadds adds the given values to the Redis set with the given key.
//
// If any of the values are already in the set, the function returns nil for err.
//
// If the Redis SADD command fails, the function returns an error
// wrapped with an httpErrors.Error.
//
// Note that the returned error may be wrapped with an
// httpErrors.Error, if the error is related to Redis communication.
/*******  2a6ac737-847a-48fe-92c2-37493ba511e5  *******/
func (r *RedisRepo[M]) Sadds(ctx context.Context, key string, values []string) error {
	if err := r.RedisClient.SAdd(ctx, key, values).Err(); err != nil {
		return err
	}
	return nil
}

/*************  ✨ Windsurf Command ⭐  *************/
/*
Srem removes the given value from the Redis set with the given key.

If the value is not in the set, the function returns nil for err.

If the Redis SREM command fails, the function returns an error
wrapped with an httpErrors.Error.

Note that the returned error may be wrapped with an
httpErrors.Error, if the error is related to Redis communication.
*/
/*******  7fa51355-a7fd-41de-bafe-7b4bbb79b378  *******/
func (r *RedisRepo[M]) Srem(ctx context.Context, key string, value string) error {
	if err := r.RedisClient.SRem(ctx, key, value).Err(); err != nil {
		return err
	}
	return nil
}

/*************  ✨ Windsurf Command ⭐  *************/
/*
SIsMember checks if the given value is in the Redis set with the given key.

If the Redis SISMEMBER command fails, the function returns false for the result
and an error wrapped with an httpErrors.Error.

Note that the returned error may be wrapped with an
httpErrors.Error, if the error is related to Redis communication.
*/
/*******  63397f97-2cdb-4199-bdbb-aa8150efcfd7  *******/
func (r *RedisRepo[M]) SIsMember(ctx context.Context, key string, value string) (bool, error) {
	result := r.RedisClient.SIsMember(ctx, key, value)
	if result.Err() != nil {
		return false, result.Err()
	}
	return result.Val(), nil
}

// func (r *RedisRepo[M]) SMembers(ctx context.Context, key string) ([]string, error) {
// 	result := r.RedisClient.SPop(ctx, key)
// 	if result.Err() != nil {
// 		return nil, result.Err()
// 	}
// 	return result.Val(), nil
// }
