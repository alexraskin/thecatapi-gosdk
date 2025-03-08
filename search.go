package thecatapi

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/alexraskin/thecatapi/internal/httpclient"
)

type CatImageSearchOptions func(*CatImageSearchParams)

func defaultImageSearchParams() CatImageSearchParams {
	return CatImageSearchParams{
		Page:      1,
		Limit:     1,
		Size:      SizeFull,
		Format:    FormatJSON,
		Order:     OrderRandom,
		HasBreeds: false,
	}
}

func WithImageSearchPage(page int) CatImageSearchOptions {
	return func(params *CatImageSearchParams) {
		params.Page = page
	}
}

func WithImageSearchLimit(limit int) CatImageSearchOptions {
	return func(params *CatImageSearchParams) {
		params.Limit = limit
	}
}

func WithImageSearchSize(size ImageSize) CatImageSearchOptions {
	return func(params *CatImageSearchParams) {
		params.Size = size
	}
}

func WithImageSearchMimeTypes(mimeTypes []string) CatImageSearchOptions {
	return func(params *CatImageSearchParams) {
		params.MimeTypes = mimeTypes
	}
}

func WithImageSearchFormat(format Format) CatImageSearchOptions {
	return func(params *CatImageSearchParams) {
		params.Format = format
	}
}

func WithImageSearchHasBreeds(hasBreeds bool) CatImageSearchOptions {
	return func(params *CatImageSearchParams) {
		params.HasBreeds = hasBreeds
	}
}

func WithImageSearchOrder(order OrderType) CatImageSearchOptions {
	return func(params *CatImageSearchParams) {
		params.Order = order
	}
}

func (p *CatImageSearchParams) toURLValues() url.Values {
	values := url.Values{}
	if p.Page > 0 {
		values.Add("page", strconv.Itoa(p.Page))
	}
	if p.Limit > 0 {
		values.Add("limit", strconv.Itoa(p.Limit))
	}
	if p.Size != "" {
		values.Add("size", string(p.Size))
	}
	if len(p.MimeTypes) > 0 {
		values.Add("mime_types", strings.Join(p.MimeTypes, ","))
	}
	if p.Format != "" {
		values.Add("format", string(p.Format))
	}
	if p.HasBreeds {
		values.Add("has_breeds", "true")
	}
	if p.Order != "" {
		values.Add("order", string(p.Order))
	}
	return values
}

// SearchCats retrieves a list of cat images from The Cat API based on the specified search parameters.
// It allows customization of the request through functional options.
//
// Parameters:
//
//	opts - A variadic list of CatImageSearchOptions functions that modify the search parameters.
//	       These options can be used to set filters such as size, format, order, and pagination.
//
// Returns:
//
//	*[]CatImageSearchResponse - A pointer to a slice of CatImageSearchResponse structs containing information about each cat image.
//	error - An error if the request fails or if there is an issue with the response.
//
// Example usage:
//
//	cats, err := client.SearchCats(thecatapi.WithImageSearchLimit(5), thecatapi.WithImageSearchSize(thecatapi.SizeSmall))
//	if err != nil {
//	    log.Fatalf("Error searching for cats: %v", err)
//	}
//	for _, cat := range *cats {
//	    fmt.Printf("Cat ID: %s, URL: %s\n", cat.ID, cat.URL)
//	}
func (c *Client) SearchCats(opts ...CatImageSearchOptions) (*[]CatImageSearchResponse, error) {
	params := defaultImageSearchParams()

	for _, fn := range opts {
		fn(&params)
	}

	query := params.toURLValues()

	var cats []CatImageSearchResponse

	requestOpts := newRequestOptions(c, "/images/search", query, nil, &cats)

	err := httpclient.DoRequest(requestOpts)

	if err != nil {
		return nil, err
	}

	return &cats, nil
}
