package faker

import (
	"math/rand"
	"strings"

	"github.com/rmsj/faker/random"
)

// LoremProvider must be implemented by types that wants to provide data source
// for emails, URLs, etc.
type LoremProvider interface {
	Words() []string
}

// lorem provides random words and sentences - lorem ipsum.
type lorem struct {
	Provider LoremProvider
}

// NewLorem constructs a lotem  faker type value and returns it
func NewLorem(provider InternetProvider, person person) lorem {
	return lorem{
		Provider: provider,
	}
}

//Word returns a random word
func (l lorem) Word() string {
	return random.StringFromSlice(l.Provider.Words())
}

//Sentence returns a number of random words as a sentence
func (l lorem) Sentence(wordCount int) string {

	if wordCount <= 0 {
		return ""
	}

	var result []string
	for i := 0; i < wordCount; i++ {
		result = append(result, l.Word())
	}

	return strings.Join(result, " ")
}

//Paragraph returns a random paragraph
func (l lorem) Paragraph(sentenceCount int) string {
	sentenceSize := rand.Intn(10) + 1
	var result []string
	for i := 0; i < sentenceCount; i++ {
		result = append(result, l.Sentence(sentenceSize))
	}
	return strings.Join(result, " ")
}
