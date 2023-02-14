package redeco

import (
	"testing"
)

func TestPathParameterExtractionWithConversionWhenIntFieldHasPathTags(t *testing.T) {
	g := NewFromString(pathTestSource("A", "A", "int", "a"))
	s, err := g.Generate(options{handler: "bar", target: "A"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import chi "github.com/go-chi/chi/v5"
import "net/http"
import "strconv"

func barDecoder(r *http.Request) (A, error) {
	var d A
	var err error

	a := chi.URLParam(r, "a")
	aConvert, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		return d, err
	}
	d.A = int(aConvert)

	return d, err
}
`
	expectString(t, expect, s)
}

func TestPathParameterExtractionWithConversionWhenInt32FieldHasPathTags(t *testing.T) {
	g := NewFromString(pathTestSource("B", "C", "int32", "e"))
	s, err := g.Generate(options{handler: "bar", target: "B"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import chi "github.com/go-chi/chi/v5"
import "net/http"
import "strconv"

func barDecoder(r *http.Request) (B, error) {
	var d B
	var err error

	e := chi.URLParam(r, "e")
	eConvert, err := strconv.ParseInt(e, 10, 32)
	if err != nil {
		return d, err
	}
	d.C = int32(eConvert)

	return d, err
}
`
	expectString(t, expect, s)
}

func TestPathParameterExtractionWithConversionWhenInt64FieldHasPathTags(t *testing.T) {
	g := NewFromString(pathTestSource("F", "g", "int64", "h"))
	s, err := g.Generate(options{handler: "bar", target: "F"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import chi "github.com/go-chi/chi/v5"
import "net/http"
import "strconv"

func barDecoder(r *http.Request) (F, error) {
	var d F
	var err error

	h := chi.URLParam(r, "h")
	hConvert, err := strconv.ParseInt(h, 10, 64)
	if err != nil {
		return d, err
	}
	d.g = int64(hConvert)

	return d, err
}
`
	expectString(t, expect, s)
}

func TestPathParameterExtractionWithConversionWhenInt16FieldHasPathTags(t *testing.T) {
	g := NewFromString(pathTestSource("I", "K", "int16", "lama"))
	s, err := g.Generate(options{handler: "bar", target: "I"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import chi "github.com/go-chi/chi/v5"
import "net/http"
import "strconv"

func barDecoder(r *http.Request) (I, error) {
	var d I
	var err error

	lama := chi.URLParam(r, "lama")
	lamaConvert, err := strconv.ParseInt(lama, 10, 16)
	if err != nil {
		return d, err
	}
	d.K = int16(lamaConvert)

	return d, err
}
`
	expectString(t, expect, s)
}

func TestPathParameterExtractionWithConversionWhenInt8FieldHasPathTags(t *testing.T) {
	g := NewFromString(pathTestSource("your", "exhaust", "int8", "style"))
	s, err := g.Generate(options{handler: "bar", target: "your"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import chi "github.com/go-chi/chi/v5"
import "net/http"
import "strconv"

func barDecoder(r *http.Request) (your, error) {
	var d your
	var err error

	style := chi.URLParam(r, "style")
	styleConvert, err := strconv.ParseInt(style, 10, 8)
	if err != nil {
		return d, err
	}
	d.exhaust = int8(styleConvert)

	return d, err
}
`
	expectString(t, expect, s)
}
