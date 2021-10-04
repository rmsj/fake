package fake

import (
	"github.com/rmsj/fake/internal/data"
)

// Builder is a function type that returns anything
type Builder func() interface{}

// Fake is the main type for faking data
type Fake struct {
	person       PersonProvider
	internet     InternetProvider
	text         TextProvider
	lorem        LoremProvider
	dna          DNAProvider
	company      CompanyProvider
	textChain    map[int]map[string][]string // chain caches the consecutive words on the chain by prefix length
	textPrefixes map[int][]string            // prefixes caches all the prefixes prefix length
}

// New constructs an instance of Faker and returns it
func New() (Fake, error) {
	f := Fake{
		person:       data.NewPersonProvider(),
		internet:     data.NewInternetProvider(),
		text:         data.NewTextProvider(),
		lorem:        data.NewLoremProvider(),
		dna:          data.NewDNAProvider(),
		company:      data.NewCompanyProvider(),
		textChain:    make(map[int]map[string][]string),
		textPrefixes: make(map[int][]string),
	}

	return f, nil
}

// Factory builds N number of
func (f Fake) Factory(builder Builder, n int) []interface{} {
	var b []interface{}

	for i := 0; i < n; i++ {
		b = append(b, builder())
	}

	return b
}

// ChangePersonProvider changes the data provider for person related fake data generation
func (f Fake) ChangePersonProvider(p PersonProvider) {
	f.person = p
}

// ChangeInternetProvider changes the data provider for internet related fake data generation
func (f Fake) ChangeInternetProvider(i InternetProvider) {
	f.internet = i
}

// ChangeTextProvider changes the data provider for text related fake data generation
func (f Fake) ChangeTextProvider(t TextProvider) {
	f.text = t
}

// ChangeLoremProvider changes the data provider for lorem (dummy text) related fake data generation
func (f Fake) ChangeLoremProvider(l LoremProvider) {
	f.lorem = l
}

// ChangeDNAProvider changes the data provider for DNA sequence related fake data generation
func (f Fake) ChangeDNAProvider(d DNAProvider) {
	f.dna = d
}
