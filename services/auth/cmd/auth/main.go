package main

import (
	"chat-app/pkg/logger"
	"chat-app/services/auth/internal/config"
	"log/slog"
)

func main() {

	config := config.MustLoadConfig()

	logger := logger.New(config.Env)

	logger.Info("Starting auth service", slog.String("env", config.Env))



	
}

