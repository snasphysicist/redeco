package redeco

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// findPackage finds the package name from the source s
func findPackage(s string) (string, error) {
	f, err := parse(s)
	if err != nil {
		return "", err
	}
	return f.Name.Name, nil
}

// parse parses the go source code in s
func parse(s string) (*ast.File, error) {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "", s, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to parse file: %s", err)
	}
	return f, nil
}

// packageCode constructs the code which declares the package
func packageCode(g *generation) (string, error) {
	p, err := findPackage(g.src)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("package %s\n", p), nil
}
