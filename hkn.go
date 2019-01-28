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
	item, err := GetItem(id, c.BaseURL)

	return item, err
}

// Get a slice of items given a number of ids
func (c *Client) GetItems(ids []int) (Items, error) {
	items, err := GetItems(ids, c.BaseURL)

	return items, err
}

// Get a user given an id
func (c *Client) GetUser(id string) (User, error) {
	user, err := GetUser(id, c.BaseURL)

	return user, err
}

// Login a user given a username and password
func (c *Client) Login(username string, password string) (*http.Cookie, error) {
	cookie, err := Login(username, password, c.WebURL)

	return cookie, err
}
