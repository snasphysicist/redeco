package redeco

import "testing"

func TestPathParameterExtractionWithConversionWhenBoolFieldHasPathTags(t *testing.T) {
	g := NewFromString(pathTestSource("youngCall", "SorryRather", "bool", "takeUp"))
	s, err := g.Generate(options{handler: "fiveBear", target: "youngCall"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import chi "github.com/go-chi/chi/v5"
import "net/http"
import "strconv"

func fiveBearDecoder(r *http.Request) (youngCall, error) {
	var d youngCall
	var err error

	takeUp := chi.URLParam(r, "takeUp")
	takeUpConvert, err := strconv.ParseBool(takeUp)
	if err != nil {
		return d, err
	}
	d.SorryRather = takeUpConvert

	return d, err
}
`
	expectString(t, expect, ignoringGeneratedComment(s))
}
