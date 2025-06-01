package main

import (
	"fmt"

	"github.com/celsian/xteve-channel-alerts/alerts"
	"github.com/celsian/xteve-channel-alerts/channel"
	"github.com/celsian/xteve-channel-alerts/file"
	"github.com/celsian/xteve-channel-alerts/utils"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	// Get current channel list
	w, err := file.GetCurrentChannelList()
	utils.HandleErr(err)

	// Write current channel list to file
	file.WriteCurrentFile(w)
	utils.HandleErr(err)

	// Load current and previous files
	previous, current, err := file.ReadFiles()
	utils.HandleErr(err)

	// Parse both files into Channel structs
	pCh, err := channel.ParseM3U(previous)
	utils.HandleErr(err)

	cCh, err := channel.ParseM3U(current)
	utils.HandleErr(err)

	// Find missing channels
	missing := channel.CompareChannels(pCh, cCh)

	if missing == nil {
		// No changes
		fmt.Println("No changes, skipping alert.")
	} else {
		// Alert Discord
		alerts.DiscordAlert(missing)
	}

	// Cleanup files for next run
	file.CleanUpFiles()
}
