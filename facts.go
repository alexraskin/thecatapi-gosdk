package thecatapi

import (
	"net/url"
	"strconv"
)

type CatFactsOptions func(*CatFactsParams)

func defaultCatFactsParams() CatFactsParams {
	return CatFactsParams{
		Limit: 10,
		Page:  1,
		Order: OrderRandom,
	}
}

func WithCatFactsLimit(limit int) CatFactsOptions {
	return func(params *CatFactsParams) {
		params.Limit = limit
	}
}

func WithCatFactsPage(page int) CatFactsOptions {
	return func(params *CatFactsParams) {
		params.Page = page
	}
}

func WithCatFactsOrder(order OrderType) CatFactsOptions {
	return func(params *CatFactsParams) {
		params.Order = order
	}
}

func (p *CatFactsParams) toURLValues() (url.Values, error) {
	values := url.Values{}
	if p.Limit > 0 {
		values.Add("limit", strconv.Itoa(p.Limit))
	}
	if p.Page > 0 {
		values.Add("page", strconv.Itoa(p.Page))
	}

	if p.Order != "" {
		values.Add("order", string(p.Order))
	}
	return values, nil
}

func (c *Client) GetCatFacts(opts ...CatFactsOptions) (*CatFactsResponse, error) {
	params := defaultCatFactsParams()

	for _, fn := range opts {
		fn(&params)
	}

	values, err := params.toURLValues()
	if err != nil {
		return nil, err
	}

	var response CatFactsResponse

	err = c.doRequest("GET", "/facts", values, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
