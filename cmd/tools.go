package cmd

import (
	"github.com/celsian/xteve-channel-alerts/pkg/alerts"
	"github.com/urfave/cli"
)

var testEnvironmnet = cli.Command{
	Name:        "testEnvironment",
	Description: "Test the .env file is setup correctly. Sends a test alert to your Discord webhook.",
	Aliases:     []string{"test"},
	Action: func(c *cli.Context) error {
		f := setup()
		defer f.Close()

		return alerts.TestAlert()
	},
}
