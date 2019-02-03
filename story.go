package hkn

import (
	"encoding/json"
	"fmt"
)

// Given a number a limit and a url, fetch from the url and
// return the number requested if it is >= 0 and <= limit
func getNumber(number int, limit int, url string) ([]int, error) {
	if number > limit || number < 0 {
		return nil, ErrInvalidNumber
	}

	resp, err := getBody(url + jsonSuffix)

	var top []int

	if err != nil {
		return top, err
	}

	err = json.Unmarshal(resp, &top)

	if err != nil {
		return top, err
	}

	if len(top) < number {
		return top, nil
	}

	return top[:number], nil
}

// Get top stories given a number
func getTopStories(number int, apiURL string) ([]int, error) {
	resource := fmt.Sprintf("%s/%s", apiURL, "topstories")
	return getNumber(number, 500, resource)
}

// Get new stories given a number
func getNewStories(number int, apiURL string) ([]int, error) {
	resource := fmt.Sprintf("%s/%s", apiURL, "newstories")
	return getNumber(number, 500, resource)
}

// Get best stories given a number
func getBestStories(number int, apiURL string) ([]int, error) {
	resource := fmt.Sprintf("%s/%s", apiURL, "beststories")
	return getNumber(number, 500, resource)
}

// Get latest ask stories given a number
func getLatestAskStories(number int, apiURL string) ([]int, error) {
	resource := fmt.Sprintf("%s/%s", apiURL, "askstories")
	return getNumber(number, 200, resource)
}

// Get latest show stories given a number
func getLatestShowStories(number int, apiURL string) ([]int, error) {
	resource := fmt.Sprintf("%s/%s", apiURL, "showstories")
	return getNumber(number, 200, resource)
}

// Get latest job stories given a number
func getLatestJobStories(number int, apiURL string) ([]int, error) {
	resource := fmt.Sprintf("%s/%s", apiURL, "jobstories")
	return getNumber(number, 200, resource)
}
