package faker_test

import (
	"github.com/rmsj/faker"
	"github.com/rmsj/faker/provider"
	"testing"
)

func TestNew(t *testing.T) {
	pp := provider.NewEnglishPersonProvider()
	ip := provider.NewEnglishInternetProvider()

	f, err := faker.New(pp, ip)
	if err != nil {
		panic(err)
	}

	t.Log(f.FirstName())
	t.Log(f.Email())
	t.Log(f.CompanyEmail())
	t.Log(f.SafeEmail())
	t.Log(f.FreeEmail())
	t.Log(f.Url())
}
