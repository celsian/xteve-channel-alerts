package file

import (
	"fmt"
	"io"
	"net/http"
	"os"
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
	// Save new list to file
	err := os.WriteFile("file/tmp/current.m3u", w, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}

func ReadFiles() (previous []byte, current []byte, err error) {
	c, err := os.ReadFile("file/tmp/current.m3u")
	if err != nil {
		return nil, nil, fmt.Errorf("error reading current file: %v", err)
	}

	p, err := os.ReadFile("file/tmp/previous.m3u")
	if err != nil {
		// If this is the first run, the previous file won't exist. Return an empty file to allow the app to create it.
		if os.IsNotExist(err) {
			return nil, c, nil
		}
		return nil, nil, fmt.Errorf("error reading previous file: %v", err)
	}

	return p, c, nil
}

func CleanUpFiles() error {
	// Move current.m3u to previous.m3u
	err := os.Rename("file/tmp/current.m3u", "file/tmp/previous.m3u")
	if err != nil {
		return fmt.Errorf("error moving file: %v", err)
	}

	return nil
}
