package redeco

import (
	"fmt"
	"log"
	"strings"
)

// queryDeserialiseCode generates the code for deserialising
// parameters from the request query parameters
func queryDeserialiseCode(g *generation) (string, error) {
	s, err := targetStruct(g)
	if err != nil {
		return "", err
	}
	f := queryTaggedFields(s)
	if len(f) == 0 {
		log.Printf("No query tags in target struct, won't look for query parameters")
		return "", nil
	}
	c := mapping(f, func(f field) string { return queryExtractCode(g, f) })
	return strings.Join(c, "\n"), nil
}

// queryTaggedFields returns any fields in the struct with query tags
func queryTaggedFields(s sourceStruct) []field {
	return filter(s.field, func(f field) bool {
		return anyMatch(f.tags, func(t tag) bool { return t.key == "query" })
	})
}

// queryExtractCode generates code to extract the query parameter associated with f
func queryExtractCode(g *generation, f field) string {
	t := filter(f.tags, func(t tag) bool { return t.key == "query" })
	if len(t) != 1 {
		log.Panicf("Could not find unique path tag in %#v", f)
	}
	switch f.typ {
	case "string":
		return fmt.Sprintf(
			queryExtractTemplate,
			t[0].values[0],
			t[0].values[0],
			t[0].values[0],
			t[0].values[0],
			"%v",
			t[0].values[0],
			f.name,
			t[0].values[0],
		)
	case "int":
		g.newImport(iport{path: "strconv"})
		return queryIntExtractTemplate(t[0].values[0], f.name, "Int", 64, "int")
	case "int64":
		g.newImport(iport{path: "strconv"})
		return queryIntExtractTemplate(t[0].values[0], f.name, "Int", 64, "int64")
	case "int32":
		g.newImport(iport{path: "strconv"})
		return queryIntExtractTemplate(t[0].values[0], f.name, "Int", 32, "int32")
	case "int16":
		g.newImport(iport{path: "strconv"})
		return queryIntExtractTemplate(t[0].values[0], f.name, "Int", 16, "int16")
	case "int8":
		g.newImport(iport{path: "strconv"})
		return queryIntExtractTemplate(t[0].values[0], f.name, "Int", 8, "int8")
	case "uint":
		g.newImport(iport{path: "strconv"})
		return queryIntExtractTemplate(t[0].values[0], f.name, "Uint", 64, "uint")
	case "uint64":
		g.newImport(iport{path: "strconv"})
		return queryIntExtractTemplate(t[0].values[0], f.name, "Uint", 64, "uint64")
	case "uint32":
		g.newImport(iport{path: "strconv"})
		return queryIntExtractTemplate(t[0].values[0], f.name, "Uint", 32, "uint32")
	case "uint16":
		g.newImport(iport{path: "strconv"})
		return queryIntExtractTemplate(t[0].values[0], f.name, "Uint", 16, "uint16")
	case "uint8":
		g.newImport(iport{path: "strconv"})
		return queryIntExtractTemplate(t[0].values[0], f.name, "Uint", 8, "uint8")
	case "bool":
		g.newImport(iport{path: "strconv"})
		return queryBoolExtractTemplate(t[0].values[0], f.name)
	}
	log.Panicf("Don't know how to convert type '%s'", f.typ)
	return ""
}

// queryExtractTemplate is the template code for extracting a path parameter
const queryExtractTemplate = `
	%s := r.URL.Query()["%s"]
	if len(%s) != 1 {
		return d, fmt.Errorf("for query parameter '%s' expected 1 value, got '%s'", %s)
	}
	d.%s = %s[0]
`

// queryUintExtractTemplate generates code extracting & converting
// a query parameter with a (u)int* type
func queryIntExtractTemplate(param string, field string, parse string, bits uint8, typ string) string {
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

// queryBoolExtractTemplate generates code extracting & converting
// a query parameter with a bool type
func queryBoolExtractTemplate(param string, field string) string {
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
