package thecatapi

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/textproto"

	"github.com/alexraskin/thecatapi/internal/httpclient"
)

type CatImageUploadOptions func(*CatImageUploadBody)

func defaultCatImageUploadBody() CatImageUploadBody {
	return CatImageUploadBody{
		SubID:    nil,
		BreedIDs: nil,
	}
}

func WithCatImageUploadSubID(subID string) CatImageUploadOptions {
	return func(body *CatImageUploadBody) {
		body.SubID = &subID
	}
}

func WithCatImageUploadBreedIDs(breedIDs string) CatImageUploadOptions {
	return func(body *CatImageUploadBody) {
		body.BreedIDs = &breedIDs
	}
}

func encodeCatImageUploadBody(body CatImageUploadBody, fileName string) (*bytes.Buffer, string, error) {
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, fileName))
	h.Set("Content-Type", "image/jpeg")

	filePart, err := writer.CreatePart(h)
	if err != nil {
		return nil, "", fmt.Errorf("error creating form file: %v", err)
	}

	_, err = filePart.Write(body.File)
	if err != nil {
		return nil, "", fmt.Errorf("error writing file data: %v", err)
	}

	if body.SubID != nil {
		if err := writer.WriteField("sub_id", *body.SubID); err != nil {
			return nil, "", fmt.Errorf("error writing sub_id field: %v", err)
		}
	}
	if body.BreedIDs != nil {
		if err := writer.WriteField("breed_ids", *body.BreedIDs); err != nil {
			return nil, "", fmt.Errorf("error writing breed_ids field: %v", err)
		}
	}

	if err := writer.Close(); err != nil {
		return nil, "", fmt.Errorf("error closing writer: %v", err)
	}

	return &b, writer.FormDataContentType(), nil
}

// UploadImage uploads an image to The Cat API.
// It allows customization of the upload request through functional options.
//
// Parameters:
//
//	imageData - A byte slice containing the image data to be uploaded.
//	fileName - The name of the file being uploaded.
//	opts - A variadic list of CatImageUploadOptions functions that modify the upload parameters.
//	       These options can be used to set additional fields like SubID and BreedIDs.
//
// Returns:
//
//	*CatImageUploadResponse - A pointer to a CatImageUploadResponse struct containing information about the uploaded image.
//	error - An error if the upload fails or if there is an issue with the response.
//
// Example usage:
//
//	image, err := os.ReadFile("cat.jpg")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	upload, err := client.UploadImage(image, "cat.jpg", thecatapi.WithCatImageUploadSubID("my-cat"))
//	if err != nil {
//	    log.Fatalf("Error uploading image: %v", err)
//	}
//	fmt.Printf("Uploaded Image ID: %s\n", upload.ID)
func (c *Client) UploadImage(imageData []byte, fileName string, opts ...CatImageUploadOptions) (*CatImageUploadResponse, error) {
	body := defaultCatImageUploadBody()
	body.File = imageData

	for _, fn := range opts {
		fn(&body)
	}

	requestBody, contentType, err := encodeCatImageUploadBody(body, fileName)
	if err != nil {
		return nil, fmt.Errorf("error encoding request body: %v", err)
	}

	var response CatImageUploadResponse

	requestOpts := newRequestOptions(c, "/images/upload", nil, requestBody, &response)
	requestOpts.Method = "POST"
	requestOpts.ContentType = contentType

	err = httpclient.DoRequest(requestOpts)
	if err != nil {
		return nil, fmt.Errorf("error uploading image: %v", err)
	}

	return &response, nil
}
