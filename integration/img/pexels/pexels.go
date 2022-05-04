package pexels

import (
	"encoding/json"
	"errors"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"net/url"

	"github.com/rmsj/fake/integration/img"
)

type Pexels struct {
	apiKey string
}

// Image is the is the response from the Pexels search
type Image struct {
	Width  int         `json:"width"`
	Height int         `json:"height"`
	Url    string      `json:"url"`
	Alt    string      `json:"alt"`
	Data   image.Image `json:"data"`
}

// searchResponse represents the API response from Pexels
type searchResponse struct {
	TotalResults int `json:"total_results"`
	Page         int `json:"page"`
	PerPage      int `json:"per_page"`
	Photos       []struct {
		Id              int    `json:"id"`
		Width           int    `json:"width"`
		Height          int    `json:"height"`
		Url             string `json:"url"`
		Photographer    string `json:"photographer"`
		PhotographerUrl string `json:"photographer_url"`
		PhotographerId  int    `json:"photographer_id"`
		AvgColor        string `json:"avg_color"`
		Src             struct {
			Original  string `json:"original"`
			Large2X   string `json:"large2x"`
			Large     string `json:"large"`
			Medium    string `json:"medium"`
			Small     string `json:"small"`
			Portrait  string `json:"portrait"`
			Landscape string `json:"landscape"`
			Tiny      string `json:"tiny"`
		} `json:"src"`
		Liked bool   `json:"liked"`
		Alt   string `json:"alt"`
	} `json:"photos"`
	NextPage string `json:"next_page"`
}

const (
	searchAPI = "https://api.pexels.com/v1/search"
)

func New(apiKey string) Pexels {
	return Pexels{
		apiKey: apiKey,
	}
}

//GetImages searches images on Pexels API and returns a list of Images, with the desired criteria and quantity
func (p Pexels) GetImages(search img.ImageSearchSettings) ([]Image, error) {
	var images []Image

	if search.Quantity <= 0 {
		search.Quantity = 1
	}

	if search.Quantity > 10 {
		search.Quantity = 10
	}

	reqURL := fmt.Sprintf(searchAPI+"?query=%s&per_page=%d", search.Word, search.Quantity)
	u, err := url.ParseRequestURI(reqURL)
	if err != nil {
		return images, err
	}

	// Create a Bearer string by appending string access token

	// Create a new request using http
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return images, err
	}

	// add authorization header to the req
	req.Header.Add("Authorization", p.apiKey)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return images, err
	}
	defer resp.Body.Close()

	var searchResp searchResponse
	if err := json.NewDecoder(resp.Body).Decode(&searchResp); err != nil {
		return images, err
	}

	// nothing found
	if len(searchResp.Photos) == 0 {
		return images, errors.New("no images wound with current criteria")
	}

	for _, v := range searchResp.Photos {

		theImage, err := img.ImageFromURL(v.Src.Original)
		if err != nil {
			return images, err
		}

		images = append(images, Image{
			Width:  v.Width,
			Height: v.Height,
			Url:    v.Src.Original,
			Alt:    v.Alt,
			Data:   theImage,
		})
	}

	return images, nil
}
