package fake

import (
	"errors"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"reflect"
	"strings"

	"github.com/rmsj/fake/integration/img"
)

// ImageSettings provide settings for random image
type ImageSettings struct {
	Width    int
	Height   int
	Category string
	Quantity int
}

type Imager interface {
	GetImages()
}

// ImageProvider must be implemented by types that wants to provide data source
// for images.
type ImageProvider interface {
	// ImgSource default is the implementation for loremflickr.com
	ImgSource() img.Imager
	// Categories is the list of valid categories for the default image source - LoremPixel
	Categories() []string
}

//ImageUrl returns a random image url
func (f Fake) ImageUrl(search img.ImageSearchSettings) (string, error) {

	if len(search.Word) > 0 && InArray(search.Word, f.image.Categories()) == -1 {
		return "", errors.New(fmt.Sprintf("category not in the valid list %s", strings.Join(f.image.Categories(), ", ")))
	}

	imgSource := f.image.ImgSource()
	list, err := imgSource.GetImages(search)
	if err != nil {
		return "", err
	}

	theImage := list[f.randomInt(len(list)-1)]

	return theImage.Url, nil
}

//Image returns a random image.Image
func (f Fake) Image(search img.ImageSearchSettings) (image.Image, error) {
	if len(search.Word) > 0 && InArray(search.Word, f.image.Categories()) == -1 {
		return nil, errors.New(fmt.Sprintf("category not in the valid list %s", strings.Join(f.image.Categories(), ", ")))
	}

	imgSource := f.image.ImgSource()
	list, err := imgSource.GetImages(search)
	if err != nil {
		return nil, err
	}

	theImage := list[f.randomInt(len(list)-1)]

	if err != nil {
		return nil, err
	}
	return theImage.Data, nil
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
