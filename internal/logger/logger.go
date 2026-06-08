package logger

import (
	"log/slog"
	"os"
)

var log *slog.Logger

func Init(level string) {

	var logLevel slog.Level
	switch level {
	case "info":
		logLevel = slog.LevelInfo
	case "debug":
		logLevel = slog.LevelDebug
	case "error":
		logLevel = slog.LevelError
	case "warning":
		logLevel = slog.LevelWarn
	default:
		slog.Info("Set default log level")
		logLevel = slog.LevelInfo
	}

	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel})
	log = slog.New(logHandler).With("service", "email-service")
}

func Info(message string, args ...any) {
	log.Info(message, args...)
}

func Debug(message string, args ...any) {
	log.Debug(message, args...)
}

func Error(message string, args ...any) {
	log.Error(message, args...)
}

func Warn(message string, args ...any) {
	log.Warn(message, args...)
}
