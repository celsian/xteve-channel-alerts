package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/celsian/xteve-channel-alerts/utils"
	"github.com/joho/godotenv"
)

func setup() *os.File {
	// Try to load .env file, but don't panic if it doesn't exist
	err := godotenv.Load()
	if err != nil {
		// Just log the error instead of panicking
		fmt.Println("No .env file found, using environment variables directly")
	} else {
		fmt.Println("Loaded configuration from .env file")
	}

	// Setup logging as before
	logFile := utils.SetupLogging()
	
	// Log environment variable usage after logging is set up
	slog.Info("Using environment variables for configuration")
	
	return logFile
}
