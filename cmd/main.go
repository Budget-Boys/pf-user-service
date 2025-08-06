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
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	logger.InitLogger()
	defer logger.Log.Sync()

	logger.Log.Info("Logger initialized")

	eurekaClient := config.NewEurekaClient()

	if err := eurekaClient.Register(); err != nil {
		logger.Log.Info("Error registering service with Eureka: " + err.Error())
	}

	eurekaClient.StartHeartbeat()

	db := config.ConnectDatabase()
	repo := repository.NewUserRepository(db)
	redisClient := config.NewRedisClient()
	svc := service.NewUserService(repo, redisClient)
	h := handler.NewUserHandler(svc)

	jwtSecret := os.Getenv("USER_JWT_SECRET")
	authService := auth.NewAuthService(repo, jwtSecret)
	authHandler := auth.NewAuthHandler(authService)

	app := fiber.New()

	 app.Use(cors.New(cors.Config{
        AllowOrigins:     "http://localhost:5173", // ou o domínio do seu frontend
        AllowCredentials: true,
        AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
        AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
    }))
	app.Post("/users", h.Create)
	app.Get("/users", h.GetAll)
	app.Get("/users/:id", h.GetByID)
	app.Delete("/users/:id", h.Delete)
	app.Put("/users/:id", h.Update)

	app.Post("/login", authHandler.Login)

	app.Get("/ping-redis", handler.PingRedis)

	log.Fatal(app.Listen(":8080"))

}
