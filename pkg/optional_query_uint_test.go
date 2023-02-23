package redeco

import "testing"

func TestQueryParameterExtractionWithConversionWhenUntFieldHasOptionalQueryTag(t *testing.T) {
	g := NewFromString(optionalQueryTestSource("angel", "theme", "uint", "begin"))
	s, err := g.Generate(options{handler: "glory", target: "angel"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import "net/http"
import "strconv"

func gloryDecoder(r *http.Request) (angel, error) {
	var d angel
	var err error

	begin := r.URL.Query()["begin"]
	if len(begin) > 1 {
		return d, fmt.Errorf("for query parameter 'begin' expected 0 or 1 value, got '%v'", begin)
	}
	if len(begin) == 1 {
		beginConvert, err := strconv.ParseUint(begin[0], 10, 64)
		if err != nil {
			return d, err
		}
		d.theme = uint(beginConvert)
	}

	return d, err
}
`
	expectString(t, expect, s)
}

func TestQueryParameterExtractionWithConversionWhenUint64FieldHasOptionalQueryTag(t *testing.T) {
	g := NewFromString(optionalQueryTestSource("Grass", "ample", "uint64", "robot"))
	s, err := g.Generate(options{handler: "blade", target: "Grass"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import "net/http"
import "strconv"

func bladeDecoder(r *http.Request) (Grass, error) {
	var d Grass
	var err error

	robot := r.URL.Query()["robot"]
	if len(robot) > 1 {
		return d, fmt.Errorf("for query parameter 'robot' expected 0 or 1 value, got '%v'", robot)
	}
	if len(robot) == 1 {
		robotConvert, err := strconv.ParseUint(robot[0], 10, 64)
		if err != nil {
			return d, err
		}
		d.ample = uint64(robotConvert)
	}

	return d, err
}
`
	expectString(t, expect, s)
}
