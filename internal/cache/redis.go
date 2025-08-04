package cache

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("USER_REDIS_HOST"), os.Getenv("USER_REDIS_PORT")),
		Password: "",
		DB:       0,
	})

	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}
}
