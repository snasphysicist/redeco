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
	return fmt.Sprintf(
		pathExtractTemplate,
		t[0].values[0],
		t[0].values[0],
		convertCode(g, f, t[0]),
	)
}

// convertCode generates the code to convert the path parameter to the correct type
func convertCode(g *generation, f field, t tag) string {
	switch f.typ {
	case "string":
		return fmt.Sprintf("	d.%s = %s", f.name, t.values[0])
	case "int":
		g.newImport(iport{path: "strconv"})
		return fmt.Sprintf(convertTemplate,
			t.values[0], "Int", t.values[0], f.name, "int(", t.values[0], ")")
	case "int32":
		g.newImport(iport{path: "strconv"})
		return fmt.Sprintf(convertTemplate,
			t.values[0], "Int", t.values[0], f.name, "int32(", t.values[0], ")")
	case "int64":
		g.newImport(iport{path: "strconv"})
		return fmt.Sprintf(convertTemplate,
			t.values[0], "Int", t.values[0], f.name, "int64(", t.values[0], ")")
	case "int16":
		g.newImport(iport{path: "strconv"})
		return fmt.Sprintf(convertTemplate,
			t.values[0], "Int", t.values[0], f.name, "int16(", t.values[0], ")")
	case "int8":
		g.newImport(iport{path: "strconv"})
		return fmt.Sprintf(convertTemplate,
			t.values[0], "Int", t.values[0], f.name, "int8(", t.values[0], ")")
	}
	log.Panicf("Cannot convert type '%s'", f.typ)
	return ""
}

// pathExtractTemplate is the template code for extracting a path parameter
const pathExtractTemplate = `
	%s := chi.URLParam(r, "%s")
%s
`

// convertTemplate is the template code for converting path parameters to numeric types
const convertTemplate = `	%sConvert, err := strconv.Parse%s(%s, 10, 64)
	if err != nil {
		return d, err
	}
	d.%s = %s%sConvert%s`
