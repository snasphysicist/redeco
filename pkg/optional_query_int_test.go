package redeco

import "testing"

func TestQueryParameterExtractionWithConversionWhenIntFieldHasOptionalQueryTag(t *testing.T) {
	g := NewFromString(optionalQueryTestSource("hobby", "drink", "int", "march"))
	s, err := g.Generate(options{handler: "crime", target: "hobby"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import "net/http"
import "strconv"

func crimeDecoder(r *http.Request) (hobby, error) {
	var d hobby
	var err error

	march := r.URL.Query()["march"]
	if len(march) > 1 {
		return d, fmt.Errorf("for query parameter 'march' expected 0 or 1 value, got '%v'", march)
	}
	if len(march) == 1 {
		marchConvert, err := strconv.ParseInt(march[0], 10, 64)
		if err != nil {
			return d, err
		}
		d.drink = int(marchConvert)
	}

	return d, err
}
`
	expectString(t, expect, s)
}
