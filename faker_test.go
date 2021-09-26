package faker_test

import (
	"github.com/rmsj/faker"
	"github.com/rmsj/faker/provider"
	"github.com/rmsj/faker/tests"
	"reflect"
	"testing"
)

type user struct {
	firstName string
	lastName  string
	email     string
}

func TestFactory(t *testing.T) {

	pp := provider.NewEnglishPersonProvider()
	ip := provider.NewEnglishInternetProvider()

	f, err := faker.New(pp, ip)
	if err != nil {
		t.Fatal("initiating faker", err)
	}

	builder := func() interface{} {
		return user{
			firstName: f.FirstName(),
			lastName:  f.LastName(),
			email:     f.Email(),
		}
	}

	tt := []struct {
		name     string
		quantity int
	}{
		{"Single Value", 1},
		{"Multiple Values", 50},
		{"Loads of Values", 200},
	}

	t.Log("Given the need to create a variable number of value of a type with the Factory function")
	{
		for testID, test := range tt {
			// start with test 1
			testID++
			tf := func(t *testing.T) {
				t.Parallel()
				t.Logf("\tTest %d:\tWhen creating %d value of type user.", testID, test.quantity)
				{
					users := f.Factory(builder, test.quantity)

					if len(users) != test.quantity {
						t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created %d user: %d created."), testID, test.quantity, len(users))
					}
					t.Logf(tests.Success("\t", "Test %d:\tShould create %d value(s) of type user."), testID, test.quantity)

					for _, v := range users {
						u, ok := v.(user)
						if !ok {
							t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a value of type user: %v created."), testID, reflect.TypeOf(v))
						}

						if len(u.firstName) == 0 {
							t.Fatalf(tests.Failed("\t", "Test %d:\tShould have set user first name: empty."), testID)
						}

						if len(u.lastName) == 0 {
							t.Fatalf(tests.Failed("\t", "Test %d:\tShould have set user last name: empty."), testID)
						}

						if len(u.email) == 0 {
							t.Fatalf(tests.Failed("\t", "Test %d:\tShould have set user email: empty."), testID)
						}

						if !tests.ValidEmail(u.email) {
							t.Fatalf(tests.Failed("\t", "Test %d:\tShould have set user email with valid email address: %s."), testID, u.email)
						}
					}
					t.Logf(tests.Success("\t", "Test %d:\tShould set all user values correctly for all %d value(s)"), testID, test.quantity)

					v := users[0]
					u, ok := v.(user)
					if !ok {
						t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a value of type user: %v created."), testID, reflect.TypeOf(v))
					}
					t.Logf(tests.Success("\t", "Test %d:\tShould create value of type user."), testID)

					if len(u.firstName) == 0 {
						t.Fatalf(tests.Failed("\t", "Test %d:\tShould have set user first name: empty."), testID)
					}
					t.Logf(tests.Success("\t", "Test %d:\tShould set user first name: %s"), testID, u.firstName)

					if len(u.lastName) == 0 {
						t.Fatalf(tests.Failed("\t", "Test %d:\tShould have set user last name: empty."), testID)
					}
					t.Logf(tests.Success("\t", "Test %d:\tShould set user last name: %s"), testID, u.lastName)

					if len(u.email) == 0 {
						t.Fatalf(tests.Failed("\t", "Test %d:\tShould have set user email: empty."), testID)
					}
					t.Logf(tests.Success("\t", "Test %d:\tShould set user email: %s"), testID, u.email)

					if !tests.ValidEmail(u.email) {
						t.Fatalf(tests.Failed("\t", "Test %d:\tShould have set user email with valid email address: %s."), testID, u.email)
					}
					t.Logf(tests.Success("\t", "Test %d:\tShould set user email with valid email address: %s"), testID, u.email)
				}
			}

			t.Run(test.name, tf)

		}
	}
}
