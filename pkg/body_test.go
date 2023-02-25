package redeco

import (
	"fmt"
	"testing"
)

func TestNoBodyDeserialisationWhenStructHasNoJSONTags(t *testing.T) {
	g := NewFromString(`
	package foo

	type A struct {}
	`)
	s, err := g.Generate(options{handler: "bar", target: "A"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import "net/http"

func barDecoder(r *http.Request) (A, error) {
	var d A
	var err error

	return d, err
}
`
	expectString(t, expect, s)
}

func TestBodyDeserialisationWhenStructHasJSONTags(t *testing.T) {
	g := NewFromString(fmt.Sprintf(`
	package foo

	type A struct {
		A string %sjson:"A"%s
	}
	`, "`", "`"))
	s, err := g.Generate(options{handler: "bar", target: "A"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import "encoding/json"
import "io"
import "net/http"

func barDecoder(r *http.Request) (A, error) {
	var d A
	var err error

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return d, err
	}
	err = json.Unmarshal(b, &d)
	if err != nil {
		return d, err
	}

	return d, err
}
`
	expectString(t, expect, s)
}
