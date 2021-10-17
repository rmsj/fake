package fake

import (
	"errors"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"

	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// ImageSettings provide settings for random image
type ImageSettings struct {
	Width    int
	Height   int
	Category string
}

// ImageProvider must be implemented by types that wants to provide data source
// for images.
type ImageProvider interface {
	// ImgSource default is lorempixel.com
	ImgSource() string
	// Categories is the list of valid categories for the default image source - LoremPixel
	Categories() []string
}

//ImageUrl returns a random image url
func (f Fake) ImageUrl(img ImageSettings) (string, error) {

	if len(img.Category) > 0 && InArray(img.Category, f.image.Categories()) == -1 {
		return "", errors.New(fmt.Sprintf("category not in the valid list %s", strings.Join(f.image.Categories(), ", ")))
	}

	width := img.Width
	if width == 0 {
		width = 640
	}

	height := img.Height
	if height == 0 {
		height = 480
	}

	imgUrl := fmt.Sprintf(f.image.ImgSource()+"/%s/%s", strconv.Itoa(width), strconv.Itoa(height))
	if len(img.Category) > 0 {
		imgUrl = fmt.Sprintf(imgUrl+"/%s", img.Category)
	}

	u, err := url.ParseRequestURI(imgUrl)
	if err != nil {
		return "", err
	}

	resp, err := http.Get(u.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	_, _, err = image.Decode(resp.Body)
	if err != nil {
		return "", err
	}

	return u.String(), nil
}

//Image returns a random image.Image
func (f Fake) Image(img ImageSettings) (image.Image, error) {
	if len(img.Category) > 0 && InArray(img.Category, f.image.Categories()) == -1 {
		return nil, errors.New(fmt.Sprintf("category not in the valid list %s", strings.Join(f.image.Categories(), ", ")))
	}
	width := img.Width
	if width == 0 {
		width = 640
	}

	height := img.Height
	if height == 0 {
		height = 480
	}

	imgUrl := fmt.Sprintf(f.image.ImgSource()+"/%s/%s", strconv.Itoa(width), strconv.Itoa(height))
	if len(img.Category) > 0 {
		imgUrl = fmt.Sprintf(imgUrl+"/%s", img.Category)
	}
	u, err := url.ParseRequestURI(imgUrl)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	image, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}
	return image, nil
}

// InArray checks if a given value exists in a specific array and returns the index or -1 if not found
func InArray(val interface{}, array interface{}) int {
	values := reflect.ValueOf(array)

	if reflect.TypeOf(array).Kind() == reflect.Slice || values.Len() > 0 {
		for i := 0; i < values.Len(); i++ {
			if reflect.DeepEqual(val, values.Index(i).Interface()) {
				return i
			}
		}
	}

	return -1
}
