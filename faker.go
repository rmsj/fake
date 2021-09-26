package faker

// Builder is a function type that returns anything
type Builder func() interface{}

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

// Factory builds N number of
func (f Faker) Factory(builder Builder, n int) []interface{} {
	var b []interface{}

	for i := 0; i < n; i++ {
		b = append(b, builder())
	}

	return b
}
