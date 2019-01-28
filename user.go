package hkn

import (
	"encoding/json"
	"fmt"
	"net/http"
	NetUrl "net/url"
)

type User struct {
	Id        string `json:"id"`
	Delay     int    `json:"delay"`
	Created   int32  `json:"created"`
	Karma     int    `json:"karma"`
	About     string `json:"about"`
	Submitted []int  `json:"submitted"`
}

// Get a user given an id
func GetUser(id string, url string) (User, error) {
	reqUrl := fmt.Sprintf("%s/%s/%s", url, "user", id)

	resp, err := GetBody(reqUrl)

	var user User

	if err != nil {
		return user, err
	}

	err = json.Unmarshal(resp, &user)
	return user, err
}

// Login a user given a username and password
func Login(username string, password string, url string) (*http.Cookie, error) {
	reqUrl := fmt.Sprintf("%s/%s", url, "login")

	body := NetUrl.Values{}
	body.Set("acct", username)
	body.Set("pw", password)
	body.Set("goto", "news")

	cookie, err := PostAndGetCookie(reqUrl, body)

	return cookie, err
}
