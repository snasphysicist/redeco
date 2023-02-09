package redeco

import (
	"fmt"
	"go/ast"
	"log"
)

// targetStruct locates the target struct in the source
func targetStruct(g *generation) (sourceStruct, error) {
	ss, err := findStructs(g.src)
	if err != nil {
		return sourceStruct{}, err
	}
	log.Printf("Found structs in file: %#v", ss)
	ts := filter(ss, func(s sourceStruct) bool { return s.name == g.o.target })
	if len(ts) != 1 {
		return sourceStruct{}, fmt.Errorf("no/too many struct/s named '%s': %#v", g.o.target, ts)
	}
	log.Printf("Found struct %#v with name '%s'", ts[0], g.o.target)
	return ts[0], err
}

// findStructs finds all structs in the passed source code
func findStructs(s string) ([]sourceStruct, error) {
	f, err := parse(s)
	if err != nil {
		return make([]sourceStruct, 0), err
	}
	sf := newStructFinder()
	ast.Walk(sf, f)
	if len(filter(*sf.e, func(e error) bool { return e != nil })) > 0 {
		return make([]sourceStruct, 0), (*sf.e)[0]
	}
	return *sf.s, nil
}

// sourceStruct is a struct found in Go source code
type sourceStruct struct {
	name  string
	field []field
}

// structFinder is an ast.Visitor which finds all structs under the node
type structFinder struct {
	s *[]sourceStruct
	e *[]error
}

// newStructFinder creates a new structFinder, initialising as required
func newStructFinder() structFinder {
	return structFinder{
		s: asPointer(make([]sourceStruct, 0)),
		e: asPointer(make([]error, 0)),
	}
}

// Visit implements ast.Visitor on structFinder
func (s structFinder) Visit(n ast.Node) ast.Visitor {
	if ts, ok := n.(*ast.TypeSpec); ok {
		st, err := fromTypeSpec(ts)
		*s.s = append(*s.s, st...)
		*s.e = append(*s.e, err)
	}
	return s
}

// fromTypeSpec converts the TypeSpec into a sourceStruct
// if it is a a struct type and not anonymous
func fromTypeSpec(ts *ast.TypeSpec) ([]sourceStruct, error) {
	st, ok := ts.Type.(*ast.StructType)
	if !ok {
		log.Printf("Ignoring %#v, not a struct", ts)
		return make([]sourceStruct, 0), nil
	}
	if ts.Name == nil {
		log.Printf("Ignoring %#v, no name", ts)
		return make([]sourceStruct, 0), nil
	}
	fs, err := fromFieldList(st.Fields)
	if err != nil {
		return make([]sourceStruct, 0), err
	}
	return []sourceStruct{{name: ts.Name.Name, field: fs}}, nil
}

// asPointer returns a pointer to the provided item
func asPointer[T any](a T) *T {
	return &a
}
