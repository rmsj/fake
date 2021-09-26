package faker

// Faker is the main type for faking data
type Faker struct {
	person
	internet
}

// New constructs an instance of Faker and returns it
func New(p PersonProvider, i InternetProvider) (Faker, error) {
	person := newPerson(p)
	internet := NewInternet(i, person)
	f := Faker{
		person,
		internet,
	}

	return f, nil
}
