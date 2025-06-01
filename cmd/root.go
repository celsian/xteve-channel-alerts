package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/celsian/xteve-channel-alerts/alerts"
	"github.com/celsian/xteve-channel-alerts/channel"
	"github.com/celsian/xteve-channel-alerts/file"
	"github.com/urfave/cli"
)

var app = &cli.App{
	Name:  "xteve-channel-alerts",
	Usage: "A service to check that the current m3u is not missing channels present in the previous m3u.",
	Commands: []cli.Command{
		{
			Name:    "App",
			Aliases: []string{"start", "s"},
			Usage:   "Start the app",
			Action: func(c *cli.Context) (err error) {
				logFile := setup()
				defer logFile.Close()

				return root()
			},
		},
		testEnvironmnet,
	},
}

func Execute() error {
	err := app.Run(os.Args)
	if err != nil {
		return err
	}
	return err
}

func root() error {
	// Get current channel list
	w, err := file.GetCurrentChannelList()
	if err != nil {
		return err
	}

	// Write current channel list to file
	file.WriteCurrentFile(w)
	if err != nil {
		return err
	}

	// Load current and previous files
	previous, current, err := file.ReadFiles()
	if err != nil {
		return err
	}

	// Parse both files into Channel structs
	pCh, err := channel.ParseM3U(previous)
	if err != nil {
		return err
	}

	cCh, err := channel.ParseM3U(current)
	if err != nil {
		return err
	}

	// Find missing channels
	missing := channel.CompareChannels(pCh, cCh)

	if missing == nil {
		// No changes
		slog.Info("No changes, skipping alert.")
	} else {
		slog.Warn(fmt.Sprintf("Found %d Missing Channel(s):", len(missing)))
		for _, c := range missing {
			c.LogWarning()
		}
		// Alert Discord
		err = alerts.DiscordAlert(missing)
		if err != nil {
			slog.Error(fmt.Sprintf("error sending Discord alert: %v", err))
		}
	}

	// Cleanup files for next run
	file.CleanUpFiles()

	return nil
}
