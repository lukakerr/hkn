package hkn

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"net/url"
	"sync"
)

// Item represents a Hacker News item
type Item struct {
	ID          int    `json:"id"`
	Deleted     bool   `json:"deleted"`
	Type        string `json:"type"`
	By          string `json:"by"`
	Time        int32  `json:"time"`
	Text        string `json:"text"`
	Dead        bool   `json:"dead"`
	Parent      int    `json:"parent"`
	Poll        int    `json:"poll"`
	Kids        []int  `json:"kids"`
	URL         string `json:"url"`
	Score       int    `json:"score"`
	Title       string `json:"title"`
	Parts       []int  `json:"parts"`
	Descendants int    `json:"descendants"`
}

// Items represents an array of items
type Items []Item

const (
	voteURLRegex    = `<a\s+id=['"]%s_%d['"]\s+(?:[^>]*?\s+)?href=['"]([^'"]*)['"]`
	commentURLRegex = `<input\s+type=['"]hidden['"]\s+name=['"]hmac['"]\s+(?:[^>]*?\s+)?value=['"]([^'"]*)['"]`
)

// Get an item given an id
func getItem(id int, apiURL string) (Item, error) {
	reqURL := fmt.Sprintf("%s/%s/%d", apiURL, "item", id) + jsonSuffix

	resp, err := getBody(reqURL)

	var item Item

	if err != nil {
		return item, err
	}

	err = json.Unmarshal(resp, &item)
	return item, err
}

// Get items given a slice of ids
// This function is concurrent and thus does not guarantee order
func getItems(ids []int, apiURL string) (Items, error) {
	var (
		items Items
		wg    sync.WaitGroup
	)

	// Add length of ids to the WaitGroup
	wg.Add(len(ids))

	for _, id := range ids {
		// Spawn a thread for each item
		go func(id int, url string) {
			defer wg.Done()

			item, err := getItem(id, url)

			if err != nil {
				return
			}

			items = append(items, item)
		}(id, apiURL)
	}

	// Wait until all threads are done
	wg.Wait()

	return items, nil
}

// Get the most recent item id
func getMaxItemID(apiURL string) (int, error) {
	reqURL := fmt.Sprintf("%s/%s", apiURL, "maxitem") + jsonSuffix

	resp, err := getBody(reqURL)

	var id int

	if err != nil {
		return id, err
	}

	err = json.Unmarshal(resp, &id)
	return id, err
}

func vote(id int, cookie *http.Cookie, webURL string, voteType string) (bool, error) {
	reqURL := fmt.Sprintf("%s/%s?id=%d", webURL, "item", id)
	upvoteRegex := fmt.Sprintf(voteURLRegex, voteType, id)

	voteAuth, err := matchRegexFromBody(reqURL, upvoteRegex, cookie)

	if err != nil {
		return false, err
	}

	voteURL := fmt.Sprintf("%s/%s", webURL, voteAuth)
	unescaped := html.UnescapeString(voteURL)

	resp, err := getBodyWithCookie(unescaped, cookie)

	if err == nil && resp != nil {
		return true, nil
	}

	return false, err
}

// Upvote an item given an id and a cookie
func upvote(id int, cookie *http.Cookie, webURL string) (bool, error) {
	return vote(id, cookie, webURL, "up")
}

// Unvote a comment given an id and a cookie
func unvote(id int, cookie *http.Cookie, webURL string) (bool, error) {
	return vote(id, cookie, webURL, "un")
}

// Create a comment on an item given an id and content
func comment(id int, content string, cookie *http.Cookie, webURL string) (bool, error) {
	if len(content) == 0 {
		return false, ErrEmptyContent
	}

	reqURL := fmt.Sprintf("%s/%s?id=%d", webURL, "item", id)

	commentAuth, err := matchRegexFromBody(reqURL, commentURLRegex, cookie)

	if err != nil {
		return false, err
	}

	commentURL := fmt.Sprintf("%s/%s", webURL, "comment")

	body := url.Values{}
	body.Set("parent", fmt.Sprintf("%d", id))
	body.Set("goto", fmt.Sprintf("item?id=%d", id))
	body.Set("hmac", commentAuth)
	body.Set("text", content)

	resp, err := postWithCookie(commentURL, body, cookie)

	if err == nil && resp != nil {
		return true, nil
	}

	return false, err
}
