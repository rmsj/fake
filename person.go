package faker

import (
	"fmt"
	"math/rand"

	"github.com/rmsj/faker/random"
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

// person provides fake name information
type person struct {
	provider PersonProvider
}

// newPerson constructs a person type value and returns it
func newPerson(p PersonProvider) person {
	return person{
		provider: p,
	}
}

//Title returns a person title
func (p person) Title() string {
	return random.StringFromSlice(append(p.provider.TitlesMale(), p.provider.TitlesFemale()...))
}

// TitleMale get a title male randomly
func (p person) TitleMale() string {
	return random.StringFromSlice(p.provider.TitlesMale())
}

// TitleFemale get a title female randomly
func (p person) TitleFemale() string {
	return random.StringFromSlice(p.provider.TitlesFemale())
}

// FirstName gets a first name randomly
func (p person) FirstName() string {
	return random.StringFromSlice(p.provider.FirstNames())
}

// FirstNameMale gets a first name male randomly
func (p person) FirstNameMale() string {
	return random.StringFromSlice(p.provider.FirstNamesMale())
}

// FirstNameFemale returns a random female first name
func (p person) FirstNameFemale() string {
	return random.StringFromSlice(p.provider.FirstNamesFemale())
}

// LastName get fake lastname
func (p person) LastName() string {
	return random.StringFromSlice(p.provider.LastNames())
}

// Name returns a random full name with title
func (p person) Name() string {
	if n := rand.Intn(100); n%2 == 0 {
		return fmt.Sprintf("%s %s %s", random.StringFromSlice(p.provider.TitlesFemale()),
			random.StringFromSlice(p.provider.FirstNamesFemale()),
			random.StringFromSlice(p.provider.LastNames()))
	}

	return fmt.Sprintf("%s %s %s", random.StringFromSlice(p.provider.TitlesMale()),
		random.StringFromSlice(p.provider.FirstNamesMale()),
		random.StringFromSlice(p.provider.LastNames()))
}

// Gender returns a random gender
func (p person) Gender() string {
	return random.StringFromSlice(p.provider.Genders())
}

//Suffix returns a person name suffix
func (p person) Suffix() string {
	return random.StringFromSlice(p.provider.Suffixes())
}
