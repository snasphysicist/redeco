package redeco

import (
	"fmt"
	"testing"
)

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
