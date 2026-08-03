package database

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	Redis *redis.Client
	Ctx   = context.Background()
)

func ConnectRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if err := Redis.Ping(Ctx).Err(); err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}
}
