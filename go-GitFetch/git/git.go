package git

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GithubEvents struct {
	Type      string      `json:"type"`
	Repo      Repo        `json:"repo"`
	Payload   interface{} `json:"payload"`
	CreatedAt string      `json:"created_at"`
}

type Repo struct {
	Name string `json:"name"`
}

func GetEvents(username string) ([]GithubEvents, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	print(url)

	resp, err := http.Get(url)

	if err != nil {
		return []GithubEvents{}, fmt.Errorf("GitFetch request failed %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("github API error: status code %d", resp.StatusCode)
	}

	var events []GithubEvents
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, fmt.Errorf("json decoding failed %w", err)
	}
	return events, nil
}
