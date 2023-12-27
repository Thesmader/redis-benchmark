package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type DragonflyBenchmarkClient struct {
	*redis.Client
}

func NewDragonflyClient() *DragonflyBenchmarkClient {
	return &DragonflyBenchmarkClient{
		Client: redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			DB:   0,
		}),
	}
}

func (d *DragonflyBenchmarkClient) JSONSet(ctx context.Context, key string, path string, value interface{}) {
	d.Client.JSONSet(ctx, key, path, value)
}

func (d *DragonflyBenchmarkClient) JSONGet(ctx context.Context, key string, path string) {
	d.Client.JSONGet(ctx, key, path)
}
