package redeco

import "testing"

func TestPathParameterExtractionWithConversionWhenUintFieldHasPathTags(t *testing.T) {
	g := NewFromString(pathTestSource("deep", "could", "uint", "head"))
	s, err := g.Generate(options{handler: "night", target: "deep"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import chi "github.com/go-chi/chi/v5"
import "net/http"
import "strconv"

func nightDecoder(r *http.Request) (deep, error) {
	var d deep
	var err error

	head := chi.URLParam(r, "head")
	headConvert, err := strconv.ParseUint(head, 10, 64)
	if err != nil {
		return d, err
	}
	d.could = uint(headConvert)

	return d, err
}
`
	expectString(t, expect, s)
}
