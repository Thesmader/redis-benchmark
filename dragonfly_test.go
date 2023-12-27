package main

import (
	"context"
	"strconv"
	"testing"
)

func BenchmarkDragonfly(b *testing.B) {
	ctx := context.Background()
	r := NewDragonflyClient()
	defer r.Client.FlushDB(ctx)
	for i := 0; i < b.N; i++ {
		r.JSONSet(
			ctx,
			strconv.Itoa(i),
			"$",
			map[string]any{"foo": "bar"},
		)
	}
}
