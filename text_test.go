package faker_test

import (
	"strings"
	"testing"

	"github.com/rmsj/faker"
	"github.com/rmsj/faker/provider"
)

func TestRealText(t *testing.T) {

	p := provider.NewEnglishTextProvider()
	tFaker := faker.NewText(p)

	text, err := tFaker.RealText(100, 4)
	if err != nil {
		t.Fatal("error generating text", err)
	}

	t.Log(text)
	t.Log(len(strings.Fields(text)))
}
