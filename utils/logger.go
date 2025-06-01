package utils

import (
	"fmt"
	"time"
)

// Log outputs to the command line and to a log file
func Log(line string) {
	// Print to console
	ll := fmt.Sprintf("%s: %s", time.Now().Format("2006-01-02 15:04:05 MST"), line)
	fmt.Println(ll)

	// Log to file
}
