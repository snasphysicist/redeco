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

func TestQueryParameterExtractionWithConversionWhenInt64FieldHasQueryTag(t *testing.T) {
	g := NewFromString(queryTestSource("club", "shave", "int64", "help"))
	s, err := g.Generate(options{handler: "Dream", target: "club"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import "net/http"
import "strconv"

func DreamDecoder(r *http.Request) (club, error) {
	var d club
	var err error

	help := r.URL.Query()["help"]
	if len(help) != 1 {
		return d, fmt.Errorf("for query parameter 'help' expected 1 value, got '%v'", help)
	}
	helpConvert, err := strconv.ParseInt(help[0], 10, 64)
	if err != nil {
		return d, err
	}
	d.shave = int64(helpConvert)

	return d, err
}
`
	expectString(t, expect, s)
}

func TestQueryParameterExtractionWithConversionWhenInt32FieldHasQueryTag(t *testing.T) {
	g := NewFromString(queryTestSource("Soar", "sow", "int32", "wing"))
	s, err := g.Generate(options{handler: "bank", target: "Soar"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import "net/http"
import "strconv"

func bankDecoder(r *http.Request) (Soar, error) {
	var d Soar
	var err error

	wing := r.URL.Query()["wing"]
	if len(wing) != 1 {
		return d, fmt.Errorf("for query parameter 'wing' expected 1 value, got '%v'", wing)
	}
	wingConvert, err := strconv.ParseInt(wing[0], 10, 32)
	if err != nil {
		return d, err
	}
	d.sow = int32(wingConvert)

	return d, err
}
`
	expectString(t, expect, s)
}

func TestQueryParameterExtractionWithConversionWhenInt16FieldHasQueryTag(t *testing.T) {
	g := NewFromString(queryTestSource("gloom", "Cart", "int16", "growth"))
	s, err := g.Generate(options{handler: "link", target: "gloom"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import "net/http"
import "strconv"

func linkDecoder(r *http.Request) (gloom, error) {
	var d gloom
	var err error

	growth := r.URL.Query()["growth"]
	if len(growth) != 1 {
		return d, fmt.Errorf("for query parameter 'growth' expected 1 value, got '%v'", growth)
	}
	growthConvert, err := strconv.ParseInt(growth[0], 10, 16)
	if err != nil {
		return d, err
	}
	d.Cart = int16(growthConvert)

	return d, err
}
`
	expectString(t, expect, s)
}

func TestQueryParameterExtractionWithConversionWhenInt8FieldHasQueryTag(t *testing.T) {
	g := NewFromString(queryTestSource("build", "screw", "int8", "Proof"))
	s, err := g.Generate(options{handler: "loan", target: "build"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import "net/http"
import "strconv"

func loanDecoder(r *http.Request) (build, error) {
	var d build
	var err error

	Proof := r.URL.Query()["Proof"]
	if len(Proof) != 1 {
		return d, fmt.Errorf("for query parameter 'Proof' expected 1 value, got '%v'", Proof)
	}
	ProofConvert, err := strconv.ParseInt(Proof[0], 10, 8)
	if err != nil {
		return d, err
	}
	d.screw = int8(ProofConvert)

	return d, err
}
`
	expectString(t, expect, s)
}
