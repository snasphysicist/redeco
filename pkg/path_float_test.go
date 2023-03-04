package redeco

import "testing"

func TestPathParameterExtractionWithConversionWhenFloatFieldHasPathTags(t *testing.T) {
	g := NewFromString(pathTestSource("smooth", "rock", "float", "mind"))
	s, err := g.Generate(options{handler: "box", target: "smooth"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import chi "github.com/go-chi/chi/v5"
import "net/http"
import "strconv"

func boxDecoder(r *http.Request) (smooth, error) {
	var d smooth
	var err error

	mind := chi.URLParam(r, "mind")
	mindConvert, err := strconv.ParseFloat(mind, 64)
	if err != nil {
		return d, err
	}
	d.rock = float(mindConvert)

	return d, err
}
`
	expectString(t, expect, ignoringGeneratedComment(s))
}

func TestPathParameterExtractionWithConversionWhenFloat64FieldHasPathTags(t *testing.T) {
	g := NewFromString(pathTestSource("alive", "cat", "float64", "day"))
	s, err := g.Generate(options{handler: "upset", target: "alive"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import chi "github.com/go-chi/chi/v5"
import "net/http"
import "strconv"

func upsetDecoder(r *http.Request) (alive, error) {
	var d alive
	var err error

	day := chi.URLParam(r, "day")
	dayConvert, err := strconv.ParseFloat(day, 64)
	if err != nil {
		return d, err
	}
	d.cat = float64(dayConvert)

	return d, err
}
`
	expectString(t, expect, ignoringGeneratedComment(s))
}

func TestPathParameterExtractionWithConversionWhenFloat32FieldHasPathTags(t *testing.T) {
	g := NewFromString(pathTestSource("begin", "hands", "float32", "way"))
	s, err := g.Generate(options{handler: "french", target: "begin"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import chi "github.com/go-chi/chi/v5"
import "net/http"
import "strconv"

func frenchDecoder(r *http.Request) (begin, error) {
	var d begin
	var err error

	way := chi.URLParam(r, "way")
	wayConvert, err := strconv.ParseFloat(way, 32)
	if err != nil {
		return d, err
	}
	d.hands = float32(wayConvert)

	return d, err
}
`
	expectString(t, expect, ignoringGeneratedComment(s))
}
