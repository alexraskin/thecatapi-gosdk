// Package thecatapi provides a Go client for The Cat API (https://thecatapi.com/),
// a service that offers a wide variety of cat images, breeds, and related information.
//
// Getting Started
//
// To use this package, first create a client:
//
//	// Create client with default settings (no API key)
//	client := thecatapi.NewClient()
//
//	// Or with an API key (recommended for more than 10 images per request)
//	client := thecatapi.NewClient(thecatapi.WithAPIKey("YOUR-API-KEY"))
//
//	// With custom options
//	client := thecatapi.NewClient(
//		thecatapi.WithAPIKey("YOUR-API-KEY"),
//		thecatapi.WithBaseURL("https://custom-url.com"),
//		thecatapi.WithHTTPClient(&http.Client{Timeout: 10 * time.Second}),
//	)
//
// Searching for Cat Images
//
// The most common operation is searching for cat images using functional options:
//
//	// Basic search with default parameters
//	cats, err := client.SearchCats()
//
//	// Search with specific options
//	cats, err := client.SearchCats(
//		thecatapi.WithSize(thecatapi.SizeSmall),
//		thecatapi.WithMimeTypes([]string{"jpg"}),
//		thecatapi.WithLimit(10),
//		thecatapi.WithHasBreeds(true),
//		thecatapi.WithOrder(thecatapi.OrderRandom),
//		thecatapi.WithIncludeBreeds(true),
//		thecatapi.WithIncludeCategories(true),
//	)
//
// API Key Requirements
//
// While The Cat API can be used without an API key, there are limitations:
//   - Without an API key: Maximum 10 images per request
//   - With an API key: Up to 25 images per request and higher rate limits
//
// To obtain an API key, visit: https://thecatapi.com/signup
//
// Error Handling
//
// All methods return errors that should be checked:
//
//	cats, err := client.SearchCats(thecatapi.WithLimit(10))
//	if err != nil {
//		// Handle error
//		log.Fatalf("Error searching cats: %v", err)
//	}
//
// For more information about The Cat API, see: https://developers.thecatapi.com/
package thecatapi