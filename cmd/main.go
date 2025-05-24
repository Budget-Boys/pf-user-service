package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Carrega variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		log.Println("Erro ao carregar .env, usando variáveis padrão.")
	}

	// Cria conexão mínima com MySQL (não usada de fato, só pra manter dependência)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		"root", "secret", "localhost", "user_service")
	_, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Não conectou ao MySQL (ok para teste de dependência)")
	}

	// Inicia Fiber
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("pf-user-service is running!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	log.Fatal(app.Listen(":" + port))
}
