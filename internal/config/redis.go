package config

import (
	"context"
	"os"

	"user-service/internal/logger"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func NewRedisClient() *redis.Client {
	addr := os.Getenv("USER_REDIS_HOST") + ":" + os.Getenv("USER_REDIS_PORT")

	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		logger.Log.Fatal("Failed to connect to Redis", zap.Error(err))
	}

	logger.Log.Info("Connected to Redis successfully")
	return client
}
