package hkn

import (
	"encoding/json"
	"fmt"
)

// Updates represents profile and item updates
type Updates struct {
	Items    []int    `json:"items"`
	Profiles []string `json:"profiles"`
}

// Get the latest item and profile updates
func getUpdates(apiURL string) (Updates, error) {
	reqURL := fmt.Sprintf("%s/%s", apiURL, "updates") + jsonSuffix

	resp, err := getBody(reqURL)

	var updates Updates

	if err != nil {
		return updates, err
	}

	err = json.Unmarshal(resp, &updates)
	return updates, err
}
