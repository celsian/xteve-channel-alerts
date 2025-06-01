package alerts

import (
	"bytes"
	"encoding/json"
	"fmt"
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

	fmt.Println("Alerting on missing channels: ", missing)

	description := ""

	for _, c := range missing {
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

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshaling json: %v", err)
	}

	err = sendAlert(jsonData)
	if err != nil {
		return fmt.Errorf("error sending alert: %v", err)
	}

	return nil
}

func TestAlert() error {
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

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshaling json: %v", err)
	}

	sendAlert(jsonData)
	if err != nil {
		return fmt.Errorf("error sending test alert: %v", err)
	}

	return nil
}

func sendAlert(jsonData []byte) error {
	DISCORD_WEBHOOK_URL := os.Getenv("DISCORD_WEBHOOK_URL")

	resp, err := http.Post(DISCORD_WEBHOOK_URL, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		return fmt.Errorf("error making POST request: %v", err)
	}
	defer resp.Body.Close()

	return nil
}
