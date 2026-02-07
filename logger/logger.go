package logger

import (
	"log/slog"
	"os"
	"strings"
	"sync"
)

var (
	once     sync.Once
	instance *slog.Logger
)

func Init(l string) {
	once.Do(func() {
		level := parseLogLevel(l)

		handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})

		instance = slog.New(handler)
	})
}

func Get() *slog.Logger {
	if instance == nil {
		panic("logger not initialized — call logger.Init(level) first")
	}
	return instance
}

func parseLogLevel(level string) slog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
