package logger

import (
	"log/slog"
	"os"
)

var ( 
	local = "local"
	production = "production"
	test = "test"
)

func New(env string) *slog.Logger {
	

	var logger *slog.Logger
	
	switch env {
	case local, test:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case production:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	default:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	}

	return logger
}