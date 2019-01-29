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

// Get a slice of items given a number of ids
func (c *Client) GetMaxItemId() (int, error) {
	return GetMaxItemId(c.BaseURL)
}

// Get a user given an id
func (c *Client) GetUser(id string) (User, error) {
	return GetUser(id, c.BaseURL)
}

// Login a user given a username and password
func (c *Client) Login(username string, password string) (*http.Cookie, error) {
	return Login(username, password, c.WebURL)
}
