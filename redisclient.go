package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisBenchmarkClient struct {
	*redis.Client
}

func NewRedisClient() *RedisBenchmarkClient {
	return &RedisBenchmarkClient{
		Client: redis.NewClient(&redis.Options{
			Addr: "localhost:7777",
			DB:   0,
		}),
	}
}

func (r *RedisBenchmarkClient) JSONSet(ctx context.Context, key string, path string, value any) {
	r.Client.JSONSet(ctx, key, path, value)
}

func (r *RedisBenchmarkClient) JSONGet(ctx context.Context, key string, path string) {
	r.Client.JSONGet(ctx, key, path)
}
