# The Cat API Go SDK 🐱

Welcome to the [The Cat API](https://thecatapi.com/) Go SDK! This purr-fectly crafted library allows you to interact with The Cat API.

### Getting Started

```go
client := thecatapi.NewClient(thecatapi.WithAPIKey("YOUR-API-KEY"))
```

### Images

Search for some cats photos

```go
cats, err := client.SearchCats(
    thecatapi.WithSize(thecatapi.SizeSmall),
    thecatapi.WithLimit(10),
)
if err != nil {
    log.Fatalf("Error fetching cats: %v", err)
}

for _, cat := range cats {
    fmt.Printf("Cat ID: %s, URL: %s\n", cat.ID, cat.URL)
}
```

### Breeds

search for 

```go
breeds, err := client.GetBreeds(thecatapi.WithBreedLimit(5))
if err != nil {
    log.Fatalf("Error fetching breeds: %v", err)
}

for _, breed := range breeds {
    fmt.Printf("Breed: %s, Origin: %s\n", breed.Name, breed.Origin)
}
```
### Upload

```go
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
```
