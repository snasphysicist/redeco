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

import "fmt"
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
	expectString(t, expect, ignoringGeneratedComment(s))
}

func TestQueryParameterExtractionWithConversionWhenUint64FieldHasQueryTag(t *testing.T) {
	g := NewFromString(queryTestSource("stir", "seed", "uint64", "note"))
	s, err := g.Generate(options{handler: "Hook", target: "stir"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import "fmt"
import "net/http"
import "strconv"

func HookDecoder(r *http.Request) (stir, error) {
	var d stir
	var err error

	note := r.URL.Query()["note"]
	if len(note) != 1 {
		return d, fmt.Errorf("for query parameter 'note' expected 1 value, got '%v'", note)
	}
	noteConvert, err := strconv.ParseUint(note[0], 10, 64)
	if err != nil {
		return d, err
	}
	d.seed = uint64(noteConvert)

	return d, err
}
`
	expectString(t, expect, ignoringGeneratedComment(s))
}

func TestQueryParameterExtractionWithConversionWhenUint32FieldHasQueryTag(t *testing.T) {
	g := NewFromString(queryTestSource("Seize", "cross", "uint32", "raise"))
	s, err := g.Generate(options{handler: "show", target: "Seize"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import "fmt"
import "net/http"
import "strconv"

func showDecoder(r *http.Request) (Seize, error) {
	var d Seize
	var err error

	raise := r.URL.Query()["raise"]
	if len(raise) != 1 {
		return d, fmt.Errorf("for query parameter 'raise' expected 1 value, got '%v'", raise)
	}
	raiseConvert, err := strconv.ParseUint(raise[0], 10, 32)
	if err != nil {
		return d, err
	}
	d.cross = uint32(raiseConvert)

	return d, err
}
`
	expectString(t, expect, ignoringGeneratedComment(s))
}

func TestQueryParameterExtractionWithConversionWhenUint16FieldHasQueryTag(t *testing.T) {
	g := NewFromString(queryTestSource("cage", "Bench", "uint16", "truth"))
	s, err := g.Generate(options{handler: "carve", target: "cage"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import "fmt"
import "net/http"
import "strconv"

func carveDecoder(r *http.Request) (cage, error) {
	var d cage
	var err error

	truth := r.URL.Query()["truth"]
	if len(truth) != 1 {
		return d, fmt.Errorf("for query parameter 'truth' expected 1 value, got '%v'", truth)
	}
	truthConvert, err := strconv.ParseUint(truth[0], 10, 16)
	if err != nil {
		return d, err
	}
	d.Bench = uint16(truthConvert)

	return d, err
}
`
	expectString(t, expect, ignoringGeneratedComment(s))
}

func TestQueryParameterExtractionWithConversionWhenUint8FieldHasQueryTag(t *testing.T) {
	g := NewFromString(queryTestSource("snail", "salt", "uint8", "Hut"))
	s, err := g.Generate(options{handler: "fork", target: "snail"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import "fmt"
import "net/http"
import "strconv"

func forkDecoder(r *http.Request) (snail, error) {
	var d snail
	var err error

	Hut := r.URL.Query()["Hut"]
	if len(Hut) != 1 {
		return d, fmt.Errorf("for query parameter 'Hut' expected 1 value, got '%v'", Hut)
	}
	HutConvert, err := strconv.ParseUint(Hut[0], 10, 8)
	if err != nil {
		return d, err
	}
	d.salt = uint8(HutConvert)

	return d, err
}
`
	expectString(t, expect, ignoringGeneratedComment(s))
}
