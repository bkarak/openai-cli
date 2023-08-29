package openai

import (
	"errors"
	"slices"
)

type ImageResponse struct {
	Created int64   `json:"created"`
	Data    []Image `json:"data"`
}

type Image struct {
	Url string `json:"url"`
}

func ImageCreate(prompt string, cardinality int, size string, responseFormat string, user ...string) ([]Image, error) {
	ValidImageSizes := []string{"256x256", "512x512", "1024x1024"}
	ValidResponseFormats := []string{"url", "b64_json"}

	if !slices.Contains(ValidImageSizes, size) {
		return make([]Image, 0), errors.New("invalid image size")
	}

	if !slices.Contains(ValidResponseFormats, responseFormat) {
		return make([]Image, 0), errors.New("invalid response format")
	}

	if cardinality < 1 || cardinality > 10 {
		return make([]Image, 0), errors.New("cardinality should be between 1 to 10")
	}

	return make([]Image, 0), nil
}
