package redeco

import (
	"fmt"
	"log"
)

// optionalQueryExtractCode generates code to extract and convert
// a query parameter associated with the provided field & tag
func optionalQueryExtractCode(g *generation, f field, t tag) string {
	switch f.typ {
	case "bool":
		g.newImport(iport{path: "strconv"})
		return optionalQueryBoolExtractTemplate(t.values[0], f.name)
	case "int":
		g.newImport(iport{path: "strconv"})
		return optionalQueryIntExtractTemplate(t.values[0], f.name, "Int", 64, "int")
	}
	log.Panicf("Don't know how to generate code for type: %s", f.typ)
	return ""
}

// optionalQueryBoolExtractTemplate generates code extracting & converting
// a query parameter with a bool type
func optionalQueryBoolExtractTemplate(param string, field string) string {
	return fmt.Sprintf(`
	%s := r.URL.Query()["%s"]
	if len(%s) > 1 {
		return d, fmt.Errorf("for query parameter '%s' expected 0 or 1 value, got '%s'", %s)
	}
	if len(%s) == 1 {
		%sConvert, err := strconv.ParseBool(%s[0])
		if err != nil {
			return d, err
		}
		d.%s = %sConvert
	}
`, param, param, param, param, "%v", param, param, param, param, field, param)
}

// optionalQueryIntExtractTemplate generates code extracting & converting
// a query parameter with a (u)int type
func optionalQueryIntExtractTemplate(param string, field string, fn string, bits uint8, cast string) string {
	return fmt.Sprintf(`
	%s := r.URL.Query()["%s"]
	if len(%s) > 1 {
		return d, fmt.Errorf("for query parameter '%s' expected 0 or 1 value, got '%s'", %s)
	}
	if len(%s) == 1 {
		%sConvert, err := strconv.Parse%s(%s[0], 10, %d)
		if err != nil {
			return d, err
		}
		d.%s = %s(%sConvert)
	}
`, param, param, param, param, "%v", param, param, param, fn, param, bits, field, cast, param)
}
