package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type BenchmarkClient interface {
	*redis.Client
	JSONSet(ctx context.Context, key string, path string, value interface{})
	JSONGet(ctx context.Context, key string, path string)
}
