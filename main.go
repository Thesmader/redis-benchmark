package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	ddb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:7777",
		DB:   0,
	})

	s := rdb.Ping(ctx)
	fmt.Printf("Vanilla Redis: %v\n", s)
	s1 := ddb.Ping(ctx)
	fmt.Printf("Dragonfly: %v\n", s1)
}
