package thecatapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) doRequest(method, path string, query url.Values, result any) error {

	var reqURL string

	if query != nil {
		reqURL = fmt.Sprintf("%s%s?%s", c.baseURL, path, query.Encode())
	} else {
		reqURL = fmt.Sprintf("%s%s", c.baseURL, path)
	}

	fmt.Println(reqURL)

	req, err := http.NewRequest(method, reqURL, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	if c.apiKey != "" {
		req.Header.Add("x-api-key", c.apiKey)
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("error decoding response: %v", err)
		}
	}

	return nil
}

func (c *Client) doRequestWithBody(method, path string, body io.Reader, contentType string, result any) error {
	reqURL := fmt.Sprintf("%s%s", c.baseURL, path)

	req, err := http.NewRequest(method, reqURL, body)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", contentType)

	if c.apiKey != "" {
		req.Header.Add("x-api-key", c.apiKey)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("error decoding response: %v", err)
		}
	}

	return nil
}
