package main

import (
	"context"
	"strconv"
	"testing"
)

func BenchmarkRedis(b *testing.B) {
	ctx := context.Background()
	r := NewRedisClient()
	r.Client.FlushDB(ctx)
	for i := 0; i < b.N; i++ {
		// key := "benchtest_" + strconv.Itoa(i)
		r.JSONSet(
			ctx,
			strconv.Itoa(i),
			"$",
			map[string]any{"foo": "bar"},
		)
	}
}
