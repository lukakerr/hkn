package hkn

import (
	"encoding/json"
	"fmt"
	"net/http"
	NetURL "net/url"
)

// User : Struct representing a Hacker News user
type User struct {
	ID        string `json:"id"`
	Delay     int    `json:"delay"`
	Created   int32  `json:"created"`
	Karma     int    `json:"karma"`
	About     string `json:"about"`
	Submitted []int  `json:"submitted"`
}

// Get a user given an id
func getUser(id string, url string) (User, error) {
	reqURL := fmt.Sprintf("%s/%s/%s", url, "user", id) + jsonSuffix

	resp, err := getBody(reqURL)

	var user User

	if err != nil {
		return user, err
	}

	err = json.Unmarshal(resp, &user)
	return user, err
}

// Login a user given a username and password
func login(username string, password string, url string) (*http.Cookie, error) {
	reqURL := fmt.Sprintf("%s/%s", url, "login")

	body := NetURL.Values{}
	body.Set("acct", username)
	body.Set("pw", password)
	body.Set("goto", "news")

	cookie, err := postAndGetCookie(reqURL, body)

	return cookie, err
}
