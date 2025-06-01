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
	setup()

	// Get current channel list
	w, err := file.GetCurrentChannelList()
	utils.PanicOnErr(err)

	// Write current channel list to file
	file.WriteCurrentFile(w)
	utils.PanicOnErr(err)

	// Load current and previous files
	previous, current, err := file.ReadFiles()
	utils.PanicOnErr(err)

	// Parse both files into Channel structs
	pCh, err := channel.ParseM3U(previous)
	utils.PanicOnErr(err)

	cCh, err := channel.ParseM3U(current)
	utils.PanicOnErr(err)

	// Find missing channels
	missing := channel.CompareChannels(pCh, cCh)

	if missing == nil {
		// No changes
		utils.Log("No changes, skipping alert.")
	} else {
		utils.Log(fmt.Sprintf("Found %d Missing Channel(s):", len(missing)))
		for _, c := range missing {
			c.Print()
		}
		// Alert Discord
		err = alerts.DiscordAlert(missing)
		if err != nil {
			utils.Log(fmt.Sprintf("error sending Discord alert: %v", err))
		}
	}

	// Cleanup files for next run
	file.CleanUpFiles()
}

func setup() {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Errorf("error loading .env file: %v", err))
	}
}
