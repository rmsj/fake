package fake

import (
	"errors"
	"strings"
)

// DNAProvider must be implemented to provide different sets of data for DNA sequence generation
type DNAProvider interface {
	Set() []string
}

// DNASequence generates a simplified random DNA sequence of the given len
func (f Fake) DNASequence(len int) (string, error) {

	if len < 100 {
		return "", errors.New("DNA sequence must be at least 100")
	}

	var dna []string
	charSet := f.dna.Set()
	for i := 0; i < len; i++ {
		dna = append(dna, f.randomFromSlice(charSet))
	}

	return strings.Join(dna, ""), nil
}
