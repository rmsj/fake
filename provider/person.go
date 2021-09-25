package provider

import (
	"fmt"
	"math/rand"

	"github.com/rmsj/faker/random"
)

// PersonProvider must be implemented by types that wants to provide data source
// for names, etc.
type PersonProvider interface {
	titlesMale() []string
	titlesFemale() []string
	firstNames() []string
	firstNamesMale() []string
	firstNamesFemale() []string
	lastNames() []string
	genders() []string
	suffixes() []string
}

// Person provides fake name information
type Person struct {
	provider PersonProvider
}

//Title returns a person title
func (p Person) Title() string {
	return random.FromSliceOfString(append(p.provider.titlesMale(), p.provider.titlesFemale()...))
}

// TitleMale get a title male randomly
func (p Person) TitleMale() string {
	return random.FromSliceOfString(p.provider.titlesMale())
}

// TitleFemale get a title female randomly
func (p Person) TitleFemale() string {
	return random.FromSliceOfString(p.provider.titlesFemale())
}

// FirstName gets a first name randomly
func (p Person) FirstName() string {
	return random.FromSliceOfString(p.provider.firstNames())
}

// FirstNameMale gets a first name male randomly
func (p Person) FirstNameMale() string {
	return random.FromSliceOfString(p.provider.firstNamesMale())
}

// FirstNameFemale returns a random female first name
func (p Person) FirstNameFemale() string {
	return random.FromSliceOfString(p.provider.firstNamesFemale())
}

// LastName get fake lastname
func (p Person) LastName() string {
	return random.FromSliceOfString(p.provider.lastNames())
}

// Name returns a random full name with title
func (p Person) Name() string {
	if n := rand.Intn(100); n%2 == 0 {
		return fmt.Sprintf("%s %s %s", random.FromSliceOfString(p.provider.titlesFemale()),
			random.FromSliceOfString(p.provider.firstNamesFemale()),
			random.FromSliceOfString(p.provider.lastNames()))
	}

	return fmt.Sprintf("%s %s %s", random.FromSliceOfString(p.provider.titlesMale()),
		random.FromSliceOfString(p.provider.firstNamesMale()),
		random.FromSliceOfString(p.provider.lastNames()))
}

// Gender returns a random gender
func (p Person) Gender() string {
	return random.FromSliceOfString(p.provider.genders())
}

//Suffix returns a person name suffix
func (p Person) Suffix() string {
	return random.FromSliceOfString(p.provider.suffixes())
}
