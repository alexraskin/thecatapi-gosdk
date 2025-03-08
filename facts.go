package thecatapi

import (
	"net/url"
	"strconv"

	"github.com/alexraskin/thecatapi/internal/httpclient"
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

// GetCatFacts retrieves a list of cat facts from The Cat API.
// It allows customization of the request through functional options.
//
// Parameters:
//
//	opts - A variadic list of CatFactsOptions functions that modify the request parameters.
//	       These options can be used to set pagination and other query parameters.
//
// Returns:
//
//	*CatFactsResponse - A pointer to a CatFactsResponse struct containing the cat facts.
//	error - An error if the request fails or if there is an issue with the response.
//
// Example usage:
//
//	facts, err := client.GetCatFacts(thecatapi.WithCatFactsLimit(5))
//	if err != nil {
//	    log.Fatalf("Error fetching cat facts: %v", err)
//	}
//	fmt.Printf("Cat Fact: %s\n", facts.Fact)
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

	err = httpclient.DoRequest(newRequestOptions(c, "/facts", values, nil, &response))

	if err != nil {
		return nil, err
	}

	return &response, nil
}
