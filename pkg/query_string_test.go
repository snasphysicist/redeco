package redeco

import (
	"testing"
)

func TestStringQueryParameterExtractionWithoutConversionWhenStructHasQueryTags(t *testing.T) {
	g := NewFromString(queryTestSource("A", "B", "string", "C"))
	s, err := g.Generate(options{handler: "D", target: "A"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import "fmt"
import "net/http"

func DDecoder(r *http.Request) (A, error) {
	var d A
	var err error

	C := r.URL.Query()["C"]
	if len(C) != 1 {
		return d, fmt.Errorf("for query parameter 'C' expected 1 value, got '%v'", C)
	}
	d.B = C[0]

	return d, err
}
`
	expectString(t, expect, ignoringGeneratedComment(s))
}
