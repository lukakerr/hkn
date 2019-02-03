/*
Package hkn is a go module for interacting with Hacker News.

To get started simply import the package, create a client and call methods on the client:

	import (
		"fmt"

		"github.com/lukakerr/hkn"
	)

	func main() {
		client := hkn.NewClient()

		// For example, to get an item by id
		item, err := client.GetItem(8869)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%+v\n", item)
	}
*/
package hkn

import (
	"errors"
	"net/http"
)

// Client represents the hkn client
type Client struct {
	BaseURL string
	WebURL  string
}

// NewClient creates a new hkn client
func NewClient() *Client {
	c := Client{
		BaseURL: "https://hacker-news.firebaseio.com/v0",
		WebURL:  "https://news.ycombinator.com",
	}

	return &c
}

const (
	// A JSON suffix for URLs
	jsonSuffix = ".json"
)

var (
	// ErrFetching represents an error fetching a resource
	ErrFetching = errors.New("fetching resource failed")

	// ErrEmptyContent represents an error that content provided is empty
	ErrEmptyContent = errors.New("content is empty")

	// ErrInvalidAuth represents an error authenticating
	ErrInvalidAuth = errors.New("invalid username or password")

	// ErrFetchingActionURL represents an error fetching an action URL
	ErrFetchingActionURL = errors.New("fetching action URL failed")

	// ErrInvalidNumber represents an error that a number provided is invalid
	ErrInvalidNumber = errors.New("invalid number")
)

// GetItem returns a single item given an id
func (c *Client) GetItem(id int) (Item, error) {
	return getItem(id, c.BaseURL)
}

// GetItems returns a slice of items given a number of ids
func (c *Client) GetItems(ids []int) (Items, error) {
	return getItems(ids, c.BaseURL)
}

// GetMaxItemID returns the most recent item id
func (c *Client) GetMaxItemID() (int, error) {
	return getMaxItemID(c.BaseURL)
}

// GetUpdates returns the latest item and profile updates
func (c *Client) GetUpdates() (Updates, error) {
	return getUpdates(c.BaseURL)
}

// GetUser returns a user given an id
func (c *Client) GetUser(id string) (User, error) {
	return getUser(id, c.BaseURL)
}

// Login a user given a username and password and get back an authentication cookie
func (c *Client) Login(username string, password string) (*http.Cookie, error) {
	return login(username, password, c.WebURL)
}

// GetTopStories returns top stories given a number
func (c *Client) GetTopStories(number int) ([]int, error) {
	return getTopStories(number, c.BaseURL)
}

// GetNewStories returns new stories given a number
func (c *Client) GetNewStories(number int) ([]int, error) {
	return getNewStories(number, c.BaseURL)
}

// GetBestStories returns best stories given a number
func (c *Client) GetBestStories(number int) ([]int, error) {
	return getBestStories(number, c.BaseURL)
}

// GetLatestAskStories returns latest ask stories given a number
func (c *Client) GetLatestAskStories(number int) ([]int, error) {
	return getLatestAskStories(number, c.BaseURL)
}

// GetLatestShowStories returns latest show stories given a number
func (c *Client) GetLatestShowStories(number int) ([]int, error) {
	return getLatestShowStories(number, c.BaseURL)
}

// GetLatestJobStories returns latest job stories given a number
func (c *Client) GetLatestJobStories(number int) ([]int, error) {
	return getLatestJobStories(number, c.BaseURL)
}

// Upvote an item given an id and cookie and get back whether the upvote was successful
func (c *Client) Upvote(id int, cookie *http.Cookie) (bool, error) {
	return upvote(id, cookie, c.WebURL)
}

// Unvote a comment given an id and cookie and get back whether the unvote was successful
func (c *Client) Unvote(id int, cookie *http.Cookie) (bool, error) {
	return unvote(id, cookie, c.WebURL)
}

// Comment creates a comment on an item given an id, content and cookie, and returns whether the comment was successful
func (c *Client) Comment(id int, content string, cookie *http.Cookie) (bool, error) {
	return comment(id, content, cookie, c.WebURL)
}
