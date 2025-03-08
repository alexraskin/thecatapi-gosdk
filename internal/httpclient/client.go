package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type RequestOptions struct {
	Ctx         context.Context
	BaseURL     string
	APIKey      string
	Client      *http.Client
	Method      string
	Path        string
	Query       url.Values
	Body        io.Reader
	ContentType string
	Result      any
}

func DoRequest(opts RequestOptions) error {
	var reqURL string
	if opts.Query != nil {
		reqURL = fmt.Sprintf("%s%s?%s", opts.BaseURL, opts.Path, opts.Query.Encode())
	} else {
		reqURL = fmt.Sprintf("%s%s", opts.BaseURL, opts.Path)
	}

	fmt.Println(reqURL)

	req, err := http.NewRequestWithContext(opts.Ctx, opts.Method, reqURL, opts.Body)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	if opts.ContentType != "" {
		req.Header.Set("Content-Type", opts.ContentType)
	}

	if opts.APIKey != "" {
		req.Header.Add("x-api-key", opts.APIKey)
	}

	resp, err := opts.Client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	if opts.Result != nil {
		if err := json.NewDecoder(resp.Body).Decode(opts.Result); err != nil {
			return fmt.Errorf("error decoding response: %v", err)
		}
	}

	return nil
}
