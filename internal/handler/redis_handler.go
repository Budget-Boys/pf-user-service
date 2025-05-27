package handler

import (
	"github.com/gofiber/fiber/v2"
	"user-service/internal/cache"
)

func PingRedis(c *fiber.Ctx) error {
	err := cache.RedisClient.Set(cache.Ctx, "hello", "world", 0).Err()
	if err != nil {
		return c.Status(500).SendString("Failed to set value in Redis")
	}

	val, err := cache.RedisClient.Get(cache.Ctx, "hello").Result()
	if err != nil {
		return c.Status(500).SendString("Failed to get value from Redis")
	}

	return c.SendString("Redis responded: " + val)
}
