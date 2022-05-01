package fake

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rmsj/fake/internal/data"
)

// Builder is a function type that returns anything
type Builder func() any

// Fake is the main type for faking data
type Fake struct {
	person        PersonProvider
	internet      InternetProvider
	text          TextProvider
	lorem         LoremProvider
	dna           DNAProvider
	company       CompanyProvider
	image         ImageProvider
	textChain     map[int]map[string][]string // chain caches the consecutive words on the chain by prefix length
	textPrefixes  map[int][]string            // prefixes caches all the prefixes prefix length
	deterministic bool
	randSeed      int64
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
		image:        data.NewImageProvider(),
		textChain:    make(map[int]map[string][]string),
		textPrefixes: make(map[int][]string),
	}

	return f, nil
}

// Deterministic will make the generated values constant until a call to Nondeterministic
func (f *Fake) Deterministic(seed int64) {
	f.randSeed = seed
	f.deterministic = true
	fmt.Println("DETERMINISTIC:", f.deterministic)
}

// Nondeterministic reset the seed to it's default randomness for each generated value
func (f *Fake) Nondeterministic() {
	f.deterministic = false
}

// Factory builds N number of
func (f Fake) Factory(builder Builder, n int) []any {
	var b []any

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

var rd *rand.Rand

// randomFromSlice returns a random element from the given slice of string
func (f Fake) randomFromSlice(s []string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s[0]
	}

	return s[f.randomInt(len(s)-1)]
}

// randomFromSlice returns a random element from the given slice of string
func (f Fake) randomInt(interval int) int {

	//default, random generation
	if !f.deterministic {
		rand.Seed(time.Now().UnixNano())
		return rand.Intn(interval)
	}

	// The default number generator is deterministic, so it’ll produce the same SEQUENCE of numbers each time by default.
	// so by seeding it again, given the same semi-open interval, it will produce the same number
	// deterministic - same seed
	rand.Seed(f.randSeed)
	return rand.Intn(interval)

}
