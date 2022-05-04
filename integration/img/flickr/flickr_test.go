package flickr_test

import (
	"testing"

	"github.com/rmsj/fake/integration/img"
	"github.com/rmsj/fake/integration/img/flickr"
	"github.com/rmsj/fake/tests"
)

func TestGetImages(t *testing.T) {

	flk := flickr.New()

	t.Log(tests.Given("Given the need to generate image list from Flickr API integration"))
	{
		for testID := 0; testID < 6; testID++ {
			t.Logf("\tTest %d:\tWhen getting %d image(s).", testID, testID+1)
			{
				search := img.ImageSearchSettings{
					Width:    640,
					Height:   480,
					Word:     "cat",
					Quantity: testID + 1,
				}

				images, err := flk.GetImages(search)
				if err != nil {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould be able to get images from Flickr API: %v."), testID, err)
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould be able to get images from Flickr API."), testID)

				if len(images) != search.Quantity {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould have got %d images - got %d."), testID, len(images), search.Quantity)
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould have got %d image(s)."), testID, search.Quantity)

				if len(images) != search.Quantity {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould have got %d images - got %d."), testID, len(images), search.Quantity)
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould have got %d image(s)."), testID, search.Quantity)

				for _, v := range images {
					if v.Data == nil {
						t.Fatalf(tests.Failed("\t", "Test %d:\tShould have valid image data."), testID)
					}
					t.Logf(tests.Success("\t", "Test %d:\tShould have valid image data."), testID)
				}
			}
		}
	}
}
