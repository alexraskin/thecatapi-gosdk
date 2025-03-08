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

func WithSize(size ImageSize) CatImageSearchOptions {
	return func(params *CatImageSearchParams) {
		params.Size = size
	}
}

func WithMimeTypes(mimeTypes []string) CatImageSearchOptions {
	return func(params *CatImageSearchParams) {
		params.MimeTypes = mimeTypes
	}
}

func WithFormat(format Format) CatImageSearchOptions {
	return func(params *CatImageSearchParams) {
		params.Format = format
	}
}

func WithHasBreeds(hasBreeds bool) CatImageSearchOptions {
	return func(params *CatImageSearchParams) {
		params.HasBreeds = hasBreeds
	}
}

func WithOrder(order OrderType) CatImageSearchOptions {
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

func (c *Client) SearchCats(opts ...CatImageSearchOptions) ([]CatImageSearchResponse, error) {
	params := defaultImageSearchParams()

	for _, fn := range opts {
		fn(&params)
	}

	query := params.toURLValues()

	var cats []CatImageSearchResponse

	requestOpts := defaultRequestOptions(c)
	requestOpts.Path = "/images/search"
	requestOpts.Query = query
	requestOpts.Result = &cats
	requestOpts.ContentType = "application/json"

	err := httpclient.DoRequest(requestOpts)

	if err != nil {
		return nil, err
	}

	return cats, nil
}
