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
