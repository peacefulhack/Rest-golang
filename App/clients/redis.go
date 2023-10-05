package clients

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisClient struct {
	Redis *redis.Client
}
type RedisClientRepo interface {
	Set(key string, val interface{}, duration time.Duration) error
	Get(key string) (string, bool, error)
	Del(key string) error
}

func NewRedisClient(reds *redis.Client) *RedisClient {
	return &RedisClient{Redis: reds}
}

func (m *RedisClient) Set(key string, val interface{}, duration time.Duration) error {
	err := m.Redis.Set(context.Background(), key, val, duration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (m *RedisClient) Get(key string) (string, bool, error) {
	value, err := m.Redis.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return "", false, nil
	} else if err != nil {
		return "", false, err
	} else {
		return value, true, nil
	}
}

func (m *RedisClient) Del(key string) error {
	_, err := m.Redis.Del(context.Background(), key).Result()
	if err != nil {
		return err
	}
	return nil
}
