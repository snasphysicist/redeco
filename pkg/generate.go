package redeco

import (
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
)

// Generate is the main command called when used as a CLI tool
func Generate() {
	g := NewFromString("")
	o, err := parseArguments()
	if err != nil {
		log.Fatalf("Input arguments incorrect: %s", err)
	}
	c, err := g.Generate(o)
	log.Printf("Output: %s", c)
	log.Printf("Error: %s", err)
}

// generator is responsible for generating the new source code
type generator struct {
	source io.Reader
}

// generation contains/tracks information which all parts of the process need
type generation struct {
	o       options
	src     string
	imports []iport
}

// iport represents a possibly aliased import
type iport struct {
	alias string
	path  string
}

// String implements Stringer for iport, prints as a source code import
func (i iport) String() string {
	alias := ""
	if i.alias != "" {
		alias = fmt.Sprintf("%s ", i.alias)
	}
	return fmt.Sprintf(`import %s"%s"`, alias, i.path)
}

// newImport adds the import to the list, if it's not already there
func (g *generation) newImport(i iport) {
	g.imports = unique(append(g.imports, i))
}

// NewFromString creates a new generator, reading source code
// on which it was invoked from a string
func NewFromString(s string) generator {
	return generator{
		source: strings.NewReader(s),
	}
}

// Generate generates the new source code
func (g generator) Generate(o options) (string, error) {
	if o.target == "" {
		return "", errors.New("no deserialisation target was given")
	}
	b, err := io.ReadAll(g.source)
	if err != nil {
		return "", fmt.Errorf("failed to read input: %s", err)
	}
	gtn := generation{o: o, src: string(b)}
	// TODO: replace with proper logger
	log.Printf("Called with input: %s", string(b))
	p, err := packageCode(&gtn)
	if err != nil {
		return "", err
	}
	log.Printf("Called on package: %s", p)
	f, err := generateFunction(&gtn)
	if err != nil {
		return "", err
	}
	log.Printf("Generated decoding function: %s", f)
	i := importsCode(&gtn)
	return fmt.Sprintf("%s\n%s\n%s", p, i, f), nil
}
