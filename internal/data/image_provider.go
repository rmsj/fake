package data

import (
	"github.com/rmsj/fake/integration/img"
	"github.com/rmsj/fake/integration/img/flickr"
)

// ImageProvider provides data to generate random images
type ImageProvider struct{}

// NewImageProvider builds an ImageProvider and returns it
func NewImageProvider() ImageProvider {
	return ImageProvider{}
}

// ImgSource provides the URL for the images to be generated from
func (ip ImageProvider) ImgSource() img.Imager {
	return flickr.New()
}

// Categories list of valid categories to use for LoremPixel
func (ip ImageProvider) Categories() []string {
	return []string{
		"abstract",
		"animals",
		"business",
		"cats",
		"city",
		"sports",
		"fashion",
		"nature",
		"food",
		"nightlife",
		"people",
		"technics",
		"transport",
	}
}
