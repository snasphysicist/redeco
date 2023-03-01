package redeco

import "testing"

func TestQueryParameterExtractionWithConversionWhenUntFieldHasOptionalQueryTag(t *testing.T) {
	g := NewFromString(optionalQueryTestSource("angel", "theme", "uint", "begin"))
	s, err := g.Generate(options{handler: "glory", target: "angel"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import "fmt"
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

import "fmt"
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

func TestQueryParameterExtractionWithConversionWhenUint32FieldHasOptionalQueryTag(t *testing.T) {
	g := NewFromString(optionalQueryTestSource("blind", "Exile", "uint32", "night"))
	s, err := g.Generate(options{handler: "throw", target: "blind"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import "fmt"
import "net/http"
import "strconv"

func throwDecoder(r *http.Request) (blind, error) {
	var d blind
	var err error

	night := r.URL.Query()["night"]
	if len(night) > 1 {
		return d, fmt.Errorf("for query parameter 'night' expected 0 or 1 value, got '%v'", night)
	}
	if len(night) == 1 {
		nightConvert, err := strconv.ParseUint(night[0], 10, 32)
		if err != nil {
			return d, err
		}
		d.Exile = uint32(nightConvert)
	}

	return d, err
}
`
	expectString(t, expect, s)
}

func TestQueryParameterExtractionWithConversionWhenUint16FieldHasOptionalQueryTag(t *testing.T) {
	g := NewFromString(optionalQueryTestSource("aloof", "touch", "uint16", "Laser"))
	s, err := g.Generate(options{handler: "thumb", target: "aloof"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import "fmt"
import "net/http"
import "strconv"

func thumbDecoder(r *http.Request) (aloof, error) {
	var d aloof
	var err error

	Laser := r.URL.Query()["Laser"]
	if len(Laser) > 1 {
		return d, fmt.Errorf("for query parameter 'Laser' expected 0 or 1 value, got '%v'", Laser)
	}
	if len(Laser) == 1 {
		LaserConvert, err := strconv.ParseUint(Laser[0], 10, 16)
		if err != nil {
			return d, err
		}
		d.touch = uint16(LaserConvert)
	}

	return d, err
}
`
	expectString(t, expect, s)
}

func TestQueryParameterExtractionWithConversionWhenUint8FieldHasOptionalQueryTag(t *testing.T) {
	g := NewFromString(optionalQueryTestSource("tower", "swarm", "uint8", "enemy"))
	s, err := g.Generate(options{handler: "Quote", target: "tower"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import "fmt"
import "net/http"
import "strconv"

func QuoteDecoder(r *http.Request) (tower, error) {
	var d tower
	var err error

	enemy := r.URL.Query()["enemy"]
	if len(enemy) > 1 {
		return d, fmt.Errorf("for query parameter 'enemy' expected 0 or 1 value, got '%v'", enemy)
	}
	if len(enemy) == 1 {
		enemyConvert, err := strconv.ParseUint(enemy[0], 10, 8)
		if err != nil {
			return d, err
		}
		d.swarm = uint8(enemyConvert)
	}

	return d, err
}
`
	expectString(t, expect, s)
}
