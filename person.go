package fake

import (
	"fmt"
	"math/rand"

	"github.com/rmsj/fake/internal/random"
)

// PersonProvider must be implemented by types that wants to provide data source
// for names, etc.
type PersonProvider interface {
	TitlesMale() []string
	TitlesFemale() []string
	FirstNames() []string
	FirstNamesMale() []string
	FirstNamesFemale() []string
	LastNames() []string
	Genders() []string
	Suffixes() []string
}

//Title returns a Fake title
func (f Fake) Title() string {
	return random.StringFromSlice(append(f.person.TitlesMale(), f.person.TitlesFemale()...))
}

// TitleMale get a title male randomly
func (f Fake) TitleMale() string {
	return random.StringFromSlice(f.person.TitlesMale())
}

// TitleFemale get a title female randomly
func (f Fake) TitleFemale() string {
	return random.StringFromSlice(f.person.TitlesFemale())
}

// FirstName gets a first name randomly
func (f Fake) FirstName() string {
	return random.StringFromSlice(f.person.FirstNames())
}

// FirstNameMale gets a first name male randomly
func (f Fake) FirstNameMale() string {
	return random.StringFromSlice(f.person.FirstNamesMale())
}

// FirstNameFemale returns a random female first name
func (f Fake) FirstNameFemale() string {
	return random.StringFromSlice(f.person.FirstNamesFemale())
}

// LastName get fake lastname
func (f Fake) LastName() string {
	return random.StringFromSlice(f.person.LastNames())
}

// Name returns a random full name with title
func (f Fake) Name() string {
	if n := rand.Intn(100); n%2 == 0 {
		return fmt.Sprintf("%s %s %s", random.StringFromSlice(f.person.TitlesFemale()),
			random.StringFromSlice(f.person.FirstNamesFemale()),
			random.StringFromSlice(f.person.LastNames()))
	}

	return fmt.Sprintf("%s %s %s", random.StringFromSlice(f.person.TitlesMale()),
		random.StringFromSlice(f.person.FirstNamesMale()),
		random.StringFromSlice(f.person.LastNames()))
}

// Gender returns a random gender
func (f Fake) Gender() string {
	return random.StringFromSlice(f.person.Genders())
}

//Suffix returns a Fake name suffix
func (f Fake) Suffix() string {
	return random.StringFromSlice(f.person.Suffixes())
}
