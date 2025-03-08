package thecatapi

import (
	"net/http"
	"time"
)

type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

type ClientOptions func(*Client)

func defaultClient() *Client {
	return &Client{
		baseURL: "https://api.thecatapi.com/v1",
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

func WithAPIKey(apiKey string) ClientOptions {
	return func(c *Client) {
		c.apiKey = apiKey
	}
}

func WithBaseURL(url string) ClientOptions {
	return func(c *Client) {
		c.baseURL = url
	}
}

func WithHTTPClient(httpClient *http.Client) ClientOptions {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

func NewClient(opts ...ClientOptions) *Client {

	c := defaultClient()

	for _, fn := range opts {
		fn(c)
	}

	return c
}
