package hkn

import (
	"errors"
	"io/ioutil"
	"net/http"
	NetUrl "net/url"
	"strings"
)

const SUFFIX = ".json"

// Perform a GET request and return the response
func Get(url string) (*http.Response, error) {
	reqUrl := url + SUFFIX

	// Build the request
	req, err := http.NewRequest("GET", reqUrl, nil)
	req.Close = true

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

// Perform a POST request and return the response
func Post(url string, urlEncodedValues NetUrl.Values) (*http.Response, error) {
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

// Perform a GET request and return the body as a slice of bytes
func GetBody(url string) ([]byte, error) {
	resp, err := Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

// Perform a POST request and return the first cookie in the response
func PostAndGetCookie(url string, urlEncodedValues NetUrl.Values) (*http.Cookie, error) {
	resp, err := Post(url, urlEncodedValues)

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
