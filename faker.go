package faker

import "github.com/rmsj/faker/provider"

type Faker struct {
	person provider.Person
}

func New(lang string, country string) (Faker, error) {
	f := Faker{}

	return f, nil
}
