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
	c := mapping(f, func(f field) string { return pathExtractCode(s, f) })
	return strings.Join(c, "\n"), nil
}

// pathTaggedFields returns any fields in the struct with path tags
func pathTaggedFields(s sourceStruct) []field {
	return filter(s.field, func(f field) bool {
		return anyMatch(f.tags, func(t tag) bool { return t.key == "path" })
	})
}

// pathExtractCode generates code to extract the path parameter associated with f
func pathExtractCode(s sourceStruct, f field) string {
	t := filter(f.tags, func(t tag) bool { return t.key == "path" })
	if len(t) != 1 {
		log.Panicf("Could not find unique path tag in %#v", f)
	}
	return fmt.Sprintf(
		pathExtractTemplate,
		t[0].values[0],
		t[0].values[0],
		f.name,
		t[0].values[0],
	)
}

// pathExtractTemplate is the template code for extracting a path parameter
const pathExtractTemplate = `
	%s := chi.URLParam(r, "%s")
	d.%s = %s
`
