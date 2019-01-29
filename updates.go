package hkn

import (
	"encoding/json"
	"fmt"
)

type Updates struct {
	Items    []int    `json:"items"`
	Profiles []string `json:"profiles"`
}

// Get the latest item and profile updates
func GetUpdates(url string) (Updates, error) {
	reqUrl := fmt.Sprintf("%s/%s", url, "updates")

	resp, err := GetBody(reqUrl)

	var updates Updates

	if err != nil {
		return updates, err
	}

	err = json.Unmarshal(resp, &updates)
	return updates, err
}
