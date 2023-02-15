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
}

// queryExtractTemplate is the template code for extracting a path parameter
const queryExtractTemplate = `
	%s := r.URL.Query()["%s"]
	if len(%s) != 1 {
		return d, fmt.Errorf("for query parameter '%s' expected 1 value, got '%s'", %s)
	}
	d.%s = %s[0]
`
