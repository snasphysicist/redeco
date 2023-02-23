package redeco

import (
	"fmt"
	"log"
)

// requiredQueryExtractCode generates code to extract and convert
// a query parameter associated with the provided field & tag
func requiredQueryExtractCode(g *generation, f field, t tag) string {
	switch f.typ {
	case "string":
		return fmt.Sprintf(
			requiredQueryStringExtractTemplate,
			t.values[0],
			t.values[0],
			t.values[0],
			t.values[0],
			"%v",
			t.values[0],
			f.name,
			t.values[0],
		)
	case "int":
		g.newImport(iport{path: "strconv"})
		return requiredQueryIntExtractTemplate(t.values[0], f.name, "Int", 64, "int")
	case "int64":
		g.newImport(iport{path: "strconv"})
		return requiredQueryIntExtractTemplate(t.values[0], f.name, "Int", 64, "int64")
	case "int32":
		g.newImport(iport{path: "strconv"})
		return requiredQueryIntExtractTemplate(t.values[0], f.name, "Int", 32, "int32")
	case "int16":
		g.newImport(iport{path: "strconv"})
		return requiredQueryIntExtractTemplate(t.values[0], f.name, "Int", 16, "int16")
	case "int8":
		g.newImport(iport{path: "strconv"})
		return requiredQueryIntExtractTemplate(t.values[0], f.name, "Int", 8, "int8")
	case "uint":
		g.newImport(iport{path: "strconv"})
		return requiredQueryIntExtractTemplate(t.values[0], f.name, "Uint", 64, "uint")
	case "uint64":
		g.newImport(iport{path: "strconv"})
		return requiredQueryIntExtractTemplate(t.values[0], f.name, "Uint", 64, "uint64")
	case "uint32":
		g.newImport(iport{path: "strconv"})
		return requiredQueryIntExtractTemplate(t.values[0], f.name, "Uint", 32, "uint32")
	case "uint16":
		g.newImport(iport{path: "strconv"})
		return requiredQueryIntExtractTemplate(t.values[0], f.name, "Uint", 16, "uint16")
	case "uint8":
		g.newImport(iport{path: "strconv"})
		return requiredQueryIntExtractTemplate(t.values[0], f.name, "Uint", 8, "uint8")
	case "bool":
		g.newImport(iport{path: "strconv"})
		return requiredQueryBoolExtractTemplate(t.values[0], f.name)
	}
	log.Panicf("Don't know how to convert type '%s'", f.typ)
	return ""
}

// requiredQueryStringExtractTemplate is the template code for extracting a string path parameter
const requiredQueryStringExtractTemplate = `
	%s := r.URL.Query()["%s"]
	if len(%s) != 1 {
		return d, fmt.Errorf("for query parameter '%s' expected 1 value, got '%s'", %s)
	}
	d.%s = %s[0]
`

// requiredQueryIntExtractTemplate generates code extracting & converting
// a query parameter with a (u)int* type
func requiredQueryIntExtractTemplate(param string, field string, parse string, bits uint8, typ string) string {
	return fmt.Sprintf(`
	%s := r.URL.Query()["%s"]
	if len(%s) != 1 {
		return d, fmt.Errorf("for query parameter '%s' expected 1 value, got '%s'", %s)
	}
	%sConvert, err := strconv.Parse%s(%s[0], 10, %d)
	if err != nil {
		return d, err
	}
	d.%s = %s(%sConvert)
`, param, param, param, param, "%v", param, param, parse, param, bits, field, typ, param)
}

// requiredQueryBoolExtractTemplate generates code extracting & converting
// a query parameter with a bool type
func requiredQueryBoolExtractTemplate(param string, field string) string {
	return fmt.Sprintf(`
	%s := r.URL.Query()["%s"]
	if len(%s) != 1 {
		return d, fmt.Errorf("for query parameter '%s' expected 1 value, got '%s'", %s)
	}
	%sConvert, err := strconv.ParseBool(%s[0])
	if err != nil {
		return d, err
	}
	d.%s = %sConvert
`, param, param, param, param, "%v", param, param, param, field, param)
}
