package config

import (
	"log/slog"
	"os"
)

func initLogger(stage string) *slog.Logger {
	var logger *slog.Logger

	if stage == "production" {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	} else {
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}

	return logger
}
