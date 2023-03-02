package redeco

import (
	"log"
)

// bodyDeserialiseCode generates the code for deserialising
// the request body into a struct
func bodyDeserialiseCode(g *generation) (string, error) {
	s, err := targetStruct(g)
	if err != nil {
		return "", err
	}
	if !hasAJSONTag(s) {
		log.Print("Target struct has no JSON tags, won't generate JSON deserialisation code")
		return "", nil
	}
	g.newImport(iport{path: "io"})
	g.newImport(iport{path: "encoding/json"})
	return bodyDeserialiseTemplate, nil
}

// hasAJSONTag returns true if the struct s has any fields with any JSON tags
func hasAJSONTag(s sourceStruct) bool {
	return anyMatch(s.field, func(f field) bool {
		return anyMatch(f.tags, func(t tag) bool { return t.key == "json" })
	})
}

// bodyDeserialiseTemplate is a template for the code
// to JSON deserialise a request body into a struct
const bodyDeserialiseTemplate = `
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return d, err
	}
	err = json.Unmarshal(b, &d)
	if err != nil {
		return d, err
	}
`
