package redeco

import "testing"

func TestPathParameterExtractionWithConversionWhenUintFieldHasPathTags(t *testing.T) {
	g := NewFromString(pathTestSource("deep", "could", "uint", "head"))
	s, err := g.Generate(options{handler: "night", target: "deep"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import chi "github.com/go-chi/chi/v5"
import "net/http"
import "strconv"

func nightDecoder(r *http.Request) (deep, error) {
	var d deep
	var err error

	head := chi.URLParam(r, "head")
	headConvert, err := strconv.ParseUint(head, 10, 64)
	if err != nil {
		return d, err
	}
	d.could = uint(headConvert)

	return d, err
}
`
	expectString(t, expect, s)
}

func TestPathParameterExtractionWithConversionWhenUint64FieldHasPathTags(t *testing.T) {
	g := NewFromString(pathTestSource("heart", "gotta", "uint64", "all"))
	s, err := g.Generate(options{handler: "More", target: "heart"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import chi "github.com/go-chi/chi/v5"
import "net/http"
import "strconv"

func MoreDecoder(r *http.Request) (heart, error) {
	var d heart
	var err error

	all := chi.URLParam(r, "all")
	allConvert, err := strconv.ParseUint(all, 10, 64)
	if err != nil {
		return d, err
	}
	d.gotta = uint64(allConvert)

	return d, err
}
`
	expectString(t, expect, s)
}

func TestPathParameterExtractionWithConversionWhenUint32FieldHasPathTags(t *testing.T) {
	g := NewFromString(pathTestSource("Light", "fat", "uint32", "winter"))
	s, err := g.Generate(options{handler: "blade", target: "Light"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import chi "github.com/go-chi/chi/v5"
import "net/http"
import "strconv"

func bladeDecoder(r *http.Request) (Light, error) {
	var d Light
	var err error

	winter := chi.URLParam(r, "winter")
	winterConvert, err := strconv.ParseUint(winter, 10, 32)
	if err != nil {
		return d, err
	}
	d.fat = uint32(winterConvert)

	return d, err
}
`
	expectString(t, expect, s)
}

func TestPathParameterExtractionWithConversionWhenUint16FieldHasPathTags(t *testing.T) {
	g := NewFromString(pathTestSource("stupid", "Were", "uint16", "super"))
	s, err := g.Generate(options{handler: "hat", target: "stupid"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import chi "github.com/go-chi/chi/v5"
import "net/http"
import "strconv"

func hatDecoder(r *http.Request) (stupid, error) {
	var d stupid
	var err error

	super := chi.URLParam(r, "super")
	superConvert, err := strconv.ParseUint(super, 10, 16)
	if err != nil {
		return d, err
	}
	d.Were = uint16(superConvert)

	return d, err
}
`
	expectString(t, expect, s)
}
