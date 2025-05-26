package main

import (
	"log"
	"user-service/internal/config"
	"user-service/internal/handler"
	"user-service/internal/repository"
	"user-service/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db := config.ConnectDatabase()
	repo := repository.NewUserRepository(db)
	redisClient := config.NewRedisClient()
	svc := service.NewUserService(repo, redisClient)
	h := handler.NewUserHandler(svc)

	app := fiber.New()

	app.Post("/user", h.Create)
	app.Get("/users", h.GetAll)
	app.Get("/user/:id", h.GetByID)
	app.Delete("/user/:id", h.Delete)

	app.Get("/ping-redis", handler.PingRedis)

	log.Fatal(app.Listen(":9000"))

}
