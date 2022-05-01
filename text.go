package fake

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

// TextProvider must be implemented by types that wants to provide data source for texts
type TextProvider interface {
	Source() string
}

// RealText generates a text string by the Markov chain algorithm.
//
// Depending on the maxWordCount, returns a random valid looking text. The algorithm
// generates a weighted table with the specified number of words as the index and the
// possible following words as the value.// @example 'Alice, swallowing down her flamingo, and began by taking the little golden key'
// maxWordCount is the maximum number of characters the text should contain (minimum: 10)
// prefixLen determines how many words are considered for the generation of the next word.
// The minimum is 1, and it produces a higher level of randomness, although the
// generated text usually doesn't make sense. Higher index sizes (up to 5)
// produce more correct text, at the price of less randomness.
// @return string
func (f Fake) RealText(maxWordCount int, prefixLen int) (string, error) {
	if maxWordCount < 10 {
		return "", errors.New("maxWordCount must be at least 10")
	}

	if prefixLen < 1 {
		return "", errors.New("prefixLength must be at least 1")
	}

	if prefixLen > 5 {
		return "", errors.New("prefixLength must be at most 5")
	}

	err := f.generateChain(prefixLen)
	if err != nil {
		return "", err
	}

	var result []string
	words := f.textChain[prefixLen]
	next := f.randomFromSlice(f.textPrefixes[prefixLen])

	for {

		wordsOnIndex, ok := words[next]
		if !ok {
			break
		}

		// fetch a random word to append
		word := f.randomFromSlice(wordsOnIndex)

		// calculate next index
		currentPrefix := strings.Fields(next)
		currentPrefix = append(currentPrefix, word)
		// shifts prefix in one element, to remain with prefixLength
		currentPrefix = currentPrefix[1:]
		next = strings.Join(currentPrefix, " ")

		// ensure text starts with an uppercase letter
		if len(result) == 0 && !f.validTextStart(word) {
			continue
		}

		// append the element
		result = append(result, word)

		if len(result) > maxWordCount {
			break
		}
	}

	// remove the extra word
	if len(result) > maxWordCount {
		result = result[:len(result)-1]
	}

	// clear the text off some stuff
	chars := []string{"]", "^", "\\\\", "[", ".", ":", ";", "(", ")", "'", "-"}
	reg := strings.Join(chars, "")

	// last treatments - removing special characters, etc.
	text := strings.Join(result, " ")
	re := regexp.MustCompile("[" + reg + "]+")
	text = re.ReplaceAllString(text, "") + "."

	return text, nil
}

//generateChain reads text from the provided Reader and
// parses it into prefixes and suffixes that are stored in t.chain.
func (f Fake) generateChain(prefixLen int) error {

	// do we have a chain cache for this prefixLength?
	if _, ok := f.textChain[prefixLen]; !ok {

		// chain contains is a map of prefixes to a list of suffixes.
		// A prefix is a string of prefixLen words joined with spaces.
		// A suffix is a single word. A prefix can have multiple suffixes.
		chain := make(map[string][]string)

		re := regexp.MustCompile("/\\s+/u")
		text := re.ReplaceAllString(f.text.Source(), " ")

		// prefix is a Markov chain prefix of one or more words.
		var prefix []string
		var prefixes []string
		words := strings.Fields(text)

		if prefixLen > len(words) {
			return errors.New("source text not long enough")
		}

		// initial prefix generation
		for i := 0; i < prefixLen; i++ {
			prefix = append(prefix, words[i])
		}
		// remove the initial prefix from list of words
		words = words[prefixLen:]

		for i := 0; i < len(words); i++ {
			prefixIndex := strings.Join(prefix, " ")
			chainWord := words[i]

			// if we do not have anything for the existing index, add it
			if _, ok := chain[prefixIndex]; !ok {
				chain[prefixIndex] = []string{chainWord}
				prefixes = append(prefixes, prefixIndex)
			}

			prefix = append(prefix, chainWord)
			// shifts prefix in one element, to remain with prefixLength
			prefix = prefix[1:]
		}

		// cache to help with random text generation later on
		f.textPrefixes[prefixLen] = prefixes
		f.textChain[prefixLen] = chain
	}

	return nil
}

// validTextStart checks if the first character of the word is upper case
func (f Fake) validTextStart(word string) bool {
	runes := []rune(word)
	if !unicode.IsUpper(runes[0]) || !unicode.IsLetter(runes[0]) {
		return false
	}

	return true
}
