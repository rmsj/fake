package fake_test

import (
	"net"
	"net/url"
	"strings"
	"testing"

	"github.com/rmsj/fake"
	"github.com/rmsj/fake/internal/data"
	"github.com/rmsj/fake/tests"
)

var internetFake fake.Fake
var internetProvider data.InternetProvider

func setupInternetTest() {
	var err error
	internetFake, err = fake.New()

	if err != nil {
		panic("error creating fake instance")
	}
	internetProvider = data.NewInternetProvider()
}

func TestUserName(t *testing.T) {
	setupInternetTest()

	t.Log(tests.Given("Given the need to generate random username"))
	{
		for testID := 1; testID < 6; testID++ {
			t.Logf("\tTest %d:\tWhen creating an username.", testID)
			{
				username := internetFake.Username()

				if len(username) == 0 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould create valid username but it's empty."), testID)
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create valid username."), testID)
			}
		}
	}
}

func TestEmail(t *testing.T) {
	setupInternetTest()

	t.Log(tests.Given("Given the need to generate random email"))
	{
		for testID := 1; testID < 6; testID++ {
			t.Logf("\tTest %d:\tWhen creating an email addresas.", testID)
			{
				email := internetFake.Email()

				if len(email) == 0 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould create valid email but it's empty."), testID)
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create valid email."), testID)

				if !tests.ValidEmail(email) {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould create valid email: %s."), testID, email)
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create valid email."), testID)
			}
		}
	}
}

func TestSafeEmail(t *testing.T) {
	setupInternetTest()

	validDomains := internetProvider.SafeEmailDomains()
	t.Log(tests.Given("Given the need to generate random safe email"))
	{
		for testID := 1; testID < 6; testID++ {
			t.Logf("\tTest %d:\tWhen creating a safe email addresas.", testID)
			{
				email := internetFake.SafeEmail()

				if len(email) == 0 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould create valid email but it's empty."), testID)
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create valid username."), testID)

				if !tests.ValidEmail(email) {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould create valid email: %s."), testID, email)
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create valid email."), testID)

				emailParts := strings.Split(email, "@")
				if tests.InArray(emailParts[1], validDomains) == -1 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a valid email (%s) within proper domain range: \"%s\"."), testID, email, strings.Join(validDomains, ", "))
				}
			}
		}
	}
}

func TestFreeEmail(t *testing.T) {
	setupInternetTest()

	validDomains := internetProvider.FreeEmailDomains()
	t.Log(tests.Given("Given the need to generate random free email"))
	{
		for testID := 1; testID < 6; testID++ {
			t.Logf("\tTest %d:\tWhen creating a free email addresas.", testID)
			{
				email := internetFake.FreeEmail()

				if len(email) == 0 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould create valid email but it's empty."), testID)
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create valid username."), testID)

				if !tests.ValidEmail(email) {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould create valid email: %s."), testID, email)
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create valid email."), testID)

				emailParts := strings.Split(email, "@")
				if tests.InArray(emailParts[1], validDomains) == -1 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a valid email (%s) within proper domain range: \"%s\"."), testID, email, strings.Join(validDomains, ", "))
				}
			}
		}
	}
}

func TestCompanyEmail(t *testing.T) {
	setupInternetTest()

	t.Log(tests.Given("Given the need to generate random company email"))
	{
		for testID := 1; testID < 6; testID++ {
			t.Logf("\tTest %d:\tWhen creating an email addresas.", testID)
			{
				email := internetFake.CompanyEmail()

				if len(email) == 0 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould create valid email but it's empty."), testID)
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create valid email."), testID)

				if !tests.ValidEmail(email) {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould create valid email: %s."), testID, email)
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create valid email."), testID)
			}
		}
	}
}

func TestDomainName(t *testing.T) {
	t.Skip("Must implement better words for company")
}

func TestUrl(t *testing.T) {
	setupInternetTest()

	t.Log(tests.Given("Given the need to generate random URLs"))
	{
		for testID := 1; testID < 6; testID++ {
			t.Logf("\tTest %d:\tWhen creating an URL.", testID)
			{
				fakeUrl := internetFake.Url()

				if len(fakeUrl) == 0 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould create valid URL but it's empty."), testID)
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create valid URL."), testID)

				_, err := url.ParseRequestURI(fakeUrl)
				if err != nil {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould create valid URL: %s."), testID, fakeUrl)
				}
			}
		}
	}
}

func TestIPv4(t *testing.T) {
	setupInternetTest()

	t.Log(tests.Given("Given the need to generate random IPv4"))
	{
		for testID := 1; testID < 6; testID++ {
			t.Logf("\tTest %d:\tWhen creating an URL.", testID)
			{
				fakeIP := internetFake.IPv4()

				if len(fakeIP) == 0 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould create valid IPv4 but it's empty."), testID)
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create valid IPv4."), testID)

				if net.ParseIP(fakeIP) == nil {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould create valid IPv4: %s."), testID, fakeIP)
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create valid IPv4."), testID)
			}
		}
	}
}

func TestIPv6(t *testing.T) {
	setupInternetTest()

	t.Log(tests.Given("Given the need to generate random IPv4"))
	{
		for testID := 1; testID < 6; testID++ {
			t.Logf("\tTest %d:\tWhen creating an URL.", testID)
			{
				fakeIP := internetFake.IPv6()

				if len(fakeIP) == 0 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould create valid IPv6 but it's empty."), testID)
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create valid IPv6."), testID)

				if net.ParseIP(fakeIP) == nil {
					t.Fatalf(tests.Failed("\t", "Test %d:\tShould create valid IPv6: %s."), testID, fakeIP)
				}
				t.Logf(tests.Success("\t", "Test %d:\tShould create valid IPv6."), testID)
			}
		}
	}
}

func TestMacAddress(t *testing.T) {
	setupInternetTest()

	t.Log(tests.Given("Given the need to generate random IPv4"))
	{
		for testID := 1; testID < 6; testID++ {
			fakeMAC := internetFake.MacAddress()

			if len(fakeMAC) == 0 {
				t.Fatalf(tests.Failed("\t", "Test %d:\tShould create valid MacAddress but it's empty."), testID)
			}
			t.Logf(tests.Success("\t", "Test %d:\tShould create valid MacAddress."), testID)

			if _, err := net.ParseMAC(fakeMAC); err != nil {
				t.Fatalf(tests.Failed("\t", "Test %d:\tShould create MacAddress: %s - %v."), testID, fakeMAC, err)
			}
			t.Logf(tests.Success("\t", "Test %d:\tShould create valid MacAddress."), testID)
		}
	}
}
