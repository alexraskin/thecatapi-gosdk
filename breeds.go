package thecatapi

import (
	"net/url"
	"strconv"

	"github.com/alexraskin/thecatapi/internal/httpclient"
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

// GetBreeds retrieves a list of cat breeds from The Cat API.
// It allows customization of the request through functional options.
//
// Parameters:
//
//	opts - A variadic list of CatBreedOptions functions that modify the request parameters.
//	       These options can be used to set pagination and other query parameters.
//
// Returns:
//
//	*[]CatBreedResponse - A pointer to a slice of CatBreedResponse structs containing information about each breed.
//	error - An error if the request fails or if there is an issue with the response.
//
// Example usage:
//
//	breeds, err := client.GetBreeds(thecatapi.WithBreedLimit(5))
//	if err != nil {
//	    log.Fatalf("Error fetching breeds: %v", err)
//	}
//	for _, breed := range *breeds {
//	    fmt.Printf("Breed: %s, Origin: %s\n", breed.Name, breed.Origin)
//	}
func (c *Client) GetBreeds(opts ...CatBreedOptions) (*[]CatBreedResponse, error) {
	params := defaultBreedParams()

	for _, opt := range opts {
		opt(&params)
	}

	query := params.toURLValues()

	var breeds []CatBreedResponse

	err := httpclient.DoRequest(newRequestOptions(c, "/breeds", query, nil, &breeds))

	if err != nil {
		return nil, err
	}

	return &breeds, nil
}
