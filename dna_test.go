package fake_test

import (
	"regexp"
	"strings"
	"testing"

	"github.com/rmsj/fake"
	"github.com/rmsj/fake/tests"
)

func TestDNASequence(t *testing.T) {

	f, err := fake.New()

	if err != nil {
		t.Fatal("error creating fake instance")
	}

	tt := []struct {
		name         string
		sequenceSize int
	}{
		{"Short DNA Sequence", 1000},
		{"Medium DNA Sequence", 10000},
		{"Long (ish) text", 65535},
	}

	t.Log(tests.Given("Given the need to generate random DNA Sequence"))
	{
		for testID, test := range tt {
			// start with test 1
			testID++
			tf := func(t *testing.T) {
				t.Logf("\tTest %d:\tWhen creating a %s of %d characters.", testID, test.name, test.sequenceSize)
				{
					dna, err := f.DNASequence(test.sequenceSize)
					if err != nil {
						t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created DNA sequences: %v"), testID, err)
					}

					if len(dna) == 0 {
						t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a valid DNA sequence but it's empty: %s"), testID, dna)
					}
					t.Logf(tests.Success("\t", "Test %d:\tShould have create a valid DNA Sequence"), testID)

					if len(dna) != test.sequenceSize {
						t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a valid DNA sequence with %d characters length: %d"), testID, test.sequenceSize, len(dna))
					}
					t.Logf(tests.Success("\t", "Test %d:\tShould have create a valid DNA Sequence with %d characters length"), testID, test.sequenceSize)

					// check characters in sequence
					chars := []string{"A", "C", "G", "T"}
					reg := strings.Join(chars, "")

					// last treatments - removing special characters, etc.
					re := regexp.MustCompile("[" + reg + "]+")
					result := re.ReplaceAllString(dna, "")
					if len(result) > 0 {
						t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a valid DNA sequence only using characters A, C, G and T: %s"), testID, result)
					}
				}
			}

			t.Run(test.name, tf)

		}
	}

	tt = []struct {
		name         string
		sequenceSize int
	}{
		{"Too small DNA Sequence", 1},
		{"Too small DNA Sequence", 22},
		{"Too small DNA Sequence", 99},
	}

	t.Log(tests.Given("Given the need to not generate too small SNA sequence"))
	{
		for testID, test := range tt {
			// start with test 1
			testID++
			tf := func(t *testing.T) {
				t.Logf("\tTest %d:\tWhen creating a %s of %d characters.", testID, test.name, test.sequenceSize)
				{
					dna, err := f.DNASequence(test.sequenceSize)
					if err == nil {
						t.Logf(tests.Success("\t", "Test %d:\tShould not create DNA sequences smaller than 100 characters. Created one with %d characters"), testID, test.sequenceSize)
					}

					if len(dna) != 0 {
						t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created a valid DNA sequence: %s"), testID, dna)
					}
					t.Logf(tests.Success("\t", "Test %d:\tShould not have created a valid DNA Sequence"), testID)
				}
			}

			t.Run(test.name, tf)

		}
	}
}
