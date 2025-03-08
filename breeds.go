package thecatapi

import (
	"net/url"
	"strconv"
)

type CatBreedOptions func(*CatBreedParams)

func defaultBreedParams() CatBreedParams {
	return CatBreedParams{
		Page:  1,
		Limit: 10,
	}
}

func WithBreedPage(page int) CatBreedOptions {
	return func(params *CatBreedParams) {
		params.Page = page
	}
}

func WithBreedLimit(limit int) CatBreedOptions {
	return func(params *CatBreedParams) {
		params.Limit = limit
	}
}

func (p *CatBreedParams) toURLValues() url.Values {
	values := url.Values{}
	if p.Page > 0 {
		values.Add("page", strconv.Itoa(p.Page))
	}
	if p.Limit > 0 {
		values.Add("limit", strconv.Itoa(p.Limit))
	}
	return values
}

func (c *Client) GetBreeds(opts ...CatBreedOptions) ([]CatBreedResponse, error) {
	params := defaultBreedParams()

	for _, opt := range opts {
		opt(&params)
	}

	query := params.toURLValues()

	var breeds []CatBreedResponse

	err := c.doRequest("GET", "/breeds", query, &breeds)

	if err != nil {
		return nil, err
	}

	return breeds, nil
}
