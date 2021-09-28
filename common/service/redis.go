package common

import (
	"context"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisService interface {
	Set(key string, data interface{}, ttl time.Duration) error
	Get(key string) (interface{}, error)
}

type redisService struct {
	ctx    context.Context
	client *redis.Client
}

type redisConfig struct {
	host      string
	passsword string
}

func createClient(host string, password string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
	})
	return rdb
}

func getConfig() redisConfig {
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		host = "localhost:6379"
	}
	password := os.Getenv("REDIS_PASSWORD")
	if password == "" {
		password = ""
	}
	return redisConfig{
		host,
		password,
	}
}

func NewRedis() RedisService {
	config := getConfig()
	return &redisService{
		ctx:    context.Background(),
		client: createClient(config.host, config.passsword),
	}
}

func (r *redisService) Get(key string) (data interface{}, err error) {
	val, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		return nil, err
	} else if err == redis.Nil {
		return nil, nil
	}
	return val, nil
}

func (r *redisService) Set(key string, data interface{}, ttl time.Duration) (err error) {
	cmd := r.client.Set(r.ctx, key, data, ttl)
	err = cmd.Err()
	return
}
