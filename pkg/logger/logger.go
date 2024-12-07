package logger

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

// SetupLogger creates a new logger.
func SetupLogger(env string, logDirPath string) (*slog.Logger, *os.File) {
	var log *slog.Logger
	syscall.Umask(0)
	defer syscall.Umask(022)

	if _, err := os.Stat(logDirPath); os.IsNotExist(err) {
		err := os.MkdirAll(logDirPath, os.ModePerm)
		if err != nil {
			panic("failed to create log directory: " + err.Error())
		}
	}

	fileName := time.Now().Format(time.DateOnly) + ".log"
	filePath := logDirPath + "/" + fileName
	file, err := os.OpenFile(filepath.Clean(filePath), os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic("failed to open log file: " + err.Error())
	}
	writer := io.MultiWriter(os.Stdout, file)

	switch env {
	case envLocal, envDev:
		log = slog.New(slog.NewJSONHandler(writer, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(writer, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	}

	return log, file
}
