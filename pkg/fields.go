package redeco

import (
	"fmt"
	"go/ast"
	"log"
	"strings"
)

// fromFieldList parses fields and tags
// from Go source code into more useful structs
func fromFieldList(fl *ast.FieldList) ([]field, error) {
	return mapWithError(fl.List, fromField)
}

// fromField parses a field and its tags from
// Go source code into more useful structs
func fromField(f *ast.Field) (field, error) {
	n, err := nameOf(f)
	if err != nil {
		return field{}, err
	}
	typ := typeOf(f)
	t := tagsAttachedTo(f)
	return field{name: n, typ: typ, tags: t}, nil
}

// nameOf finds the name of a field in Go source code
func nameOf(f *ast.Field) (string, error) {
	if f.Names == nil || len(f.Names) == 0 {
		return "", fmt.Errorf("field %#v has no name", f)
	}
	if len(f.Names) > 1 {
		return "", fmt.Errorf("field %#v has too many names", f)
	}
	return f.Names[0].Name, nil
}

// typeOf returns the type name associated with the field
func typeOf(f *ast.Field) string {
	switch tt := f.Type.(type) {
	case *ast.Ident:
		return tt.Name
	}
	log.Panicf("Cannot deal with type %#v", f.Type)
	return ""
}

// tagsAttachedTo parses all tags attached
// to a field in Go source code
func tagsAttachedTo(f *ast.Field) []tag {
	tags := strings.Trim(f.Tag.Value, "`")
	remaining := tags
	parsed := make([]tag, 0)
	for len(remaining) != 0 {
		r := nextTagKVPair(remaining)
		if r.err != nil {
			panic(r.err)
		}
		parsed = append(parsed, tag{key: r.key, values: []string{r.value}})
		remaining = r.remaining
	}
	return parsed
}

// nextPairResult is the result of an attempt
// to extract the next key-value pair from a tag
type nextPairResult struct {
	remaining string
	key       string
	value     string
	err       error
}

// nextTagKVPair attempts to extract the next
// key-value pair from a tag
func nextTagKVPair(remaining string) nextPairResult {
	regex := `^[,]{0,1}(?P<key>[A-Za-z0-9]+):"(?P<value>[^"]+)".*`
	kv := groups(regex, remaining)
	key, keyOK := kv["key"]
	value, valueOK := kv["value"]
	if !keyOK || !valueOK {
		return nextPairResult{err: fmt.Errorf("next struct tag in '%s' is malformed, could not extract key and value", remaining)}
	}
	countNonKVCharacters := `:""`
	if strings.HasPrefix(remaining, ",") {
		countNonKVCharacters = "," + countNonKVCharacters
	}
	return nextPairResult{
		remaining: remaining[len(key)+len(value)+len(countNonKVCharacters):],
		key:       key,
		value:     value,
		err:       nil,
	}
}

// field is a representation of a field
// in a struct in Go source code
type field struct {
	name string
	typ  string
	tags []tag
}

// tag represents a key-value pair
// from a tag attached to a Go struct field
type tag struct {
	key    string
	values []string
}
