package hkn

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	createStoryFormRegex = `<input\stype=['"]hidden['"]\s+name=['"]fnid['"]\s+(?:[^>]*?\s+)?value=['"]([^'"]*)['"]`
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

// Create a story given a title, content, cookie and content key (either "text" or "url")
func createStory(title string, content string, cookie *http.Cookie, webURL string, contentKey string) (bool, error) {
	if len(title) == 0 {
		return false, ErrEmptyTitle
	}

	submitFormURL := fmt.Sprintf("%s/%s", webURL, "submit")
	fnID, err := matchRegexFromBody(submitFormURL, createStoryFormRegex, cookie)

	if err != nil {
		return false, err
	}

	submitURL := fmt.Sprintf("%s/%s", webURL, "r")

	body := url.Values{}
	body.Set("fnid", fnID)
	body.Set("fnop", "submit-page")
	body.Set("title", title)

	if contentKey == "text" {
		body.Set("url", "")
		body.Set("text", content)
	} else if contentKey == "url" {
		body.Set("url", content)
		body.Set("text", "")
	}

	resp, err := postWithCookie(submitURL, body, cookie)

	if err == nil && resp != nil {
		return true, nil
	}

	return false, err
}
