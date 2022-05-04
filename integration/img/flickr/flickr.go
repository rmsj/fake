//Package flickr implements integration with loremflickr.com to get random images.
package flickr

import (
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/url"
	"strconv"

	"github.com/rmsj/fake/integration/img"
)

type Flickr struct{}

const (
	searchAPI = "https://loremflickr.com"
)

//New returns a value of type Flickr
func New() Flickr {
	return Flickr{}
}

//GetImages searches images on Flickr and returns a list of 10 Images, with the desired criteria and quantity
func (f Flickr) GetImages(search img.ImageSearchSettings) ([]img.Image, error) {
	var images []img.Image

	if search.Quantity <= 0 {
		search.Quantity = 1
	}

	if search.Quantity > 10 {
		search.Quantity = 10
	}

	if search.Width == 0 {
		search.Width = 640
	}

	if search.Height == 0 {
		search.Height = 480
	}

	imgUrl := fmt.Sprintf(searchAPI+"/%s/%s", strconv.Itoa(search.Width), strconv.Itoa(search.Height))
	if len(search.Word) > 0 {
		imgUrl = fmt.Sprintf(imgUrl+"/%s", search.Word)
	}
	u, err := url.ParseRequestURI(imgUrl)
	if err != nil {
		return images, err
	}

	for i := 0; i < search.Quantity; i++ {

		theImage, err := img.ImageFromURL(u.String())
		if err != nil {
			return images, err
		}

		images = append(images, img.Image{
			Width:  search.Width,
			Height: search.Height,
			Url:    u.String(),
			Alt:    "",
			Data:   theImage,
		})
	}

	return images, nil
}
