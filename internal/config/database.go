package config

import (
	"fmt"
	"os"

	"user-service/internal/logger"
	"user-service/internal/model"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("USER_DB_USER"),
		os.Getenv("USER_DB_PASSWORD"),
		os.Getenv("USER_DB_HOST"),
		os.Getenv("USER_DB_PORT"),
		os.Getenv("USER_DB_NAME"),
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
