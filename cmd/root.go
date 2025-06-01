package cmd

import (
	"context"
	"os"

	"github.com/urfave/cli"
)

var app = &cli.App{
	Name:  "xteve-channel-alerts",
	Usage: "A service to alert a discord webhook when xteve channels disappear",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "config",
			Required: true,
		},
	},
}

func Execute(ctx context.Context) error {
	err := app.Run(os.Args)
	if err != nil {
		return err
	}
	return err
}
