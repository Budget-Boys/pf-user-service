package config

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"user-service/internal/logger"
)

func NewRedisClient() *redis.Client {
	addr := os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")

	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		logger.Log.Fatal("Failed to connect to Redis", zap.Error(err))
	}

	logger.Log.Info("Connected to Redis successfully")
	return client
}
