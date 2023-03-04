package redeco

import (
	"testing"
)

func TestStringPathParameterExtractionWithoutConversionWhenStructHasPathTags(t *testing.T) {
	g := NewFromString(pathTestSource("A", "A", "string", "a"))
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
	expectString(t, expect, ignoringGeneratedComment(s))
}
