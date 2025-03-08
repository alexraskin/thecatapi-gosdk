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

func WithSubID(subID string) CatImageUploadOptions {
	return func(body *CatImageUploadBody) {
		body.SubID = &subID
	}
}

func WithBreedIDs(breedIDs string) CatImageUploadOptions {
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

func (c *Client) UploadImage(imageData []byte, fileName string, opts ...CatImageUploadOptions) (*CatImageUploadResponse, error) {
	body := defaultCatImageUploadBody()
	body.File = imageData

	for _, opt := range opts {
		opt(&body)
	}

	requestBody, contentType, err := encodeCatImageUploadBody(body, fileName)
	if err != nil {
		return nil, fmt.Errorf("error encoding request body: %v", err)
	}

	var response CatImageUploadResponse

	requestOpts := defaultRequestOptions(c)
	requestOpts.Method = "POST"
	requestOpts.Path = "/images/upload"
	requestOpts.Body = requestBody
	requestOpts.ContentType = contentType
	requestOpts.Result = &response

	err = httpclient.DoRequest(requestOpts)
	if err != nil {
		return nil, fmt.Errorf("error uploading image: %v", err)
	}

	return &response, nil
}
