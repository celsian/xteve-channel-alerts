package utils

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

func SetupLogging() *os.File {
	// Create log directory if it doesn't exist
	err := os.MkdirAll("log", 0755)
	if err != nil {
		panic(fmt.Errorf("error creating log directory: %v", err))
	}

	// Open the file for writing, create if it doesn't exist
	f, err := os.OpenFile("log/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Errorf("error opening app.log file: %v", err))
	}

	// Create a MultiWriter that writes to both os.Stdout and the file
	multiWriter := io.MultiWriter(os.Stdout, f)

	// Create and set the default logger
	logger := slog.New(slog.NewTextHandler(multiWriter, nil))
	slog.SetDefault(logger)

	slog.Info("App starting, Log started")

	return f
}
