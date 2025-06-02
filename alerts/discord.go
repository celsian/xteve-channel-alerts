package alerts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/celsian/xteve-channel-alerts/channel"
)

type DiscordPayload struct {
	Username string   `json:"username"`
	Embeds   []Embeds `json:"embeds"`
}

type Embeds struct {
	Title       string `json:"title"`
	Color       string `json:"color"`
	Description string `json:"description"`
}

func DiscordAlert(missing []channel.Channel) error {

	slog.Info("Alerting Discord")

	description := ""

	for _, c := range missing {
		if len(description) >= 7000 {
			description += "...\n"
			description += "**Too many channels to list, check application log**"
			break
		}
		description += fmt.Sprintf("**%s** %s - %s\n", c.Number, c.Title, c.GroupTitle)
	}

	embed := Embeds{
		Title:       "Missing channels found",
		Color:       "16711680",
		Description: description,
	}

	data := DiscordPayload{
		Username: "xTeVe",
		Embeds:   []Embeds{embed},
	}

	err := sendAlert(data)
	if err != nil {
		return fmt.Errorf("error sending alert: %v", err)
	}

	return nil
}

func TestAlert() error {
	slog.Info("Sending test alert to Discord.")

	data := DiscordPayload{
		Username: "xTeVe",
		Embeds: []Embeds{
			{
				Title:       "Test Alert",
				Color:       "65280",
				Description: "Alert send was successful!",
			},
		},
	}

	err := sendAlert(data)
	if err != nil {
		return fmt.Errorf("error sending test alert: %v", err)
	}

	return nil
}

func MissingPreviousM3U() error {
	data := DiscordPayload{
		Username: "xTeVe",
		Embeds: []Embeds{
			{
				Title:       "Missing Previous M3U",
				Color:       "16776960",
				Description: "Previous M3U file was not found, this is expected on first run. If this is not your first run, check the application logs.",
			},
		},
	}

	err := sendAlert(data)
	if err != nil {
		return fmt.Errorf("error sending missing m3u alert: %v", err)
	}

	return nil
}

func sendAlert(data DiscordPayload) error {
	DISCORD_WEBHOOK_URL := os.Getenv("DISCORD_WEBHOOK_URL")

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshaling json: %v", err)
	}

	resp, err := http.Post(DISCORD_WEBHOOK_URL, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		return fmt.Errorf("error making POST request: %v", err)
	}
	defer resp.Body.Close()

	return nil
}
