package redeco

import (
	"fmt"
	"log"
)

// optionalQueryExtractCode generates code to extract and convert
// a query parameter associated with the provided field & tag
func optionalQueryExtractCode(g *generation, f field, t tag) string {
	switch f.typ {
	case "string":
		return optionalQueryStringExtractTemplate(t.values[0], f.name)
	case "bool":
		g.newImport(iport{path: "strconv"})
		return optionalQueryBoolExtractTemplate(t.values[0], f.name)
	case "int":
		g.newImport(iport{path: "strconv"})
		return optionalQueryIntExtractTemplate(t.values[0], f.name, "Int", 64, "int")
	case "int64":
		g.newImport(iport{path: "strconv"})
		return optionalQueryIntExtractTemplate(t.values[0], f.name, "Int", 64, "int64")
	case "int32":
		g.newImport(iport{path: "strconv"})
		return optionalQueryIntExtractTemplate(t.values[0], f.name, "Int", 32, "int32")
	case "int16":
		g.newImport(iport{path: "strconv"})
		return optionalQueryIntExtractTemplate(t.values[0], f.name, "Int", 16, "int16")
	case "int8":
		g.newImport(iport{path: "strconv"})
		return optionalQueryIntExtractTemplate(t.values[0], f.name, "Int", 8, "int8")
	case "uint":
		g.newImport(iport{path: "strconv"})
		return optionalQueryIntExtractTemplate(t.values[0], f.name, "Uint", 64, "uint")
	case "uint64":
		g.newImport(iport{path: "strconv"})
		return optionalQueryIntExtractTemplate(t.values[0], f.name, "Uint", 64, "uint64")
	case "uint32":
		g.newImport(iport{path: "strconv"})
		return optionalQueryIntExtractTemplate(t.values[0], f.name, "Uint", 32, "uint32")
	case "uint16":
		g.newImport(iport{path: "strconv"})
		return optionalQueryIntExtractTemplate(t.values[0], f.name, "Uint", 16, "uint16")
	case "uint8":
		g.newImport(iport{path: "strconv"})
		return optionalQueryIntExtractTemplate(t.values[0], f.name, "Uint", 8, "uint8")
	}
	log.Panicf("Don't know how to generate code for type: %s", f.typ)
	return ""
}

// optionalQueryStringExtractTemplate generates code extracting an optional string type query parameter
func optionalQueryStringExtractTemplate(param string, field string) string {
	v := safeVariableName(param)
	return fmt.Sprintf(`
	%s := r.URL.Query()["%s"]
	if len(%s) > 1 {
		return d, fmt.Errorf("for query parameter '%s' expected 0 or 1 value, got '%s'", %s)
	}
	if len(%s) == 1 {
		d.%s = %s[0]
	}
`, v, param, v, param, "%v", v, v, field, v)
}

// optionalQueryBoolExtractTemplate generates code extracting & converting
// a query parameter with a bool type
func optionalQueryBoolExtractTemplate(param string, field string) string {
	v := safeVariableName(param)
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
`, v, param, v, param, "%v", v, v, v, v, field, v)
}

// optionalQueryIntExtractTemplate generates code extracting & converting
// a query parameter with a (u)int type
func optionalQueryIntExtractTemplate(param string, field string, fn string, bits uint8, cast string) string {
	v := safeVariableName(param)
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
`, v, param, v, param, "%v", v, v, v, fn, v, bits, field, cast, v)
}
