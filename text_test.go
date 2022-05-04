package fake_test

import (
	"strings"
	"testing"

	"github.com/rmsj/fake"
	"github.com/rmsj/fake/tests"
)

func TestRealText(t *testing.T) {

	f, err := fake.New()

	if err != nil {
		t.Fatal("error creating fake instance")
	}

	tt := []struct {
		name      string
		wordCount int
		prefixLen int
	}{
		{"Short Text", 100, 2},
		{"Medium text", 5000, 2},
		{"Long text", 20000, 2},
	}

	t.Log(tests.Given("Given the need to create generate fake 'real text'"))
	{
		for testID, test := range tt {
			// start with test 1
			testID++
			tf := func(t *testing.T) {
				t.Logf("\tTest %d:\tWhen creating a %s of %d words.", testID, test.name, test.wordCount)
				{
					text, err := f.RealText(test.wordCount, test.prefixLen)
					if err != nil {
						t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created text but errored: %v"), testID, err)
					}

					if len(text) == 0 {
						t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a valid text but it's empty: %s"), testID, text)
					}
					t.Logf(tests.Success("\t", "Test %d:\tShould have create a valid text"), testID)

					wordCount := len(strings.Fields(text))
					if wordCount != test.wordCount {
						t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a text with %d words: %d created."), testID, test.wordCount, wordCount)
					}
					t.Logf(tests.Success("\t", "Test %d:\tShould a text with %d words."), testID, test.wordCount)
				}
			}

			t.Run(test.name, tf)

		}
	}
}
