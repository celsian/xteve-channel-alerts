package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/celsian/xteve-channel-alerts/pkg/alerts"
	"github.com/celsian/xteve-channel-alerts/pkg/channel"
	"github.com/celsian/xteve-channel-alerts/pkg/file"
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
	slog.Info("xTeVe Channel Alerts starting channel check...")

	// Get current channel list
	w, err := file.GetCurrentChannelList()
	if err != nil {
		return err
	}

	// Write current channel list to file
	if err = file.WriteCurrentFile(w); err != nil {
		return err
	}

	// Load current and previous files
	previousM3U, currentM3U, err := file.ReadFiles()
	if err != nil {
		return err
	}

	// Parse both files into Channel structs
	pCh := channel.ParseM3U(previousM3U)
	cCh := channel.ParseM3U(currentM3U)

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
	if err = file.CleanUpFiles(); err != nil {
		slog.Error(fmt.Sprintf("error cleaning up files: %v", err))
		return err
	}

	return nil
}
