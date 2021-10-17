package fake_test

import (
	"strings"
	"testing"

	"github.com/rmsj/fake"
	"github.com/rmsj/fake/internal/data"
	"github.com/rmsj/fake/tests"
)

var loremFake fake.Fake
var loremProvider data.LoremProvider

func setupLoremTest() {
	var err error
	loremFake, err = fake.New()

	if err != nil {
		panic("error creating fake instance")
	}
	loremProvider = data.NewLoremProvider()
}

func TestWord(t *testing.T) {
	setupLoremTest()

	validWords := loremProvider.Words()
	t.Log(tests.Given("Given the need to generate lorem ipsum word"))
	{
		for testID := 1; testID < 6; testID++ {
			t.Logf("\tTest %d:\tWhen creating a random word.", testID)
			{
				lorem := loremFake.Word()

				if len(lorem) == 0 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould create valid lorem ipsum word but it's empty."), testID)
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create valid lorem ipsum word."), testID)

				if tests.InArray(lorem, validWords) == -1 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a valid lorem ipsum within proper range: \"%s\"."), testID, strings.Join(validWords, ", "))
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create a valid lorem ipsum sentence: \"%s\"."), testID, lorem)
			}
		}
	}
}

func TestSentence(t *testing.T) {
	setupLoremTest()

	tt := []struct {
		name      string
		wordCount int
	}{
		{"Invalid Length", 0},
		{"Short Sentence", 25},
		{"Medium Sentence", 100},
		{"Long (ish) Sentence", 255},
	}
	validWords := loremProvider.Words()
	t.Log(tests.Given("Given the need to generate random Lorem Ipsum text sequence"))
	{
		for testID, test := range tt {
			// start with test 1
			testID++
			tf := func(t *testing.T) {
				t.Logf("\tTest %d:\tWhen creating a %s of %d words.", testID, test.name, test.wordCount)
				{
					lorem := loremFake.Sentence(test.wordCount)

					if len(lorem) == 0 && test.wordCount == 0 {
						t.Logf(tests.Success("\t", "Test %d:\tShould not create valid lorem ipsum sentence with invalid length <= 0."), testID)
						return
					}

					if len(lorem) == 0 {
						t.Fatalf(tests.Failed("\t", "Test %d:\tShould create valid lorem ipsum sentence but it's empty."), testID)
					}
					t.Logf(tests.Success("\t", "Test %d:\tShould create valid lorem ipsum sentence."), testID)

					wordCount := len(strings.Fields(lorem))
					if wordCount != test.wordCount {
						t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a valid lorem ipsum sentence with %d words long: %d"), testID, test.wordCount, wordCount)
					}
					t.Logf(tests.Success("\t", "Test %d:\tShould have created a valid lorem ipsum sentence with %d words long"), testID, test.wordCount)

					for _, word := range strings.Fields(lorem) {
						if tests.InArray(word, validWords) == -1 {
							t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a valid lorem ipsum (%s) within proper range: \"%s\"."), testID, word, strings.Join(validWords, ", "))
						}
					}
					t.Logf(tests.Success("\t", "Test %d:\tShould create a valid lorem ipsum sentence."), testID)
				}
			}

			t.Run(test.name, tf)

		}
	}
}

func TestParagraph(t *testing.T) {
	setupLoremTest()

	tt := []struct {
		name          string
		sentenceCount int
	}{
		{"Invalid Length", 0},
		{"Short Paragraph", 2},
		{"Medium Sentence", 10},
		{"Long (ish) Sentence", 20},
	}
	validWords := loremProvider.Words()
	t.Log(tests.Given("Given the need to generate random Lorem Ipsum text sequence"))
	{
		for testID, test := range tt {
			// start with test 1
			testID++
			tf := func(t *testing.T) {
				t.Logf("\tTest %d:\tWhen creating a %s of %d words.", testID, test.name, test.sentenceCount)
				{
					lorem := loremFake.Paragraph(test.sentenceCount)

					if len(lorem) == 0 && test.sentenceCount == 0 {
						t.Logf(tests.Success("\t", "Test %d:\tShould not create valid lorem ipsum sentence with invalid length <= 0."), testID)
						return
					}

					if len(lorem) == 0 {
						t.Fatalf(tests.Failed("\t", "Test %d:\tShould create valid lorem ipsum sentence but it's empty."), testID)
					}
					t.Logf(tests.Success("\t", "Test %d:\tShould create valid lorem ipsum sentence."), testID)

					for _, word := range strings.Fields(lorem) {
						if tests.InArray(word, validWords) == -1 {
							t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a valid lorem ipsum (%s) within proper range: \"%s\"."), testID, word, strings.Join(validWords, ", "))
						}
					}
					t.Logf(tests.Success("\t", "Test %d:\tShould create a valid lorem ipsum sentence."), testID)
				}
			}

			t.Run(test.name, tf)

		}
	}
}
