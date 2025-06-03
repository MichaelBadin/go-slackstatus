package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const slackAPIURL = "https://slack.com/api/users.profile.set"

type slackProfile struct {
	StatusText       string `json:"status_text"`
	StatusEmoji      string `json:"status_emoji"`
	StatusExpiration int    `json:"status_expiration"`
}

type slackPayload struct {
	Profile slackProfile `json:"profile"`
}

func SetStatus(token, text, emoji string, expiration int) error {
	payload := slackPayload{
		Profile: slackProfile{
			StatusText:       text,
			StatusEmoji:      emoji,
			StatusExpiration: expiration,
		},
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", slackAPIURL, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	if ok, _ := result["ok"].(bool); !ok {
		return fmt.Errorf("Slack error: %v", result["error"])
	}

	return nil
}

func ClearStatus(token string) error {
	return SetStatus(token, "", "", 0)
}
