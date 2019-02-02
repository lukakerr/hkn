package hkn

import (
	"errors"
	"io/ioutil"
	"net/http"
	NetUrl "net/url"
	"strings"
)

// Perform a GET request and return the response
func get(url string, cookie *http.Cookie) (*http.Response, error) {
	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	req.Close = true

	if cookie != nil {
		req.AddCookie(cookie)
	}

	if err != nil {
		return nil, errors.New("Error fetching repository")
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Get the content from a http response and close the response
func getContent(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

// GetBody : Perform a GET request and return the body as a slice of bytes
func GetBody(url string) ([]byte, error) {
	resp, err := get(url, nil)

	if err != nil {
		return nil, err
	}

	return getContent(resp)
}

// GetBodyWithCookie : Perform a GET request with a cookie and return the body as a slice of bytes
func GetBodyWithCookie(url string, cookie *http.Cookie) ([]byte, error) {
	resp, err := get(url, cookie)

	if err != nil {
		return nil, err
	}

	return getContent(resp)
}

// Perform a POST request and return the response
func post(url string, urlEncodedValues NetUrl.Values) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(urlEncodedValues.Encode()))
	req.Close = true

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Access-Control-Allow-Origin", "*")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// PostAndGetCookie : Perform a POST request and return the first cookie in the response
func PostAndGetCookie(url string, urlEncodedValues NetUrl.Values) (*http.Cookie, error) {
	resp, err := post(url, urlEncodedValues)

	if err != nil {
		return &http.Cookie{}, err
	}

	defer resp.Body.Close()

	cookies := resp.Cookies()

	if len(cookies) == 0 {
		return &http.Cookie{}, errors.New("Invalid username or password")
	}

	return cookies[0], nil
}
