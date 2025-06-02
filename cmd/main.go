package main

import (
	"log"
	"user-service/internal/auth"
	"user-service/internal/config"
	"user-service/internal/handler"
	"user-service/internal/logger"
	"user-service/internal/repository"
	"user-service/internal/service"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	logger.InitLogger()
	defer logger.Log.Sync()

	logger.Log.Info("Logger initialized")

	db := config.ConnectDatabase()
	repo := repository.NewUserRepository(db)
	redisClient := config.NewRedisClient()
	svc := service.NewUserService(repo, redisClient)
	h := handler.NewUserHandler(svc)

	jwtSecret := os.Getenv("JWT_SECRET")
	authService := auth.NewAuthService(repo, jwtSecret)
	authHandler := auth.NewAuthHandler(authService)

	app := fiber.New()

	app.Post("/user", h.Create)
	app.Get("/users", h.GetAll)
	app.Get("/user/:id", h.GetByID)
	app.Delete("/user/:id", h.Delete)

	app.Post("/login", authHandler.Login)

	app.Get("/ping-redis", handler.PingRedis)

	log.Fatal(app.Listen(":9000"))

}
