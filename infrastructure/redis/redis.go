package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/Ayeye11/se-thr/config"
	"github.com/redis/go-redis/v9"
)

type redisDatabase struct {
	client *redis.Client
	ttl    time.Duration
}

func NewRedisDB(cfg config.ConfigRedis) (*redisDatabase, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.Db,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return &redisDatabase{client, cfg.TTL}, nil
}

func (r *redisDatabase) GetClient() *redis.Client {
	return r.client
}

func (r *redisDatabase) GetTTL() time.Duration {
	return r.ttl
}

func (r *redisDatabase) Close() error {
	return r.client.Close()
}
