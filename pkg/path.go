package redeco

import (
	"fmt"
	"log"
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
	c := mapping(f, func(f field) string { return pathExtractCode(g, f) })
	return strings.Join(c, "\n"), nil
}

// pathTaggedFields returns any fields in the struct with path tags
func pathTaggedFields(s sourceStruct) []field {
	return filter(s.field, func(f field) bool {
		return anyMatch(f.tags, func(t tag) bool { return t.key == "path" })
	})
}

// pathExtractCode generates code to extract the path parameter associated with f
func pathExtractCode(g *generation, f field) string {
	t := filter(f.tags, func(t tag) bool { return t.key == "path" })
	if len(t) != 1 {
		log.Panicf("Could not find unique path tag in %#v", f)
	}
	v := safeVariableName(t[0].values[0])
	return fmt.Sprintf(
		pathExtractTemplate,
		v,
		t[0].values[0],
		convertCode(g, f, t[0]),
	)
}

// convertCode generates the code to convert the path parameter to the correct type
func convertCode(g *generation, f field, t tag) string {
	switch f.typ {
	case "string":
		return fmt.Sprintf("	d.%s = %s", f.name, safeVariableName(t.values[0]))
	case "int":
		g.newImport(iport{path: "strconv"})
		return convertIntTemplate(t.values[0], "Int", 64, f.name, "int")
	case "int32":
		g.newImport(iport{path: "strconv"})
		return convertIntTemplate(t.values[0], "Int", 32, f.name, "int32")
	case "int64":
		g.newImport(iport{path: "strconv"})
		return convertIntTemplate(t.values[0], "Int", 64, f.name, "int64")
	case "int16":
		g.newImport(iport{path: "strconv"})
		return convertIntTemplate(t.values[0], "Int", 16, f.name, "int16")
	case "int8":
		g.newImport(iport{path: "strconv"})
		return convertIntTemplate(t.values[0], "Int", 8, f.name, "int8")
	case "uint":
		g.newImport(iport{path: "strconv"})
		return convertIntTemplate(t.values[0], "Uint", 64, f.name, "uint")
	case "uint64":
		g.newImport(iport{path: "strconv"})
		return convertIntTemplate(t.values[0], "Uint", 64, f.name, "uint64")
	case "uint32":
		g.newImport(iport{path: "strconv"})
		return convertIntTemplate(t.values[0], "Uint", 32, f.name, "uint32")
	case "uint16":
		g.newImport(iport{path: "strconv"})
		return convertIntTemplate(t.values[0], "Uint", 16, f.name, "uint16")
	case "uint8":
		g.newImport(iport{path: "strconv"})
		return convertIntTemplate(t.values[0], "Uint", 8, f.name, "uint8")
	case "float":
		g.newImport(iport{path: "strconv"})
		return convertFloatTemplate(t.values[0], 64, f.name, "float")
	case "float64":
		g.newImport(iport{path: "strconv"})
		return convertFloatTemplate(t.values[0], 64, f.name, "float64")
	case "float32":
		g.newImport(iport{path: "strconv"})
		return convertFloatTemplate(t.values[0], 32, f.name, "float32")
	case "bool":
		g.newImport(iport{path: "strconv"})
		return convertBoolTemplate(t.values[0], f.name)
	}
	log.Panicf("Cannot convert type '%s'", f.typ)
	return ""
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
