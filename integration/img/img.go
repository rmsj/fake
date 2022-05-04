package img

import (
	"image"
	"net/http"
	"net/url"
)

type Image struct {
	Width  int         `json:"width"`
	Height int         `json:"height"`
	Url    string      `json:"url"`
	Alt    string      `json:"alt"`
	Data   image.Image `json:"data"`
}

// ImageSearchSettings provide settings for random image search
type ImageSearchSettings struct {
	Width    int
	Height   int
	Word     string
	Quantity int
}

//Imager defines the behaviour to be implemented by the image providers
type Imager interface {
	GetImages(search ImageSearchSettings) ([]Image, error)
}

//ImageFromURL gets a valid image url and converts to image.Image
func ImageFromURL(imgURL string) (image.Image, error) {
	u, err := url.ParseRequestURI(imgURL)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}
	return img, nil
}
