package hkn

import (
	"net/http"
)

// Client : Struct representing the hkn client
type Client struct {
	BaseURL string
	WebURL  string
}

// NewClient : Create a new hkn client
func NewClient() *Client {
	c := Client{
		BaseURL: "https://hacker-news.firebaseio.com/v0",
		WebURL:  "https://news.ycombinator.com",
	}

	return &c
}

const (
	// JSONSuffix : A JSON suffix for URLs
	JSONSuffix = ".json"
)

// GetItem : Get a single item given an id
func (c *Client) GetItem(id int) (Item, error) {
	return GetItem(id, c.BaseURL)
}

// GetItems : Get a slice of items given a number of ids
func (c *Client) GetItems(ids []int) (Items, error) {
	return GetItems(ids, c.BaseURL)
}

// GetMaxItemID : Get the most recent item id
func (c *Client) GetMaxItemID() (int, error) {
	return GetMaxItemID(c.BaseURL)
}

// GetUpdates : Get the latest item and profile updates
func (c *Client) GetUpdates() (Updates, error) {
	return GetUpdates(c.BaseURL)
}

// GetUser : Get a user given an id
func (c *Client) GetUser(id string) (User, error) {
	return GetUser(id, c.BaseURL)
}

// Login : Login a user given a username and password
func (c *Client) Login(username string, password string) (*http.Cookie, error) {
	return Login(username, password, c.WebURL)
}

// GetTopStories : Get top stories given a number
func (c *Client) GetTopStories(number int) ([]int, error) {
	return GetTopStories(number, c.BaseURL)
}

// GetNewStories : Get new stories given a number
func (c *Client) GetNewStories(number int) ([]int, error) {
	return GetNewStories(number, c.BaseURL)
}

// GetBestStories : Get best stories given a number
func (c *Client) GetBestStories(number int) ([]int, error) {
	return GetBestStories(number, c.BaseURL)
}

// GetLatestAskStories : Get latest ask stories given a number
func (c *Client) GetLatestAskStories(number int) ([]int, error) {
	return GetLatestAskStories(number, c.BaseURL)
}

// GetLatestShowStories : Get latest show stories given a number
func (c *Client) GetLatestShowStories(number int) ([]int, error) {
	return GetLatestShowStories(number, c.BaseURL)
}

// GetLatestJobStories : Get latest job stories given a number
func (c *Client) GetLatestJobStories(number int) ([]int, error) {
	return GetLatestJobStories(number, c.BaseURL)
}

// Upvote : Upvote an item given an id
func (c *Client) Upvote(id int, cookie *http.Cookie) (bool, error) {
	return Upvote(id, cookie, c.WebURL)
}
