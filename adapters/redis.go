package adapters

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type (
	RedisAdapter struct {
		AbstractAdapter
		client *redis.Client
	}
)

func NewRedisAdapter(dep *DepContainer) (*RedisAdapter, error) {
	redisURLOpts, err := redis.ParseURL(dep.RedisURL)
	if err != nil {
		return nil, err
	}

	return &RedisAdapter{
		client: redis.NewClient(redisURLOpts),
	}, nil
}

func (adapter *RedisAdapter) Get(key string) (string, error) {
	res, err := adapter.client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}

	return res, nil
}

func (adapter *RedisAdapter) Set(key string, data interface{}, exp time.Duration) error {
	jsonData, err := json.Marshal(&data)
	if err != nil {
		return err
	}

	return adapter.client.SetNX(context.Background(), key, jsonData, exp).Err()
}

func (adapter *RedisAdapter) Delete(key string) error {
	return adapter.client.Del(context.Background(), key).Err()
}
