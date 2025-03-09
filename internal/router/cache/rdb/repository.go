package rdb

import (
	"time"

	"github.com/Ayeye11/AuthCache/internal/common/types"
	"github.com/redis/go-redis/v9"
)

type RedisCache interface {
	// Roles
	SaveRole(role *types.Role, perms []*types.Permission) error
	GetRole(roleID int) (*types.Role, []*types.Permission, error)
}

func NewCache(client *redis.Client, ttl time.Duration) RedisCache {
	return &cache{client, ttl}
}

type cache struct {
	client *redis.Client
	ttl    time.Duration
}
