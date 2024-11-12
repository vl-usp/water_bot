package logger

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	logDir = "./logs"
)

var (
	log  *slog.Logger
	file *os.File
	once sync.Once
)

func Get() *slog.Logger {
	once.Do(func() {
		filename := time.Now().Format("2006-01-02") + ".log"
		err := os.MkdirAll(logDir, os.ModePerm)
		if err != nil {
			panic("failed to create log directory: " + err.Error())
		}

		fp := filepath.Clean(logDir + "/" + filename)
		file, err = os.OpenFile(fp, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			panic("failed to open log file: " + err.Error())
		}

		writer := io.MultiWriter(os.Stdout, file)
		log = slog.New(slog.NewTextHandler(writer, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	})

	return log
}
