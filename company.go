package fake

import (
	"strings"
)

// CompanyProvider must be implemented to provide all data for company related fake data generation
type CompanyProvider interface {
	NameFormats() []string
	Suffixes() []string
	CatchWords() [][]string
	BuzzWords() [][]string
	JobTitles() []string
	EIN(rand1, rand2 int) string
}

// Company generates a fake random company name and returns it
func (f Fake) Company() string {

	name := f.randomFromSlice(f.company.NameFormats())

	for strings.Contains(name, "lastName") {
		name = strings.Replace(name, "{{lastName}}", f.LastName(), 1)
	}
	if strings.Contains(name, "companySuffix") {
		name = strings.ReplaceAll(name, "{{companySuffix}}", f.randomFromSlice(f.company.Suffixes()))
	}

	return name
}

// JobTitle returns a random job title from data source
func (f Fake) JobTitle() string {
	return f.randomFromSlice(f.company.JobTitles())
}

// CatchPhrase builds a random catch phrase from source
func (f Fake) CatchPhrase() string {
	var catchPhrase []string
	for _, words := range f.company.CatchWords() {
		catchPhrase = append(catchPhrase, f.randomFromSlice(words))
	}

	return strings.Join(catchPhrase, " ")
}

// BuzzPhrase generates random phrase with buzz words
func (f Fake) BuzzPhrase() string {
	var buzzPhrase []string
	for _, words := range f.company.BuzzWords() {
		buzzPhrase = append(buzzPhrase, f.randomFromSlice(words))
	}

	return strings.Join(buzzPhrase, " ")
}

// EIN creates a random Employer Identification Number
func (f Fake) EIN() string {
	return f.company.EIN(f.randomInt(98), f.randomInt(9999998))
}
