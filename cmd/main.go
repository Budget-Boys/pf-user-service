package main

import (
	"log"
	"user-service/internal/config"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
    db := config.ConnectDatabase()

    sqlDB, err := db.DB()
    if err != nil {
        log.Fatal("Error getting sql.DB:", err)
    }

    if err := sqlDB.Ping(); err != nil {
        log.Fatal("Failed to ping DB:", err)
    }

    log.Println("Ping DB successful — migrations OK")
}
