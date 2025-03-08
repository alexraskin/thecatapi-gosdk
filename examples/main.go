package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alexraskin/thecatapi"
)

func main() {
	client := thecatapi.NewClient(thecatapi.WithAPIKey("YOUR API KEY"))

	cats, err := client.SearchCats(thecatapi.WithImageSearchLimit(10))
	if err != nil {
		log.Fatal(err)
	}

	for _, cat := range cats {
		fmt.Println(cat.URL)
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
