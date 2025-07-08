package channel

import (
	"bytes"
	"fmt"
	"log/slog"
	"regexp"
	"slices"
)

type Channel struct {
	Number     string
	Title      string
	GroupTitle string
}

func ParseM3U(file []byte) []Channel {
	var channels []Channel

	r := regexp.MustCompile(`#EXTINF:[^\n]*tvg-name="([^"]+)"[^\n]*tvg-id="([^"]+)"[^\n]*group-title="([^"]+)"[^\n]*`)

	lines := bytes.SplitSeq(file, []byte("\n"))

	for line := range lines {
		c := r.FindAllStringSubmatch(string(line), -1)
		if len(c) > 0 {
			channels = append(channels, Channel{Number: c[0][2], Title: c[0][1], GroupTitle: c[0][3]})
		}
	}

	return channels
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

func (c Channel) Log() {
	slog.Info(fmt.Sprintf("%s - %s - %s", c.Number, c.Title, c.GroupTitle))
}

func (c Channel) LogWarning() {
	slog.Warn(fmt.Sprintf("%s - %s - %s", c.Number, c.Title, c.GroupTitle))
}
