package trakt

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	libraryVersion = "0.1"
	userAgent      = "go-trakt/" + libraryVersion
	defaultBaseURL = "https://api-v2launch.trakt.tv"
)

type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Base URL for API requests. Defaults to https://api-v2launch.trakt.tv/,
	// but can be set to the sandbox url. BaseURL should always end with a
	// trailing slash.
	BaseURL *url.URL

	// User agent used when communicating with the trakt.tv API.
	UserAgent string

	// Client ID used for identifying your application
	ClientID string

	Calendar *CalendarService
}

func NewClient(clientID string, client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: client, BaseURL: baseURL, UserAgent: userAgent, ClientID: clientID}
	c.Calendar = &CalendarService{client: c}

	return c
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if c.ClientID == "" {
		return nil, errors.New("c.ClientID is empty")
	}

	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("trakt-api-version", "2")
	req.Header.Set("trakt-api-key", c.ClientID)
	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")

	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	return resp, err
}

type Error struct {
	Code        int
	Description string
}

func (e Error) Error() string {
	return fmt.Sprintf("trakt: %s (%d)", e.Description, e.Code)
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	var desc string
	switch r.StatusCode {
	case 400:
		desc = "Bad Request - request couldn't be parsed"
	case 401:
		desc = "Unauthorized - OAuth must be provided"
	case 403:
		desc = "Forbidden - invalid API key or unapproved app"
	case 404:
		desc = "Not Found - method exists, but no record found"
	case 405:
		desc = "Method Not Found - method doesn't exist"
	case 409:
		desc = "Conflict - resource already created"
	case 412:
		desc = "Precondition Failed - use application/json content type"
	case 422:
		desc = "Unprocessable Entity - validation errors"
	case 429:
		desc = "Rate Limit Exceeded"
	case 500:
		desc = "Server Error"
	case 503:
		desc = "Service Unavailable - server overloaded"
	default:
		desc = "Uknown error"
	}

	return Error{Code: r.StatusCode, Description: desc}
}
