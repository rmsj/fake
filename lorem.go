package fake

import (
	"math/rand"
	"strings"

	"github.com/rmsj/fake/internal/random"
)

// LoremProvider must be implemented by types that wants to provide data source
// for emails, URLs, etc.
type LoremProvider interface {
	Words() []string
}

//Word returns a random word
func (f Fake) Word() string {
	return random.StringFromSlice(f.lorem.Words())
}

//Sentence returns a number of random words as a sentence
func (f Fake) Sentence(wordCount int) string {

	if wordCount <= 0 {
		return ""
	}

	var result []string
	for i := 0; i < wordCount; i++ {
		result = append(result, f.Word())
	}

	return strings.Join(result, " ")
}

//Paragraph returns a random paragraph
func (f Fake) Paragraph(sentenceCount int) string {
	sentenceSize := rand.Intn(10) + 1
	var result []string
	for i := 0; i < sentenceCount; i++ {
		result = append(result, f.Sentence(sentenceSize))
	}
	return strings.Join(result, " ")
}
