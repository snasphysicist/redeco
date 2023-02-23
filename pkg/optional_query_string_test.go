package redeco

import "testing"

func TestOptionalStringQueryParameterExtractionWithoutConversionWhenStructHasQueryTags(t *testing.T) {
	g := NewFromString(optionalQueryTestSource("p", "q", "string", "t"))
	s, err := g.Generate(options{handler: "s", target: "p"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import "net/http"

func sDecoder(r *http.Request) (p, error) {
	var d p
	var err error

	t := r.URL.Query()["t"]
	if len(t) > 1 {
		return d, fmt.Errorf("for query parameter 't' expected 0 or 1 value, got '%v'", t)
	}
	if len(t) == 1 {
		d.q = t[0]
	}

	return d, err
}
`
	expectString(t, expect, s)
}
