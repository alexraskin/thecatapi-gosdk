package thecatapi

import (
	"context"
	"net/http"
	"time"

	"github.com/alexraskin/thecatapi/internal/httpclient"
)

type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

type ClientOptions func(*Client)

func defaultRequestOptions(c *Client) httpclient.RequestOptions {
	return httpclient.RequestOptions{
		Ctx:     context.Background(),
		BaseURL: c.baseURL,
		APIKey:  c.apiKey,
		Client:  c.httpClient,
		Method:  "GET",
		Query:   nil,
		Body:    nil,
	}
}

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
