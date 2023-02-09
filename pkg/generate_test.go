package redeco

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNoBodyDeserialisationWhenStructHasNoJSONTags(t *testing.T) {
	g := NewFromString(`
	package foo

	type A struct {}
	`)
	s, err := g.Generate(options{handler: "bar", target: "A"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import "net/http"

func barDecoder(r *http.Request) (A, error) {
	var d A
	var err error

	return d, err
}
`
	expectString(t, expect, s)
}

func TestBodyDeserialisationWhenStructHasJSONTags(t *testing.T) {
	g := NewFromString(fmt.Sprintf(`
	package foo

	type A struct {
		A string %sjson:"A"%s
	}
	`, "`", "`"))
	s, err := g.Generate(options{handler: "bar", target: "A"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import "io"
import "json"
import "net/http"

func barDecoder(r *http.Request) (A, error) {
	var d A
	var err error

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return d, err
	}
	err = json.Unmarshal(b, &d)
	if err != nil {
		return d, err
	}

	return d, err
}
`
	expectString(t, expect, s)
}

func TestStringPathParameterExtractionWithoutConversionWhenStructHasPathTags(t *testing.T) {
	g := NewFromString(fmt.Sprintf(`
	package foo

	type A struct {
		A string %spath:"a"%s
	}
	`, "`", "`"))
	s, err := g.Generate(options{handler: "bar", target: "A"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import chi "github.com/go-chi/chi/v5"
import "net/http"

func barDecoder(r *http.Request) (A, error) {
	var d A
	var err error

	a := chi.URLParam(r, "a")
	d.A = a

	return d, err
}
`
	expectString(t, expect, s)
}

func TestNumberPathParameterExtractionWithConversionWhenStructHasPathTags(t *testing.T) {
	g := NewFromString(fmt.Sprintf(`
	package foo

	type A struct {
		A int %spath:"a"%s
	}
	`, "`", "`"))
	s, err := g.Generate(options{handler: "bar", target: "A"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import chi "github.com/go-chi/chi/v5"
import "net/http"
import "strconv"

func barDecoder(r *http.Request) (A, error) {
	var d A
	var err error

	a := chi.URLParam(r, "a")
	aConvert, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		return d, err
	}
	d.A = int(aConvert)

	return d, err
}
`
	expectString(t, expect, s)
}

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
