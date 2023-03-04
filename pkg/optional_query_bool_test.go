package redeco

import "testing"

func TestQueryParameterExtractionWithConversionWhenBoolFieldHasOptionalQueryTag(t *testing.T) {
	g := NewFromString(optionalQueryTestSource("FeverRange", "roundAmuse", "bool", "RebelFrame"))
	s, err := g.Generate(options{handler: "snailReign", target: "FeverRange"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import "fmt"
import "net/http"
import "strconv"

func snailReignDecoder(r *http.Request) (FeverRange, error) {
	var d FeverRange
	var err error

	RebelFrame := r.URL.Query()["RebelFrame"]
	if len(RebelFrame) > 1 {
		return d, fmt.Errorf("for query parameter 'RebelFrame' expected 0 or 1 value, got '%v'", RebelFrame)
	}
	if len(RebelFrame) == 1 {
		RebelFrameConvert, err := strconv.ParseBool(RebelFrame[0])
		if err != nil {
			return d, err
		}
		d.roundAmuse = RebelFrameConvert
	}

	return d, err
}
`
	expectString(t, expect, ignoringGeneratedComment(s))
}
