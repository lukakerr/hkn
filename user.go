package hkn

import (
	"encoding/json"
	"fmt"
	"net/http"
	NetUrl "net/url"
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

// GetUser : Get a user given an id
func GetUser(id string, url string) (User, error) {
	reqURL := fmt.Sprintf("%s/%s/%s", url, "user", id) + JSONSuffix

	resp, err := GetBody(reqURL)

	var user User

	if err != nil {
		return user, err
	}

	err = json.Unmarshal(resp, &user)
	return user, err
}

// Login : Login a user given a username and password
func Login(username string, password string, url string) (*http.Cookie, error) {
	reqURL := fmt.Sprintf("%s/%s", url, "login")

	body := NetUrl.Values{}
	body.Set("acct", username)
	body.Set("pw", password)
	body.Set("goto", "news")

	cookie, err := PostAndGetCookie(reqURL, body)

	return cookie, err
}
