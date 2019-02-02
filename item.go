package hkn

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	NetURL "net/url"
	"regexp"
	"sync"
)

// Item : A Hacker News item
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

// Items : An array of items
type Items []Item

const (
	voteURLRegex    = `<a\s+id=['"]%s_%d['"]\s+(?:[^>]*?\s+)?href=['"]([^'"]*)['"]`
	commentURLRegex = `<input\s+type=['"]hidden['"]\s+name=['"]hmac['"]\s+(?:[^>]*?\s+)?value=['"]([^'"]*)['"]`
)

// GetItem : Get an item given an id
func GetItem(id int, url string) (Item, error) {
	reqURL := fmt.Sprintf("%s/%s/%d", url, "item", id) + JSONSuffix

	resp, err := GetBody(reqURL)

	var item Item

	if err != nil {
		return item, err
	}

	err = json.Unmarshal(resp, &item)
	return item, err
}

// GetItems : Get items given a slice of ids
// This function is parallelised and thus does not guarantee order
func GetItems(ids []int, url string) (Items, error) {
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

			item, err := GetItem(id, url)

			if err != nil {
				return
			}

			items = append(items, item)
		}(id, url)
	}

	// Wait until all threads are done
	wg.Wait()

	return items, nil
}

// GetMaxItemID : Get the most recent item id
func GetMaxItemID(url string) (int, error) {
	reqURL := fmt.Sprintf("%s/%s", url, "maxitem") + JSONSuffix

	resp, err := GetBody(reqURL)

	var id int

	if err != nil {
		return id, err
	}

	err = json.Unmarshal(resp, &id)
	return id, err
}

func matchRegexFromBody(url string, regex string, cookie *http.Cookie) (string, error) {
	resp, err := GetBodyWithCookie(url, cookie)

	if err != nil {
		return "", err
	}

	r := regexp.MustCompile(regex)

	result := r.FindStringSubmatch(string(resp))

	if len(result) == 2 {
		return result[1], nil
	}

	return "", ErrFetchingActionURL
}

func vote(id int, cookie *http.Cookie, url string, voteType string) (bool, error) {
	reqURL := fmt.Sprintf("%s/%s?id=%d", url, "item", id)
	upvoteRegex := fmt.Sprintf(voteURLRegex, voteType, id)

	voteAuth, err := matchRegexFromBody(reqURL, upvoteRegex, cookie)

	if err != nil {
		return false, err
	}

	voteURL := fmt.Sprintf("%s/%s", url, voteAuth)
	unescaped := html.UnescapeString(voteURL)

	resp, err := GetBodyWithCookie(unescaped, cookie)

	if err == nil && resp != nil {
		return true, nil
	}

	return false, err
}

// Upvote : Upvote an item given an id and a cookie
func Upvote(id int, cookie *http.Cookie, url string) (bool, error) {
	return vote(id, cookie, url, "up")
}

// Unvote : Unvote a comment given an id and a cookie
func Unvote(id int, cookie *http.Cookie, url string) (bool, error) {
	return vote(id, cookie, url, "un")
}

// Comment : Create a comment on an item given an id and content
func Comment(id int, content string, cookie *http.Cookie, url string) (bool, error) {
	if len(content) == 0 {
		return false, ErrEmptyContent
	}

	reqURL := fmt.Sprintf("%s/%s?id=%d", url, "item", id)

	commentAuth, err := matchRegexFromBody(reqURL, commentURLRegex, cookie)

	if err != nil {
		return false, err
	}

	commentURL := fmt.Sprintf("%s/%s", url, "comment")

	body := NetURL.Values{}
	body.Set("parent", fmt.Sprintf("%d", id))
	body.Set("goto", fmt.Sprintf("item?id=%d", id))
	body.Set("hmac", commentAuth)
	body.Set("text", content)

	resp, err := PostWithCookie(commentURL, body, cookie)

	if err == nil && resp != nil {
		return true, nil
	}

	return false, err
}
