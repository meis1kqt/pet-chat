package main

import (
	"chat-app/pkg/logger"
	"chat-app/services/auth/internal/app"
	"chat-app/services/auth/internal/config"
	"log/slog"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	config := config.MustLoadConfig()

	logger := logger.New(config.Env)

	logger.Info("Starting auth service", slog.String("env", config.Env))

	application := app.New(logger, config.GRPC.Port, config.Database.DSN(), config.JWT.Secret, config.JWT.TTL)


	if err := application.Run(); err != nil {
		logger.Error("we have some trouble", "error", err)
	}
	
}

