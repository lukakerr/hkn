package hkn

import (
	"encoding/json"
	"fmt"
)

// Updates : Struct representing profile and item updates
type Updates struct {
	Items    []int    `json:"items"`
	Profiles []string `json:"profiles"`
}

// Get the latest item and profile updates
func getUpdates(url string) (Updates, error) {
	reqURL := fmt.Sprintf("%s/%s", url, "updates") + jsonSuffix

	resp, err := getBody(reqURL)

	var updates Updates

	if err != nil {
		return updates, err
	}

	err = json.Unmarshal(resp, &updates)
	return updates, err
}
