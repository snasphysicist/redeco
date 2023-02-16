package redeco

import (
	"testing"
)

func TestQueryParameterExtractionWithConversionWhenUintFieldHasQueryTag(t *testing.T) {
	g := NewFromString(queryTestSource("stroke", "hook", "uint", "worth"))
	s, err := g.Generate(options{handler: "dorm", target: "stroke"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import "net/http"
import "strconv"

func dormDecoder(r *http.Request) (stroke, error) {
	var d stroke
	var err error

	worth := r.URL.Query()["worth"]
	if len(worth) != 1 {
		return d, fmt.Errorf("for query parameter 'worth' expected 1 value, got '%v'", worth)
	}
	worthConvert, err := strconv.ParseUint(worth[0], 10, 64)
	if err != nil {
		return d, err
	}
	d.hook = uint(worthConvert)

	return d, err
}
`
	expectString(t, expect, s)
}
