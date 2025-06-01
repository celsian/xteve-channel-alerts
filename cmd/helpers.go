package cmd

import (
	"fmt"
	"os"

	"github.com/celsian/xteve-channel-alerts/utils"
	"github.com/joho/godotenv"
)

func setup() *os.File {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Errorf("error loading .env file: %v", err))
	}

	return utils.SetupLogging()
}
