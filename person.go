package fake

import (
	"fmt"
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
	return f.randomFromSlice(append(f.person.TitlesMale(), f.person.TitlesFemale()...))
}

// TitleMale get a title male randomly
func (f Fake) TitleMale() string {
	return f.randomFromSlice(f.person.TitlesMale())
}

// TitleFemale get a title female randomly
func (f Fake) TitleFemale() string {
	return f.randomFromSlice(f.person.TitlesFemale())
}

// FirstName gets a first name randomly
func (f Fake) FirstName() string {
	return f.randomFromSlice(f.person.FirstNames())
}

// FirstNameMale gets a first name male randomly
func (f Fake) FirstNameMale() string {
	return f.randomFromSlice(f.person.FirstNamesMale())
}

// FirstNameFemale returns a random female first name
func (f Fake) FirstNameFemale() string {
	return f.randomFromSlice(f.person.FirstNamesFemale())
}

// LastName get fake lastname
func (f Fake) LastName() string {
	return f.randomFromSlice(f.person.LastNames())
}

// Name returns a random full name with title
func (f Fake) Name() string {
	if n := f.randomInt(100); n%2 == 0 {
		return fmt.Sprintf("%s %s %s", f.randomFromSlice(f.person.TitlesFemale()),
			f.randomFromSlice(f.person.FirstNamesFemale()),
			f.randomFromSlice(f.person.LastNames()))
	}

	return fmt.Sprintf("%s %s %s", f.randomFromSlice(f.person.TitlesMale()),
		f.randomFromSlice(f.person.FirstNamesMale()),
		f.randomFromSlice(f.person.LastNames()))
}

// Gender returns a random gender
func (f Fake) Gender() string {
	return f.randomFromSlice(f.person.Genders())
}

//Suffix returns a Fake name suffix
func (f Fake) Suffix() string {
	return f.randomFromSlice(f.person.Suffixes())
}
