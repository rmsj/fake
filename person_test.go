package fake_test

import (
	"strings"
	"testing"

	"github.com/rmsj/fake"
	"github.com/rmsj/fake/internal/data"
	"github.com/rmsj/fake/tests"
)

var personProvider data.PersonProvider
var personFake fake.Fake

func setupPersonTest() {
	var err error
	personFake, err = fake.New()

	if err != nil {
		panic("error creating fake instance")
	}
	personProvider = data.NewPersonProvider()
}

func TestTitle(t *testing.T) {
	setupPersonTest()

	validTitles := append(personProvider.TitlesMale(), personProvider.TitlesFemale()...)
	t.Log(tests.Given("Given the need to test creating person title"))
	{
		for testID := 1; testID < 5; testID++ {
			t.Logf("\tTest %d:\tWhen creating a Title of any gender.", testID)
			{
				title := personFake.Title()
				if len(title) == 0 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a title but empty string."), testID)
				}

				if tests.InArray(title, validTitles) == -1 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a valid title within: \"%s\"."), testID, strings.Join(validTitles, ", "))
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create a valid title: \"%s\"."), testID, title)
			}
		}
	}
}

func TestTitleFemale(t *testing.T) {
	setupPersonTest()

	validTitles := personProvider.TitlesFemale()
	t.Log(tests.Given("Given the need to test creating person title"))
	{
		for testID := 1; testID < 5; testID++ {
			t.Logf("\tTest %d:\tWhen creating a Female Title.", testID)
			{
				title := personFake.TitleFemale()
				if len(title) == 0 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a title but empty string."), testID)
				}

				if tests.InArray(title, validTitles) == -1 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a valid title within: \"%s\"."), testID, strings.Join(validTitles, ", "))
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create a valid title: \"%s\"."), testID, title)
			}
		}
	}
}

func TestTitleMale(t *testing.T) {
	setupPersonTest()

	validTitles := personProvider.TitlesMale()
	t.Log(tests.Given("Given the need to test creating person title"))
	{
		for testID := 1; testID < 5; testID++ {
			t.Logf("\tTest %d:\tWhen creating a Male Title.", testID)
			{
				title := personFake.TitleMale()
				if len(title) == 0 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a title but empty string."), testID)
				}

				if tests.InArray(title, validTitles) == -1 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a valid title within: \"%s\"."), testID, strings.Join(validTitles, ", "))
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create a valid title: \"%s\"."), testID, title)
			}
		}
	}
}

func TestFirstName(t *testing.T) {
	setupPersonTest()

	validNames := personProvider.FirstNames()
	t.Log(tests.Given("Given the need to test creating person first name"))
	{
		for testID := 1; testID < 5; testID++ {
			t.Logf("\tTest %d:\tWhen creating a first name of any gender.", testID)
			{
				firstName := personFake.FirstName()

				if len(firstName) == 0 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a first name but empty string."), testID)
				}

				if tests.InArray(firstName, validNames) == -1 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a valid first name within: \"%s\"."), testID, strings.Join(validNames, ", "))
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create a valid first name: \"%s\"."), testID, firstName)
			}
		}
	}
}

func TestFirstNameFemale(t *testing.T) {
	setupPersonTest()

	validNames := personProvider.FirstNamesFemale()
	t.Log(tests.Given("Given the need to test creating female person first name"))
	{
		for testID := 1; testID < 5; testID++ {
			t.Logf("\tTest %d:\tWhen creating a first name of female gender.", testID)
			{
				firstName := personFake.FirstNameFemale()
				if len(firstName) == 0 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a female first name but empty string."), testID)
				}

				if tests.InArray(firstName, validNames) == -1 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a valid female first name within: \"%s\"."), testID, strings.Join(validNames, ", "))
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create a valid female first name: \"%s\"."), testID, firstName)
			}
		}
	}
}

func TestFirstNameMale(t *testing.T) {
	setupPersonTest()

	validNames := personProvider.FirstNamesMale()
	t.Log(tests.Given("Given the need to test creating male person first name"))
	{
		for testID := 1; testID < 5; testID++ {
			t.Logf("\tTest %d:\tWhen creating a first name of male gender.", testID)
			{
				firstName := personFake.FirstNameMale()
				if len(firstName) == 0 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a male first name but empty string."), testID)
				}

				if tests.InArray(firstName, validNames) == -1 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a valid male first name within: \"%s\"."), testID, strings.Join(validNames, ", "))
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create a valid male first name: \"%s\"."), testID, firstName)
			}
		}
	}
}

func TestLastName(t *testing.T) {
	setupPersonTest()

	validNames := personProvider.LastNames()
	t.Log(tests.Given("Given the need to test creating a person last name"))
	{
		for testID := 1; testID < 5; testID++ {
			t.Logf("\tTest %d:\tWhen creating a last name.", testID)
			{
				lastName := personFake.LastName()
				if len(lastName) == 0 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a male last name but empty string."), testID)
				}

				if tests.InArray(lastName, validNames) == -1 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a valid last name within: \"%s\"."), testID, strings.Join(validNames, ", "))
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create a valid last name: \"%s\"."), testID, lastName)
			}
		}
	}
}

func TestName(t *testing.T) {
	setupPersonTest()

	t.Log(tests.Given("Given the need to test creating a person name"))
	{
		for testID := 1; testID < 5; testID++ {
			t.Logf("\tTest %d:\tWhen creating a person name.", testID)
			{
				name := personFake.Name()
				if len(name) == 0 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a person name but empty string."), testID)
				}

				if len(strings.Fields(name)) < 3 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a valid name with title, first and last name: \"%s\"."), testID, name)
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create a valid name with title, first and last name: \"%s\"."), testID, name)
			}
		}
	}
}

func TestGender(t *testing.T) {
	setupPersonTest()

	validGenders := personProvider.Genders()
	t.Log(tests.Given("Given the need to test creating person gender"))
	{
		for testID := 1; testID < 5; testID++ {
			t.Logf("\tTest %d:\tWhen creating a person gender.", testID)
			{
				gender := personFake.Gender()
				if len(gender) == 0 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a person gender."), testID)
				}

				if tests.InArray(gender, validGenders) == -1 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a valid gender within: \"%s\"."), testID, strings.Join(validGenders, ", "))
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create a valid gender: \"%s\"."), testID, gender)
			}
		}
	}
}

func TestSuffix(t *testing.T) {
	setupPersonTest()

	validSuffixes := personProvider.Suffixes()
	t.Log(tests.Given("Given the need to test creating person name suffix"))
	{
		for testID := 1; testID < 5; testID++ {
			t.Logf("\tTest %d:\tWhen creating a person name suffix.", testID)
			{
				suffix := personFake.Suffix()
				if len(suffix) == 0 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a person name suffix."), testID)
				}

				if tests.InArray(suffix, validSuffixes) == -1 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a valid suffix within: \"%s\"."), testID, strings.Join(validSuffixes, ", "))
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create a valid suffix: \"%s\"."), testID, suffix)
			}
		}
	}
}
