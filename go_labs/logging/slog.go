package logging

import (
	"log"
	"log/slog"
	"os"
)

func tourOfSlog() {
	slog.Info("hello world")
	slog.Info("hello, world", "user", "test")
	slog.Info("hello, world", "user", os.Getenv("USER"))
	slog.Debug("hello world")
	slog.Warn("hello world")
	slog.Error("hello world")
	log.Println("hello world")

	logger := slog.Default()
	logger.Info("hello world")
}
