package main

import (
	"context"
	"fmt"

	"log/slog"

	"github.com/redis/go-redis/v9"
)

const channel = "message_queue"

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	defer rdb.Close()

	ctx := context.Background()

	subscriber := rdb.Subscribe(ctx, channel)
	channel := subscriber.Channel()

	slog.Info("Waiting for messages...")

	for msg := range channel {
		slog.Info(fmt.Sprintf("Received: %s", msg.Payload))
	}
}
