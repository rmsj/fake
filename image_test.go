package fake_test

import (
	"image"
	"net/http"
	"net/url"
	"testing"

	"github.com/rmsj/fake"
	"github.com/rmsj/fake/integration/img"
	"github.com/rmsj/fake/internal/data"
	"github.com/rmsj/fake/tests"
)

var imgFake fake.Fake
var imgProvider data.ImageProvider

func setupImageTest() {
	var err error
	imgFake, err = fake.New()

	if err != nil {
		panic("error creating fake instance")
	}
	imgProvider = data.NewImageProvider()
}

func TestImgUrl(t *testing.T) {

	setupImageTest()

	tt := []struct {
		name string
		img  img.ImageSearchSettings
	}{
		{"invalid category image", img.ImageSearchSettings{0, 0, "invalid", 1}},
		{"100 x 100 image", img.ImageSearchSettings{100, 100, "cats", 1}},
		{"200 x 200 image", img.ImageSearchSettings{200, 200, "people", 1}},
		{"random size image", img.ImageSearchSettings{0, 0, "animals", 1}},
	}

	t.Log(tests.Given("Given the need to generate random image url"))
	{
		for testID, test := range tt {
			// start with test 1
			testID++
			tf := func(t *testing.T) {
				t.Logf("\tTest %d:\tWhen creating a %s.", testID, test.name)
				{
					imgUrl, err := imgFake.ImageUrl(test.img)
					if err != nil && test.img.Word == "invalid" {
						t.Logf(tests.Success("\t", "Test %d:\tShould not create image with invalid category for default image provider."), testID)
						return
					}

					if err != nil {
						t.Fatalf(tests.Failed("\t", "Test %d:\tShould create valid image URL: %v"), testID, err)
					}

					if len(imgUrl) == 0 {
						t.Fatalf(tests.Failed("\t", "Test %d:\tShould create valid image URL but it's empty."), testID)
					}

					_, err = url.ParseRequestURI(imgUrl)
					if err != nil {
						t.Fatalf(tests.Failed("\t", "Test %d:\tShould create valid image URL %s."), testID, imgUrl)
					}
					t.Logf(tests.Success("\t", "Test %d:\tShould create valid image URL."), testID)

					resp, err := http.Get(imgUrl)
					if err != nil {
						t.Fatalf(tests.Failed("\t", "Test %d:\tNot able to get image via URL %s."), testID, imgUrl)
					}
					defer resp.Body.Close()

					_, _, err = image.Decode(resp.Body)
					if err != nil {
						t.Fatalf(tests.Failed("\t", "Test %d:\tShould be able to create valid image: %v."), testID, err)
					}
					t.Logf(tests.Success("\t", "Test %d:\tShould be able to create a valid image."), testID)
				}
			}

			t.Run(test.name, tf)

		}
	}
}

func TestImage(t *testing.T) {

	setupImageTest()

	tt := []struct {
		name string
		img  img.ImageSearchSettings
	}{
		{"invalid category image", img.ImageSearchSettings{0, 0, "invalid", 1}},
		{"100 x 100 image", img.ImageSearchSettings{100, 100, "cats", 1}},
		{"200 x 200 image", img.ImageSearchSettings{200, 200, "people", 1}},
		{"random size image", img.ImageSearchSettings{0, 0, "animals", 1}},
	}
	t.Log(tests.Given("Given the need to generate random images"))
	{
		for testID, test := range tt {
			// start with test 1
			testID++
			tf := func(t *testing.T) {
				t.Logf("\tTest %d:\tWhen creating a %s.", testID, test.name)
				{
					_, err := imgFake.ImageUrl(test.img)
					if err != nil && test.img.Word == "invalid" {
						t.Logf(tests.Failed("\t", "Test %d:\tShould not create image with invalid category for default image provider."), testID)
						return
					}
					t.Logf(tests.Success("\t", "Test %d:\tShould be able to create a valid image."), testID)
				}
			}

			t.Run(test.name, tf)

		}
	}
}
