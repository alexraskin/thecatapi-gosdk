package thecatapi

import (
	"context"
	"errors"
	"net/url"
	"strconv"

	"github.com/alexraskin/thecatapi/internal/httpclient"
)

type CatByIDImageOption func(*CatByIDImageParams)

func WithCatImageID(id string) CatByIDImageOption {
	return func(params *CatByIDImageParams) {
		params.ID = id
	}
}

func (c *Client) GetCatImageByID(opts ...CatByIDImageOption) (*CatByIDImageResponse, error) {
	params := CatByIDImageParams{}
	for _, fn := range opts {
		fn(&params)
	}

	if params.ID == "" {
		return nil, errors.New("image ID is required; use WithCatImageID")
	}

	var response CatByIDImageResponse

	requestOpts := httpclient.RequestOptions{
		Ctx:         context.Background(),
		BaseURL:     c.baseURL,
		APIKey:      c.apiKey,
		Client:      c.httpClient,
		Method:      "GET",
		Path:        "/images/" + params.ID,
		Query:       nil,
		Body:        nil,
		ContentType: "application/json",
		Result:      &response,
	}

	err := httpclient.DoRequest(requestOpts)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

type YourCatImagesOption func(*YourCatImagesQueryParams)

func defaultYourCatImagesQueryParams() YourCatImagesQueryParams {
	return YourCatImagesQueryParams{
		Limit: 10,
	}
}

func WithYourCatImagesLimit(limit int) YourCatImagesOption {
	return func(params *YourCatImagesQueryParams) {
		params.Limit = limit
	}
}

func WithYourCatImagesPage(page int) YourCatImagesOption {
	return func(params *YourCatImagesQueryParams) {
		params.Page = page
	}
}

func WithYourCatImagesOrder(order string) YourCatImagesOption {
	return func(params *YourCatImagesQueryParams) {
		params.Order = order
	}
}

func WithYourCatImagesSubID(subID string) YourCatImagesOption {
	return func(params *YourCatImagesQueryParams) {
		params.SubID = subID
	}
}

func WithYourCatImagesBreedIDs(breedIDs string) YourCatImagesOption {
	return func(params *YourCatImagesQueryParams) {
		params.BreedIDs = breedIDs
	}
}

func WithYourCatImagesCategoryIDs(categoryIDs string) YourCatImagesOption {
	return func(params *YourCatImagesQueryParams) {
		params.CategoryIDs = categoryIDs
	}
}

func WithYourCatImagesFormat(format string) YourCatImagesOption {
	return func(params *YourCatImagesQueryParams) {
		params.Format = format
	}
}

func WithYourCatImagesOriginalFilename(originalFilename string) YourCatImagesOption {
	return func(params *YourCatImagesQueryParams) {
		params.OriginalFilename = originalFilename
	}
}

func WithYourCatImagesUserID(userID string) YourCatImagesOption {
	return func(params *YourCatImagesQueryParams) {
		params.UserID = userID
	}
}

func (p *YourCatImagesQueryParams) toURLValues() (url.Values, error) {
	values := url.Values{}
	if p.Limit > 0 {
		values.Add("limit", strconv.Itoa(p.Limit))
	}
	if p.Limit > 10 {
		return nil, errors.New("limit must be less than or equal to 10")
	}
	if p.Page > 0 {
		values.Add("page", strconv.Itoa(p.Page))
	}
	if p.Order != "" {
		values.Add("order", p.Order)
	}
	if p.SubID != "" {
		values.Add("sub_id", p.SubID)
	}
	if p.BreedIDs != "" {
		values.Add("breed_ids", p.BreedIDs)
	}
	if p.CategoryIDs != "" {
		values.Add("category_ids", p.CategoryIDs)
	}
	if p.Format != "" {
		values.Add("format", p.Format)
	}
	if p.OriginalFilename != "" {
		values.Add("original_filename", p.OriginalFilename)
	}
	if p.UserID != "" {
		values.Add("user_id", p.UserID)
	}
	return values, nil
}

func (c *Client) GetYourCatImages(opts ...YourCatImagesOption) (*YourCatImagesResponse, error) {
	params := defaultYourCatImagesQueryParams()

	for _, fn := range opts {
		fn(&params)
	}

	values, err := params.toURLValues()
	if err != nil {
		return nil, err
	}

	var response YourCatImagesResponse

	requestOpts := defaultRequestOptions(c)
	requestOpts.Path = "/images/search"
	requestOpts.Query = values
	requestOpts.Result = &response
	requestOpts.ContentType = "application/json"

	err = httpclient.DoRequest(requestOpts)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
