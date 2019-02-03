package hkn

import (
	"errors"
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
	// A JSON suffix for URLs
	jsonSuffix = ".json"
)

var (
	// ErrFetching : An error fetching a resource
	ErrFetching = errors.New("Error fetching resource")

	// ErrEmptyContent : An error that content provided is empty
	ErrEmptyContent = errors.New("Content is empty")

	// ErrInvalidAuth : An error authenticating
	ErrInvalidAuth = errors.New("Invalid username or password")

	// ErrFetchingActionURL : An error fetching an action URL
	ErrFetchingActionURL = errors.New("Error fetching action URL")

	// ErrInvalidNumber : An error that a number provided is invalid
	ErrInvalidNumber = errors.New("Invalid number")
)

// GetItem : Get a single item given an id
func (c *Client) GetItem(id int) (Item, error) {
	return getItem(id, c.BaseURL)
}

// GetItems : Get a slice of items given a number of ids
func (c *Client) GetItems(ids []int) (Items, error) {
	return getItems(ids, c.BaseURL)
}

// GetMaxItemID : Get the most recent item id
func (c *Client) GetMaxItemID() (int, error) {
	return getMaxItemID(c.BaseURL)
}

// GetUpdates : Get the latest item and profile updates
func (c *Client) GetUpdates() (Updates, error) {
	return getUpdates(c.BaseURL)
}

// GetUser : Get a user given an id
func (c *Client) GetUser(id string) (User, error) {
	return getUser(id, c.BaseURL)
}

// Login : Login a user given a username and password
func (c *Client) Login(username string, password string) (*http.Cookie, error) {
	return login(username, password, c.WebURL)
}

// GetTopStories : Get top stories given a number
func (c *Client) GetTopStories(number int) ([]int, error) {
	return getTopStories(number, c.BaseURL)
}

// GetNewStories : Get new stories given a number
func (c *Client) GetNewStories(number int) ([]int, error) {
	return getNewStories(number, c.BaseURL)
}

// GetBestStories : Get best stories given a number
func (c *Client) GetBestStories(number int) ([]int, error) {
	return getBestStories(number, c.BaseURL)
}

// GetLatestAskStories : Get latest ask stories given a number
func (c *Client) GetLatestAskStories(number int) ([]int, error) {
	return getLatestAskStories(number, c.BaseURL)
}

// GetLatestShowStories : Get latest show stories given a number
func (c *Client) GetLatestShowStories(number int) ([]int, error) {
	return getLatestShowStories(number, c.BaseURL)
}

// GetLatestJobStories : Get latest job stories given a number
func (c *Client) GetLatestJobStories(number int) ([]int, error) {
	return getLatestJobStories(number, c.BaseURL)
}

// Upvote : Upvote an item given an id
func (c *Client) Upvote(id int, cookie *http.Cookie) (bool, error) {
	return upvote(id, cookie, c.WebURL)
}

// Unvote : Unvote a comment given an id
func (c *Client) Unvote(id int, cookie *http.Cookie) (bool, error) {
	return unvote(id, cookie, c.WebURL)
}

// Comment : Create a comment on an item given an id and content
func (c *Client) Comment(id int, content string, cookie *http.Cookie) (bool, error) {
	return comment(id, content, cookie, c.WebURL)
}
