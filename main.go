package main

import (
	"github.com/celsian/xteve-channel-alerts/cmd"
	"github.com/celsian/xteve-channel-alerts/pkg/utils"
)

func main() {
	err := cmd.Execute()
	utils.PanicOnErr(err)
}
