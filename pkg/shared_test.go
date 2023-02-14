package redeco

import (
	"encoding/json"
	"testing"
)

// expectString fails the test if the two strings are not equal
// and prints in detail which characters do not match
func expectString(t *testing.T, expect string, actual string) {
	if len(expect) != len(actual) {
		t.Errorf("Strings are of different lengths: %d != %d", len(actual), len(expect))
	}
	if expect != actual {
		t.Errorf("Actual '%s' != expected '%s'", actual, expect)
	}
	for i := 0; i < min(len(expect), len(actual)); i++ {
		if expect[i] != actual[i] {
			t.Errorf("Difference at %d: '%s' != '%s'", i,
				mustMarshal(actual[i:i+1]), mustMarshal(expect[i:i+1]))
		}
	}
}

// min returns the smaller of the two values
func min(l int, r int) int {
	if l < r {
		return l
	}
	return r
}

// mustMarshal json.Marshal s the input to string, panics on failure
func mustMarshal(i interface{}) string {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	return string(b)
}
