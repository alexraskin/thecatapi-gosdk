package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alexraskin/thecatapi"
)

func main() {
	client := thecatapi.NewClient(thecatapi.WithAPIKey(os.Getenv("THECATAPI_API_KEY")))

	idLookup, err := client.GetCatImageByID(thecatapi.WithCatImageID("oR3LMBqEZ"))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(idLookup.ID)
	fmt.Println(idLookup.URL)
	fmt.Println(idLookup.Width)
	fmt.Println(idLookup.Height)

	breeds, err := client.GetBreeds(thecatapi.WithBreedLimit(5))
	if err != nil {
		log.Fatal(err)
	}

	for _, breed := range *breeds {
		fmt.Println(breed.Name)
	}

	cats, err := client.SearchCats(
		thecatapi.WithImageSearchLimit(5),
		thecatapi.WithImageSearchOrder(thecatapi.OrderRandom),
		thecatapi.WithImageSearchSize(thecatapi.SizeSmall),
		thecatapi.WithImageSearchFormat(thecatapi.FormatJSON),
		thecatapi.WithImageSearchHasBreeds(true),
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, cat := range *cats {
		fmt.Println(cat.URL)
		fmt.Println(cat.ID)
	}

	filePath := "testdata/cat.jpg"

	image, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	upload, err := client.UploadImage(image, "cosmo.jpg")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(upload.ID)
}
