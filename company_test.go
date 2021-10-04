package fake_test

import (
	"strings"
	"testing"

	"github.com/rmsj/fake"
	"github.com/rmsj/fake/internal/data"
	"github.com/rmsj/fake/tests"
)

var companyFake fake.Fake
var companyProvider data.CompanyProvider

func setupCompanyTest() {
	var err error
	companyFake, err = fake.New()

	if err != nil {
		panic("error creating fake instance")
	}
	companyProvider = data.NewCompanyProvider()
}

func TestCompany(t *testing.T) {
	setupCompanyTest()

	t.Log(tests.Given("Given the need to generate company name"))
	{
		for testID := 1; testID < 6; testID++ {
			t.Logf("\tTest %d:\tWhen creating a company name.", testID)
			{
				company := companyFake.Company()

				if len(company) == 0 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould create valid company name but it's empty."), testID)
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create valid company name."), testID)
			}
		}
	}
}

func TestJobTitle(t *testing.T) {
	setupCompanyTest()

	validJobTitles := companyProvider.JobTitles()
	t.Log(tests.Given("Given the need to generate random job titles"))
	{
		for testID := 1; testID < 6; testID++ {
			t.Logf("\tTest %d:\tWhen creating a job title.", testID)
			{
				jt := companyFake.JobTitle()

				if len(jt) == 0 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould create valid job title but it's empty."), testID)
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create valid job title."), testID)

				if tests.InArray(jt, validJobTitles) == -1 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a valid job title within proper range: \"%s\"."), testID, jt)
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create a valid job title: \"%s\"."), testID, jt)
			}
		}
	}
}

func TestCatchPhrase(t *testing.T) {
	setupCompanyTest()

	t.Log(tests.Given("Given the need to generate random catch phrases"))
	{
		for testID := 1; testID < 6; testID++ {
			t.Logf("\tTest %d:\tWhen creating a company catch phrase.", testID)
			{
				catch := companyFake.CatchPhrase()

				if len(catch) == 0 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould create valid company catch phrase but it's empty."), testID)
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create valid company catch phrase."), testID)

				expectedLen := len(companyProvider.CatchWords())
				if len(strings.Fields(catch)) != expectedLen {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould create valid company catch phrase with %d words: %s."), testID, expectedLen, catch)
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create valid company catch phrase with %d words: %s."), testID, expectedLen, catch)
			}
		}
	}
}

func TestBuzzPhrase(t *testing.T) {
	setupCompanyTest()

	t.Log(tests.Given("Given the need to generate random buzz phrases"))
	{
		for testID := 1; testID < 6; testID++ {
			t.Logf("\tTest %d:\tWhen creating a company buzz phrase.", testID)
			{
				buzz := companyFake.BuzzPhrase()

				if len(buzz) == 0 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould create valid company buzz phrase but it's empty."), testID)
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create valid company buzz phrase."), testID)

				expectedLen := len(companyProvider.BuzzWords())
				if len(strings.Fields(buzz)) != expectedLen {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould create valid company buzz phrase with %d words: %s."), testID, expectedLen, buzz)
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create valid company buzz phrase with %d words: %s."), testID, expectedLen, buzz)
			}
		}
	}
}

func TestEIN(t *testing.T) {
	setupCompanyTest()

	t.Log(tests.Given("Given the need to generate random buzz phrases"))
	{
		for testID := 1; testID < 6; testID++ {
			t.Logf("\tTest %d:\tWhen creating a company buzz phrase.", testID)
			{
				buzz := companyFake.BuzzPhrase()

				if len(buzz) == 0 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould create valid company buzz phrase but it's empty."), testID)
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create valid company buzz phrase."), testID)

				expectedLen := len(companyProvider.BuzzWords())
				if len(strings.Fields(buzz)) != expectedLen {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould create valid company buzz phrase with %d words: %s."), testID, expectedLen, buzz)
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create valid company buzz phrase with %d words: %s."), testID, expectedLen, buzz)
			}
		}
	}
}
