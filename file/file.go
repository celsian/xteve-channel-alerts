package file

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"github.com/celsian/xteve-channel-alerts/alerts"
)

func GetCurrentChannelList() ([]byte, error) {
	XTEVE_URL := os.Getenv("XTEVE_URL")

	// Get new (current) channel list
	resp, err := http.Get(XTEVE_URL)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %v", err)
	}
	defer resp.Body.Close()

	w, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	return w, nil
}

func WriteCurrentFile(w []byte) error {
	// Define the file path
	filePath := "data/m3us/current.m3u"
	
	// Create the directory if it doesn't exist
	dir := filepath.Dir(filePath)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("error creating directory: %v", err)
	}
	
	// Save new list to file
	err = os.WriteFile(filePath, w, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}

func ReadFiles() (previous []byte, current []byte, err error) {
	c, err := os.ReadFile("data/m3us/current.m3u")
	if err != nil {
		return nil, nil, fmt.Errorf("error reading current file: %v", err)
	}

	p, err := os.ReadFile("data/m3us/previous.m3u")
	if err != nil {
		// If this is the first run, the previous file won't exist. Return an empty file to allow the app to create it.
		if os.IsNotExist(err) {
			slog.Warn("Previous m3u file not found, this is expected on first run.")
			err = alerts.MissingPreviousM3U()
			if err != nil {
				return nil, c, err
			}
			return nil, c, nil
		}
		return nil, nil, fmt.Errorf("error reading previous file: %v", err)
	}

	return p, c, nil
}

func CleanUpFiles() error {
	// Move current.m3u to previous.m3u
	err := os.Rename("data/m3us/current.m3u", "data/m3us/previous.m3u")
	if err != nil {
		return fmt.Errorf("error moving file: %v", err)
	}

	return nil
}
