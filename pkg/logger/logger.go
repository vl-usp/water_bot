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
	once         sync.Once
)

// Get creates a new slog.Logger instance if it's not exists. Otherwise it returns g
func Get(pkg string, fn string) *slog.Logger {
	once.Do(func() {
		syscall.Umask(0)
		defer syscall.Umask(022)

		if _, err := os.Stat(logDirPath); os.IsNotExist(err) {
			err = os.MkdirAll(logDirPath, os.ModePerm)
			if err != nil {
				panic(fmt.Sprintf("failed to create folder: %v", err))
			}
		}

		fileName := time.Now().Format(time.DateOnly) + ".log"
		filePath := logDirPath + "/" + fileName
		loggerFile, err := os.OpenFile(filepath.Clean(filePath), os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			panic(fmt.Sprintf("failed to create folder: %v", err))
		}
		writer := io.MultiWriter(os.Stdout, loggerFile)

		globalLogger = slog.New(slog.NewTextHandler(writer, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	})

	return globalLogger.With("pkg", pkg, "fn", fn)
}

// CloseFile closes the logger file
func CloseFile() error {
	return loggerFile.Close()
}
