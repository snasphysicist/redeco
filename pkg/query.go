package redeco

import (
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
		log.Panicf("Could not find unique query tag in %#v", f)
	}
	if parameterIsOptional(t[0]) {
		return optionalQueryExtractCode(g, f, t[0])
	}
	return requiredQueryExtractCode(g, f, t[0])
}

// parameterIsOptional returns true iff the tag values
// indicate that the query parameter may be omitted
func parameterIsOptional(t tag) bool {
	return anyMatch(t.values, func(s string) bool { return s == "optional" })
}

// attachConversionImports attaches to the generation
func attachConversionImports(g *generation) {
	g.newImport(iport{path: "fmt"})
	g.newImport(iport{path: "strconv"})
}
