package redeco

import (
	"encoding/json"
	"fmt"
	"strings"
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

// pathTestSource generates the source code for path parameter test
func pathTestSource(name string, field string, typ string, param string) string {
	return fmt.Sprintf(`
	package foo

	type %s struct {
		%s %s %spath:"%s"%s
	}
	`, name, field, typ, "`", param, "`")
}

// queryTestSource generates the source code for query parameter tests
func queryTestSource(name string, field string, typ string, param string) string {
	return fmt.Sprintf(`
	package foo

	type %s struct {
		%s %s %squery:"%s"%s
	}
	`, name, field, typ, "`", param, "`")
}

// optionalQueryTestSource generates the source code for query parameter tests
func optionalQueryTestSource(name string, field string, typ string, param string) string {
	return fmt.Sprintf(`
	package foo

	type %s struct {
		%s %s %squery:"%s,optional"%s
	}
	`, name, field, typ, "`", param, "`")
}

// ignoringGeneratedComment strips out the "Code generated" comment
// to focus on the actual generated code
func ignoringGeneratedComment(s string) string {
	ls := strings.Split(s, "\n")
	ls = filter(ls, func(l string) bool {
		return !strings.Contains(l, "Code generated") && !strings.Contains(l, "DO NOT EDIT")
	})
	ls = ls[:len(ls)-1] // removes an extra newline added for the Code generated comment
	return strings.Join(ls, "\n")
}
