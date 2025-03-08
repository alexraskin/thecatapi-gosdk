# The Cat API Go SDK 🐱

Welcome to the [The Cat API](https://thecatapi.com/) Go SDK! This purr-fectly crafted library allows you to interact with The Cat API.

### Getting Started
Create a client:

```go
// Create a client with your API key
client := thecatapi.NewClient(thecatapi.WithAPIKey("YOUR-API-KEY"))
```

### Fetching Cat Images

SearchCats:

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

### Exploring Breeds

Breeds:

```go
breeds, err := client.GetBreeds(thecatapi.WithBreedLimit(5))
if err != nil {
    log.Fatalf("Error fetching breeds: %v", err)
}

for _, breed := range breeds {
    fmt.Printf("Breed: %s, Origin: %s\n", breed.Name, breed.Origin)
}
```
### Uploading Your Cat Images

Upload your cat photos

```go
err := client.UploadImage("path/to/your/cat.jpg", thecatapi.WithSubID("my-cat"), thecatapi.WithBreedIDs("beng"))
if err != nil {
    log.Fatalf("Error uploading image: %v", err)
}
```
