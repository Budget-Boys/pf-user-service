package config

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"user-service/internal/logger"
	"user-service/internal/model"
)

func ConnectDatabase() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Log.Fatal("Failed to connect to database", zap.Error(err))
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		logger.Log.Fatal("Failed to run migrations", zap.Error(err))
	}

	logger.Log.Info("Connected to MySQL successfully")
	return db
}
