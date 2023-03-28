package redeco

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// pathDeserialiseCode generates the code for deserialising
// parameters from the request path
func pathDeserialiseCode(g *generation) (string, error) {
	s, err := targetStruct(g)
	if err != nil {
		return "", err
	}
	f := pathTaggedFields(s)
	if len(f) == 0 {
		log.Printf("No path tags in target struct, won't look for path parameters")
		return "", nil
	}
	g.newImport(iport{alias: "chi", path: "github.com/go-chi/chi/v5"})
	c, err := mapWithError(f, func(f field) (string, error) { return pathExtractCode(g, f) })
	return strings.Join(c, "\n"), err
}

// pathTaggedFields returns any fields in the struct with path tags
func pathTaggedFields(s sourceStruct) []field {
	return filter(s.field, func(f field) bool {
		return anyMatch(f.tags, func(t tag) bool { return t.key == "path" })
	})
}

// pathExtractCode generates code to extract the path parameter associated with f
func pathExtractCode(g *generation, f field) (string, error) {
	t := filter(f.tags, func(t tag) bool { return t.key == "path" })
	if len(t) != 1 {
		log.Panicf("Could not find unique path tag in %#v", f)
	}
	v := safeVariableName(t[0].values[0])
	cc, err := convertCode(g, f, t[0])
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(
		pathExtractTemplate,
		v,
		t[0].values[0],
		cc,
	), nil
}

// convertCode generates the code to convert the path parameter to the correct type
func convertCode(g *generation, f field, t tag) (string, error) {
	if strings.HasPrefix(f.typ, "int") {
		return convertIntCode(g, f, t)
	}
	if strings.HasPrefix(f.typ, "uint") {
		return convertUintCode(g, f, t)
	}
	if strings.HasPrefix(f.typ, "float") {
		return convertFloatCode(g, f, t)
	}
	switch f.typ {
	case "string":
		return fmt.Sprintf("	d.%s = %s", f.name, safeVariableName(t.values[0])), nil
	case "bool":
		g.newImport(iport{path: "strconv"})
		return convertBoolTemplate(t.values[0], f.name), nil
	}
	log.Panicf("Cannot convert type '%s'", f.typ)
	return "", nil
}

// convertIntCode generates the code for reading & converting an int(X) path variable
func convertIntCode(g *generation, f field, t tag) (string, error) {
	g.newImport(iport{path: "strconv"})
	if f.typ == "int" {
		return convertIntTemplate(t.values[0], "Int", 64, f.name, "int"), nil
	}
	bitSize, err := strconv.ParseInt(strings.ReplaceAll(f.typ, "int", ""), 10, 8)
	if err != nil {
		return "", err
	}
	return convertIntTemplate(
		t.values[0], "Int", uint8(bitSize), f.name, fmt.Sprintf("int%d", bitSize)), nil
}

// convertUintCode generates the code for reading & converting an int(X) path variable
func convertUintCode(g *generation, f field, t tag) (string, error) {
	g.newImport(iport{path: "strconv"})
	if f.typ == "uint" {
		return convertIntTemplate(t.values[0], "Uint", 64, f.name, "uint"), nil
	}
	bitSize, err := strconv.ParseInt(strings.ReplaceAll(f.typ, "uint", ""), 10, 8)
	if err != nil {
		return "", err
	}
	return convertIntTemplate(
		t.values[0], "Uint", uint8(bitSize), f.name, fmt.Sprintf("uint%d", bitSize)), nil
}

// convertFloatCode generates the code for reading & converting an float(X) path variable
func convertFloatCode(g *generation, f field, t tag) (string, error) {
	g.newImport(iport{path: "strconv"})
	if f.typ == "float" {
		return convertFloatTemplate(t.values[0], 64, f.name, "float"), nil
	}
	bitSize, err := strconv.ParseInt(strings.ReplaceAll(f.typ, "float", ""), 10, 8)
	if err != nil {
		return "", err
	}
	return convertFloatTemplate(
		t.values[0], uint8(bitSize), f.name, fmt.Sprintf("float%d", bitSize)), nil
}

// pathExtractTemplate is the template code for extracting a path parameter
const pathExtractTemplate = `
	%s := chi.URLParam(r, "%s")
%s
`

// convertIntTemplate is the template code for converting path parameters to (u)int numeric types
func convertIntTemplate(param string, parse string, bits uint8, field string, cast string) string {
	v := safeVariableName(param)
	return fmt.Sprintf(`	%sConvert, err := strconv.Parse%s(%s, 10, %d)
	if err != nil {
		return d, err
	}
	d.%s = %s(%sConvert)`,
		v, parse, v, bits, field, cast, v)
}

// convertFloatTemplate is the template code for converting path parameters to float numeric types
func convertFloatTemplate(param string, bits uint8, field string, cast string) string {
	v := safeVariableName(param)
	return fmt.Sprintf(`	%sConvert, err := strconv.ParseFloat(%s, %d)
	if err != nil {
		return d, err
	}
	d.%s = %s(%sConvert)`,
		v, v, bits, field, cast, v)
}

// convertBoolTemplate is the template code for converting path parameters to float numeric types
func convertBoolTemplate(param string, field string) string {
	v := safeVariableName(param)
	return fmt.Sprintf(`	%sConvert, err := strconv.ParseBool(%s)
	if err != nil {
		return d, err
	}
	d.%s = %sConvert`,
		v, v, field, v)
}
