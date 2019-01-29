package hkn

import (
	"net/http"
)

type Client struct {
	BaseURL string
	WebURL  string
}

func NewClient() *Client {
	c := Client{
		BaseURL: "https://hacker-news.firebaseio.com/v0",
		WebURL:  "https://news.ycombinator.com",
	}

	return &c
}

// Get a single item given an id
func (c *Client) GetItem(id int) (Item, error) {
	return GetItem(id, c.BaseURL)
}

// Get a slice of items given a number of ids
func (c *Client) GetItems(ids []int) (Items, error) {
	return GetItems(ids, c.BaseURL)
}

// Get the most recent item id
func (c *Client) GetMaxItemId() (int, error) {
	return GetMaxItemId(c.BaseURL)
}

// Get the latest item and profile updates
func (c *Client) GetUpdates() (Updates, error) {
	return GetUpdates(c.BaseURL)
}

// Get a user given an id
func (c *Client) GetUser(id string) (User, error) {
	return GetUser(id, c.BaseURL)
}

// Login a user given a username and password
func (c *Client) Login(username string, password string) (*http.Cookie, error) {
	return Login(username, password, c.WebURL)
}

// Get top stories given a number
func (c *Client) GetTopStories(number int) ([]int, error) {
	return GetTopStories(number, c.BaseURL)
}

// Get new stories given a number
func (c *Client) GetNewStories(number int) ([]int, error) {
	return GetNewStories(number, c.BaseURL)
}

// Get best stories given a number
func (c *Client) GetBestStories(number int) ([]int, error) {
	return GetBestStories(number, c.BaseURL)
}

// Get latest ask stories given a number
func (c *Client) GetLatestAskStories(number int) ([]int, error) {
	return GetLatestAskStories(number, c.BaseURL)
}

// Get latest show stories given a number
func (c *Client) GetLatestShowStories(number int) ([]int, error) {
	return GetLatestShowStories(number, c.BaseURL)
}

// Get latest job stories given a number
func (c *Client) GetLatestJobStories(number int) ([]int, error) {
	return GetLatestJobStories(number, c.BaseURL)
}
