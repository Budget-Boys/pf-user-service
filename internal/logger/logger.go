// Package logger initializes a global structured logger using the zap library.
//
// Available levels:
// - Debug: for detailed development information
// - Info: normal informational messages (e.g., "Cache hit", "User created")
// - Warn: warnings that don't break execution but deserve attention
// - Error: failures that should be handled or monitored
// - Fatal: critical errors that terminate the application
//
// Usage example:
//   logger.Log.Info("User authenticated", zap.String("userID", "123"))
//   logger.Log.Error("Failed to save to database", zap.Error(err))
//
// In development, we use zap.NewDevelopment() for human-readable logs.
// In production, you can switch to zap.NewProduction() (JSON, optimized).

package logger

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

func InitLogger() {
	var err error

	Log, err = zap.NewDevelopment() // Zap in readable mode for development
	// Log, err = zap.NewProduction() // For production: optimized JSON

	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}
}
