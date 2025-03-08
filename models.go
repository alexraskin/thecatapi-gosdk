package thecatapi

type ImageSize string
type Format string
type OrderType string

const (
	SizeThumb ImageSize = "thumb"
	SizeSmall ImageSize = "small"
	SizeMed   ImageSize = "med"
	SizeFull  ImageSize = "full"

	FormatJSON Format = "json"
	FormatSrc  Format = "src"

	OrderRandom OrderType = "RANDOM"
	OrderAsc    OrderType = "ASC"
	OrderDesc   OrderType = "DESC"
)

type CatImageSearchParams struct {
	Size      ImageSize `json:"size,omitempty"`
	MimeTypes []string  `json:"mime_types,omitempty"`
	Format    Format    `json:"format,omitempty"`
	HasBreeds bool      `json:"has_breeds,omitempty"`
	Order     OrderType `json:"order,omitempty"`
	Page      int       `json:"page,omitempty"`
	Limit     int       `json:"limit,omitempty"`
}

type CatImageSearchResponse struct {
	ID     string `json:"id"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type CatBreedParams struct {
	Page  int `json:"page,omitempty"`
	Limit int `json:"limit,omitempty"`
}

type Weight struct {
	Imperial string `json:"imperial,omitempty"`
	Metric   string `json:"metric,omitempty"`
}

type BreedImage struct {
	ID     string `json:"id"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
}

type CatBreedResponse struct {
	ID               string     `json:"id"`
	Name             string     `json:"name"`
	CfaURL           string     `json:"cfa_url,omitempty"`
	VetstreetURL     string     `json:"vetstreet_url,omitempty"`
	VcahospitalsURL  string     `json:"vcahospitals_url,omitempty"`
	Temperament      string     `json:"temperament,omitempty"`
	Origin           string     `json:"origin,omitempty"`
	CountryCodes     string     `json:"country_codes,omitempty"`
	CountryCode      string     `json:"country_code,omitempty"`
	Description      string     `json:"description,omitempty"`
	LifeSpan         string     `json:"life_span,omitempty"`
	Indoor           int        `json:"indoor,omitempty"`
	Lap              int        `json:"lap,omitempty"`
	AltNames         string     `json:"alt_names,omitempty"`
	Adaptability     int        `json:"adaptability,omitempty"`
	AffectionLevel   int        `json:"affection_level,omitempty"`
	ChildFriendly    int        `json:"child_friendly,omitempty"`
	DogFriendly      int        `json:"dog_friendly,omitempty"`
	EnergyLevel      int        `json:"energy_level,omitempty"`
	Grooming         int        `json:"grooming,omitempty"`
	HealthIssues     int        `json:"health_issues,omitempty"`
	Intelligence     int        `json:"intelligence,omitempty"`
	SheddingLevel    int        `json:"shedding_level,omitempty"`
	SocialNeeds      int        `json:"social_needs,omitempty"`
	StrangerFriendly int        `json:"stranger_friendly,omitempty"`
	Vocalisation     int        `json:"vocalisation,omitempty"`
	Experimental     int        `json:"experimental,omitempty"`
	Hairless         int        `json:"hairless,omitempty"`
	Natural          int        `json:"natural,omitempty"`
	Rare             int        `json:"rare,omitempty"`
	Rex              int        `json:"rex,omitempty"`
	SuppressedTail   int        `json:"suppressed_tail,omitempty"`
	ShortLegs        int        `json:"short_legs,omitempty"`
	WikipediaURL     string     `json:"wikipedia_url,omitempty"`
	Hypoallergenic   int        `json:"hypoallergenic,omitempty"`
	ReferenceImageID string     `json:"reference_image_id,omitempty"`
	Image            BreedImage `json:"image,omitempty"`
	Weight           Weight     `json:"weight,omitempty"`
}

type CatImageUploadBody struct {
	File     []byte
	SubID    *string
	BreedIDs *string
}

type CatImageUploadResponse struct {
	ID               string `json:"id"`
	URL              string `json:"url"`
	Width            int    `json:"width"`
	Height           int    `json:"height"`
	OriginalFilename string `json:"original_filename"`
	Pending          int    `json:"pending"`
	Approved         int    `json:"approved"`
}

type CatByIDImageParams struct {
	ID string `json:"id"`
}

type CatByIDImageResponse struct {
	ID         string   `json:"id"`
	URL        string   `json:"url"`
	Width      int      `json:"width"`
	Height     int      `json:"height"`
	MimeType   string   `json:"mime_type"`
	Categories []string `json:"categories,omitempty"`
	BreedsIDs  string   `json:"breeds_ids,omitempty"`
}

type YourCatImagesQueryParams struct {
	Limit            int    `json:"limit,omitempty"`
	Page             int    `json:"page,omitempty"`
	Order            string `json:"order,omitempty"`
	SubID            string `json:"sub_id,omitempty"`
	BreedIDs         string `json:"breed_ids,omitempty"`
	CategoryIDs      string `json:"category_ids,omitempty"`
	Format           string `json:"format,omitempty"`
	OriginalFilename string `json:"original_filename,omitempty"`
	UserID           string `json:"user_id,omitempty"`
}

type YourCatImagesResponse struct {
	ID               string   `json:"id"`
	URL              string   `json:"url"`
	Width            *int     `json:"width"`
	Height           *int     `json:"height"`
	SubID            string   `json:"sub_id"`
	CreatedAt        string   `json:"created_at"`
	OriginalFilename string   `json:"original_filename"`
	BreedIDs         string   `json:"breed_ids"`
	Breeds           []string `json:"breeds"`
}

type CatFactsParams struct {
	Limit int       `json:"limit"`
	Page  int       `json:"page"`
	Order OrderType `json:"order"`
}

type CatFactsResponse struct {
	ID       string `json:"id"`
	Fact     string `json:"fact"`
	BreedIDs string `json:"breed_ids"`
	Title    string `json:"title"`
}
