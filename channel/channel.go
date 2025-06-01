package channel

import (
	"bytes"
	"regexp"
	"slices"
)

type Channel struct {
	Number     string
	Title      string
	GroupTitle string
}

func ParseM3U(file []byte) ([]Channel, error) {
	var channels []Channel

	r := regexp.MustCompile(`#EXTINF:[^\n]*tvg-name="([^"]+)"[^\n]*tvg-id="([^"]+)"[^\n]*group-title="([^"]+)"[^\n]*`)

	lines := bytes.Split(file, []byte("\n"))

	for _, line := range lines {
		c := r.FindAllStringSubmatch(string(line), -1)
		if len(c) > 0 {
			channels = append(channels, Channel{Number: c[0][2], Title: c[0][1], GroupTitle: c[0][3]})
		}
	}

	return channels, nil
}

func CompareChannels(previous []Channel, current []Channel) []Channel {
	var missing []Channel

	// If any channel in the previous set is not present in the current set, add it to the missing slice
	for _, c := range previous {
		if slices.Index(current, c) == -1 {
			missing = append(missing, c)
		}
	}

	return missing
}
