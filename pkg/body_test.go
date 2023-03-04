package redeco

import (
	"fmt"
	"strings"
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
	expectString(t, expect, ignoringGeneratedComment(s))
}

func TestGeneratedFileContainsACommentLineIndicatingTheFileContainsGeneratedCode(t *testing.T) {
	g := NewFromString(`
	package foo

	type A struct {}
	`)
	s, err := g.Generate(options{handler: "bar", target: "A"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	ls := strings.Split(s, "\n")
	c := filter(ls, func(l string) bool {
		return strings.HasPrefix(l, "// Code generated ") && strings.HasSuffix(l, " DO NOT EDIT.")
	})
	if len(c) == 0 {
		t.Errorf("Output does not contain a generated code indicator comment, lines were %v", ls)
	}
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
	expectString(t, expect, ignoringGeneratedComment(s))
}
