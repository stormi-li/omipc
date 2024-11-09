package omipc

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func NewClient(opts *redis.Options) *Client {
	return &Client{
		redisClient: redis.NewClient(opts),
		ctx:         context.Background(),
	}
}
