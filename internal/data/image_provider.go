package data

// ImageProvider provides data to generate random images
type ImageProvider struct{}

// NewImageProvider builds an ImageProvider and returns it
func NewImageProvider() ImageProvider {
	return ImageProvider{}
}

// ImgSource provides the URL for the images to be generated from
func (ip ImageProvider) ImgSource() string {
	return "https://loremflickr.com"
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
