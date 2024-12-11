package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"sync"
	"syscall"
	"time"
)

const (
	logDirPath = "logs"
)

var (
	globalLogger *slog.Logger
	loggerFile   *os.File
	currentDate  string
	mu           sync.Mutex
)

// Get creates a new slog.Logger instance if it's not exists. Otherwise it returns g
func Get(pkg string, fn string) *slog.Logger {
	mu.Lock()
	defer mu.Unlock()

	today := time.Now().Format(time.DateOnly)

	if globalLogger == nil || today != currentDate {
		// Close the previous file if it exists
		if loggerFile != nil {
			_ = loggerFile.Close()
		}

		// Update the current date
		currentDate = today

		// Ensure the log directory exists
		syscall.Umask(0)
		defer syscall.Umask(022)
		if _, err := os.Stat(logDirPath); os.IsNotExist(err) {
			err = os.MkdirAll(logDirPath, os.ModePerm)
			if err != nil {
				panic(fmt.Sprintf("failed to create folder: %v", err))
			}
		}

		// Create a new log file for today
		fileName := today + ".log"
		filePath := filepath.Join(logDirPath, fileName)
		loggerFile, err := os.OpenFile(filepath.Clean(filePath), os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			panic(fmt.Sprintf("failed to create log file: %v", err))
		}

		writer := io.MultiWriter(os.Stdout, loggerFile)

		// Create a new logger instance
		globalLogger = slog.New(slog.NewTextHandler(writer, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	}

	return globalLogger.WithGroup(pkg).With("fn", fn)
}

// CloseFile closes the logger file
func CloseFile() error {
	mu.Lock()
	defer mu.Unlock()
	if loggerFile != nil {
		return loggerFile.Close()
	}
	return nil
}
