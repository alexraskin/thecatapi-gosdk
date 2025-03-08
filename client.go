package thecatapi

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/alexraskin/thecatapi/internal/httpclient"
)

type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

type ClientOptions func(*Client)

func newRequestOptions(c *Client, path string, query url.Values, body io.Reader, result any) httpclient.RequestOptions {
	return httpclient.RequestOptions{
		Ctx:         context.Background(),
		BaseURL:     c.baseURL,
		APIKey:      c.apiKey,
		Client:      c.httpClient,
		Method:      "GET",
		Path:        path,
		Query:       query,
		Body:        body,
		Result:      result,
		ContentType: "application/json",
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

// Client is a struct that provides methods to interact with The Cat API.
// It allows users to perform various operations such as searching for cat images, retrieving cat breeds, and uploading images.
//
// Fields:
//
//	apiKey - The API key used for authenticating requests to The Cat API.
//	baseURL - The base URL for The Cat API endpoints.
//	httpClient - The HTTP client used to make requests.
//
// Methods:
//
//	NewClient - Initializes a new Client with optional configurations such as API key, base URL, and custom HTTP client.
//	GetBreeds - Retrieves a list of cat breeds with optional query parameters.
//	GetCatFacts - Retrieves a list of cat facts with optional query parameters.
//	SearchCats - Searches for cat images based on specified search parameters.
//	UploadImage - Uploads an image to The Cat API with optional upload parameters.
//	GetYourCatImages - Retrieves a list of your cat images based on specified query parameters.
//
// Example usage:
//
//	client := thecatapi.NewClient(thecatapi.WithAPIKey("YOUR-API-KEY"))
//	breeds, err := client.GetBreeds(thecatapi.WithBreedLimit(5))
//	if err != nil {
//	    log.Fatalf("Error fetching breeds: %v", err)
//	}
//	for _, breed := range breeds {
//	    fmt.Printf("Breed: %s, Origin: %s\n", breed.Name, breed.Origin)
//	}
func NewClient(opts ...ClientOptions) *Client {
	c := defaultClient()
	for _, fn := range opts {
		fn(c)
	}
	return c
}
