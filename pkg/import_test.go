package redeco

import "testing"

func TestCanParseSourceContainingStructsWithImportedTypes(t *testing.T) {
	g := NewFromString(`
	package foo

	import "github.com/google/uuid"

	type A struct {}

	type B struct {
		U uuid.UUID
	}
	`)
	_, err := g.Generate(options{handler: "bar", target: "A"})
	if err != nil {
		t.Errorf("Generation failed with %s", err)
	}
}
