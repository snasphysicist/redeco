package redeco

import "testing"

func TestQueryParameterExtractionWithConversionWhenIntFieldHasOptionalQueryTag(t *testing.T) {
	g := NewFromString(optionalQueryTestSource("hobby", "drink", "int", "march"))
	s, err := g.Generate(options{handler: "crime", target: "hobby"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import "fmt"
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
	expectString(t, expect, ignoringGeneratedComment(s))
}

func TestQueryParameterExtractionWithConversionWhenInt64FieldHasOptionalQueryTag(t *testing.T) {
	g := NewFromString(optionalQueryTestSource("Aware", "spend", "int64", "world"))
	s, err := g.Generate(options{handler: "stage", target: "Aware"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import "fmt"
import "net/http"
import "strconv"

func stageDecoder(r *http.Request) (Aware, error) {
	var d Aware
	var err error

	world := r.URL.Query()["world"]
	if len(world) > 1 {
		return d, fmt.Errorf("for query parameter 'world' expected 0 or 1 value, got '%v'", world)
	}
	if len(world) == 1 {
		worldConvert, err := strconv.ParseInt(world[0], 10, 64)
		if err != nil {
			return d, err
		}
		d.spend = int64(worldConvert)
	}

	return d, err
}
`
	expectString(t, expect, ignoringGeneratedComment(s))
}

func TestQueryParameterExtractionWithConversionWhenInt32FieldHasOptionalQueryTag(t *testing.T) {
	g := NewFromString(optionalQueryTestSource("pause", "Budge", "int32", "think"))
	s, err := g.Generate(options{handler: "hover", target: "pause"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import "fmt"
import "net/http"
import "strconv"

func hoverDecoder(r *http.Request) (pause, error) {
	var d pause
	var err error

	think := r.URL.Query()["think"]
	if len(think) > 1 {
		return d, fmt.Errorf("for query parameter 'think' expected 0 or 1 value, got '%v'", think)
	}
	if len(think) == 1 {
		thinkConvert, err := strconv.ParseInt(think[0], 10, 32)
		if err != nil {
			return d, err
		}
		d.Budge = int32(thinkConvert)
	}

	return d, err
}
`
	expectString(t, expect, ignoringGeneratedComment(s))
}

func TestQueryParameterExtractionWithConversionWhenInt16FieldHasOptionalQueryTag(t *testing.T) {
	g := NewFromString(optionalQueryTestSource("charm", "river", "int16", "Doubt"))
	s, err := g.Generate(options{handler: "smash", target: "charm"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import "fmt"
import "net/http"
import "strconv"

func smashDecoder(r *http.Request) (charm, error) {
	var d charm
	var err error

	Doubt := r.URL.Query()["Doubt"]
	if len(Doubt) > 1 {
		return d, fmt.Errorf("for query parameter 'Doubt' expected 0 or 1 value, got '%v'", Doubt)
	}
	if len(Doubt) == 1 {
		DoubtConvert, err := strconv.ParseInt(Doubt[0], 10, 16)
		if err != nil {
			return d, err
		}
		d.river = int16(DoubtConvert)
	}

	return d, err
}
`
	expectString(t, expect, ignoringGeneratedComment(s))
}

func TestQueryParameterExtractionWithConversionWhenInt8FieldHasOptionalQueryTag(t *testing.T) {
	g := NewFromString(optionalQueryTestSource("amuse", "choke", "int8", "smoke"))
	s, err := g.Generate(options{handler: "Medal", target: "amuse"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import "fmt"
import "net/http"
import "strconv"

func MedalDecoder(r *http.Request) (amuse, error) {
	var d amuse
	var err error

	smoke := r.URL.Query()["smoke"]
	if len(smoke) > 1 {
		return d, fmt.Errorf("for query parameter 'smoke' expected 0 or 1 value, got '%v'", smoke)
	}
	if len(smoke) == 1 {
		smokeConvert, err := strconv.ParseInt(smoke[0], 10, 8)
		if err != nil {
			return d, err
		}
		d.choke = int8(smokeConvert)
	}

	return d, err
}
`
	expectString(t, expect, ignoringGeneratedComment(s))
}
