package redeco

import "testing"

func TestPathParameterExtractionWithConversionWhenFloatFieldHasPathTags(t *testing.T) {
	g := NewFromString(pathTestSource("smooth", "rock", "float", "mind"))
	s, err := g.Generate(options{handler: "box", target: "smooth"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import chi "github.com/go-chi/chi/v5"
import "net/http"
import "strconv"

func boxDecoder(r *http.Request) (smooth, error) {
	var d smooth
	var err error

	mind := chi.URLParam(r, "mind")
	mindConvert, err := strconv.ParseFloat(mind, 64)
	if err != nil {
		return d, err
	}
	d.rock = float(mindConvert)

	return d, err
}
`
	expectString(t, expect, s)
}
