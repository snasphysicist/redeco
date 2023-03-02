package redeco

import "testing"

func TestQueryParameterExtractionWithConversionWhenBoolFieldHasQueryTag(t *testing.T) {
	g := NewFromString(queryTestSource("SheetDrug", "GrudgeLock", "bool", "ClueSoil"))
	s, err := g.Generate(options{handler: "ScreamHeart", target: "SheetDrug"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
	expect := `package foo

import "fmt"
import "net/http"
import "strconv"

func ScreamHeartDecoder(r *http.Request) (SheetDrug, error) {
	var d SheetDrug
	var err error

	ClueSoil := r.URL.Query()["ClueSoil"]
	if len(ClueSoil) != 1 {
		return d, fmt.Errorf("for query parameter 'ClueSoil' expected 1 value, got '%v'", ClueSoil)
	}
	ClueSoilConvert, err := strconv.ParseBool(ClueSoil[0])
	if err != nil {
		return d, err
	}
	d.GrudgeLock = ClueSoilConvert

	return d, err
}
`
	expectString(t, expect, s)
}
