package thecatapi

import (
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

// GetCatImageByID retrieves a specific cat image from The Cat API using the provided image ID.
// It allows customization of the request through functional options.
//
// Parameters:
//
//	opts - A variadic list of CatByIDImageOption functions that modify the request parameters.
//	       These options can be used to set the image ID and other query parameters.
//
// Returns:
//
//	*CatByIDImageResponse - A pointer to a CatByIDImageResponse struct containing information about the retrieved cat image.
//	error - An error if the request fails, if there is an issue with the response, or if the image ID is not provided.
//
// Example usage:
//
//	image, err := client.GetCatImageByID(thecatapi.WithCatImageID("abc123"))
//	if err != nil {
//	    log.Fatalf("Error fetching cat image: %v", err)
//	}
//	fmt.Printf("Image ID: %s, URL: %s\n", image.ID, image.URL)
func (c *Client) GetCatImageByID(opts ...CatByIDImageOption) (*CatByIDImageResponse, error) {
	params := CatByIDImageParams{}
	for _, fn := range opts {
		fn(&params)
	}

	if params.ID == "" {
		return nil, errors.New("image ID is required; use WithCatImageID")
	}

	var response CatByIDImageResponse

	err := httpclient.DoRequest(newRequestOptions(c, "/images/"+params.ID, nil, nil, &response))

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

// GetYourCatImages retrieves a list of your cat images from The Cat API based on the specified query parameters.
// It allows customization of the request through functional options.
//
// Parameters:
//
//	opts - A variadic list of YourCatImagesOption functions that modify the query parameters.
//	       These options can be used to set filters such as limit, page, order, and other query parameters.
//
// Returns:
//
//	*YourCatImagesResponse - A pointer to a YourCatImagesResponse struct containing information about the retrieved cat images.
//	error - An error if the request fails or if there is an issue with the response.
//
// Example usage:
//
//	images, err := client.GetYourCatImages(thecatapi.WithYourCatImagesLimit(5), thecatapi.WithYourCatImagesOrder("asc"))
//	if err != nil {
//	    log.Fatalf("Error fetching your cat images: %v", err)
//	}
//	for _, image := range images.Images {
//	    fmt.Printf("Image ID: %s, URL: %s\n", image.ID, image.URL)
//	}
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

	err = httpclient.DoRequest(newRequestOptions(c, "/images/search", values, nil, &response))
	if err != nil {
		return nil, err
	}

	return &response, nil
}
