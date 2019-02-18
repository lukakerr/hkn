package hkn

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

// Create a new HTTP client with sensible options
func newHTTPClient() *http.Client {
	return &http.Client{
		// A 10 second timeout
		Timeout: time.Second * 10,

		// Don't follow 301 redirects
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
}

// Perform a GET request and return the response
func get(resource string, cookie *http.Cookie) (*http.Response, error) {
	// Build the request
	req, err := http.NewRequest("GET", resource, nil)
	req.Close = true

	if cookie != nil {
		req.AddCookie(cookie)
	}

	if err != nil {
		return nil, ErrFetching
	}

	client := newHTTPClient()

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

// Perform a GET request and return the body as a slice of bytes
func getBody(resource string) ([]byte, error) {
	resp, err := get(resource, nil)

	if err != nil {
		return nil, err
	}

	return getContent(resp)
}

// Perform a GET request with a cookie and return the body as a slice of bytes
func getBodyWithCookie(resource string, cookie *http.Cookie) ([]byte, error) {
	resp, err := get(resource, cookie)

	if err != nil {
		return nil, err
	}

	return getContent(resp)
}

// Perform a POST request and return the response
func post(resource string, urlEncodedValues url.Values, cookie *http.Cookie) (*http.Response, error) {
	req, err := http.NewRequest("POST", resource, strings.NewReader(urlEncodedValues.Encode()))

	if err != nil {
		return nil, err
	}

	req.Close = true

	if cookie != nil {
		req.AddCookie(cookie)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Access-Control-Allow-Origin", "*")

	client := newHTTPClient()

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Perform a POST request with a cookie
func postWithCookie(resource string, urlEncodedValues url.Values, cookie *http.Cookie) ([]byte, error) {
	resp, err := post(resource, urlEncodedValues, cookie)

	if err != nil {
		return nil, err
	}

	return getContent(resp)
}

// Perform a POST request and return the first cookie in the response
func postAndGetCookie(resource string, urlEncodedValues url.Values) (*http.Cookie, error) {
	resp, err := post(resource, urlEncodedValues, nil)

	if err != nil {
		return &http.Cookie{}, err
	}

	defer resp.Body.Close()

	cookies := resp.Cookies()

	if len(cookies) == 0 {
		return &http.Cookie{}, ErrInvalidAuth
	}

	return cookies[0], nil
}

func matchRegexFromBody(webURL string, regex string, cookie *http.Cookie) (string, error) {
	resp, err := getBodyWithCookie(webURL, cookie)

	if err != nil {
		return "", err
	}

	r := regexp.MustCompile(regex)

	result := r.FindStringSubmatch(string(resp))

	if len(result) == 2 {
		return result[1], nil
	}

	return "", ErrFetchingActionURL
}
