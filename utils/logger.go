package utils

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
)

func SetupLogging() *os.File {
	// Create logs directory if it doesn't exist
	logsDir := "data/logs"
	err := os.MkdirAll(logsDir, 0755)
	if err != nil {
		panic(fmt.Errorf("error creating logs directory: %v", err))
	}

	// Open the file for writing, create if it doesn't exist
	logPath := filepath.Join(logsDir, "app.log")
	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Errorf("error opening app.log file: %v", err))
	}

	// Create a MultiWriter that writes to both os.Stdout and the file
	multiWriter := io.MultiWriter(os.Stdout, f)

	// Create and set the default logger
	logger := slog.New(slog.NewTextHandler(multiWriter, nil))
	slog.SetDefault(logger)

	return f
}
