package redeco

import (
	"testing"
)

func TestQueryParameterExtractionWithConversionWhenIntFieldHasQueryTag(t *testing.T) {
	g := NewFromString(queryTestSource("live", "state", "int", "pledge"))
	s, err := g.Generate(options{handler: "bee", target: "live"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import "net/http"
import "strconv"

func beeDecoder(r *http.Request) (live, error) {
	var d live
	var err error

	pledge := r.URL.Query()["pledge"]
	if len(pledge) != 1 {
		return d, fmt.Errorf("for query parameter 'pledge' expected 1 value, got '%v'", pledge)
	}
	pledgeConvert, err := strconv.ParseInt(pledge[0], 10, 64)
	if err != nil {
		return d, err
	}
	d.state = int(pledgeConvert)

	return d, err
}
`
	expectString(t, expect, s)
}
