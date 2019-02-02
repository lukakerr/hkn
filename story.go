package hkn

import (
	"encoding/json"
	"errors"
	"fmt"
)

// GetNumber : Given a number a limit and a url, fetch from the url and
// return the number requested if it is >= 0 and <= limit
func GetNumber(number int, limit int, url string) ([]int, error) {
	if number > limit || number < 0 {
		msg := fmt.Sprintf("Invalid number. %d is not within the bounds of 0 and %d", number, limit)
		return nil, errors.New(msg)
	}

	resp, err := GetBody(url + JSONSuffix)

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

// GetTopStories : Get top stories given a number
func GetTopStories(number int, url string) ([]int, error) {
	resource := fmt.Sprintf("%s/%s", url, "topstories")
	return GetNumber(number, 500, resource)
}

// GetNewStories : Get new stories given a number
func GetNewStories(number int, url string) ([]int, error) {
	resource := fmt.Sprintf("%s/%s", url, "newstories")
	return GetNumber(number, 500, resource)
}

// GetBestStories : Get best stories given a number
func GetBestStories(number int, url string) ([]int, error) {
	resource := fmt.Sprintf("%s/%s", url, "beststories")
	return GetNumber(number, 500, resource)
}

// GetLatestAskStories : Get latest ask stories given a number
func GetLatestAskStories(number int, url string) ([]int, error) {
	resource := fmt.Sprintf("%s/%s", url, "askstories")
	return GetNumber(number, 200, resource)
}

// GetLatestShowStories : Get latest show stories given a number
func GetLatestShowStories(number int, url string) ([]int, error) {
	resource := fmt.Sprintf("%s/%s", url, "showstories")
	return GetNumber(number, 200, resource)
}

// GetLatestJobStories : Get latest job stories given a number
func GetLatestJobStories(number int, url string) ([]int, error) {
	resource := fmt.Sprintf("%s/%s", url, "jobstories")
	return GetNumber(number, 200, resource)
}
