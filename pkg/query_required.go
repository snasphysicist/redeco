package redeco

import (
	"fmt"
	"strconv"
	"strings"
)

// requiredQueryExtractCode generates code to extract and convert
// a query parameter associated with the provided field & tag
func requiredQueryExtractCode(g *generation, f field, t tag) (string, error) {
	if strings.HasPrefix(f.typ, "int") {
		return requiredQueryIntExtractCode(g, f, t)
	}
	if strings.HasPrefix(f.typ, "uint") {
		return requiredQueryUintExtractCode(g, f, t)
	}
	switch f.typ {
	case "string":
		v := safeVariableName(t.values[0])
		return fmt.Sprintf(
			requiredQueryStringExtractTemplate,
			v,
			t.values[0],
			v,
			t.values[0],
			"%v",
			v,
			f.name,
			v,
		), nil
	case "bool":
		attachConversionImports(g)
		return requiredQueryBoolExtractTemplate(t.values[0], f.name), nil
	}
	return "", fmt.Errorf("don't know how to convert type '%s'", f.typ)
}

// requiredQueryIntExtractCode generates the code for reading & converting an int(X) required query parameter
func requiredQueryIntExtractCode(g *generation, f field, t tag) (string, error) {
	attachConversionImports(g)
	if f.typ == "int" {
		return requiredQueryIntExtractTemplate(t.values[0], f.name, "Int", 64, "int"), nil
	}
	bitSize, err := strconv.ParseInt(strings.ReplaceAll(f.typ, "int", ""), 10, 8)
	if err != nil {
		return "", err
	}
	return requiredQueryIntExtractTemplate(
		t.values[0], f.name, "Int", uint8(bitSize), fmt.Sprintf("int%d", bitSize)), nil
}

// requiredQueryUintExtractCode generates the code for reading & converting a uint(X) required query parameter
func requiredQueryUintExtractCode(g *generation, f field, t tag) (string, error) {
	attachConversionImports(g)
	if f.typ == "uint" {
		return requiredQueryIntExtractTemplate(t.values[0], f.name, "Uint", 64, "uint"), nil
	}
	bitSize, err := strconv.ParseInt(strings.ReplaceAll(f.typ, "uint", ""), 10, 8)
	if err != nil {
		return "", err
	}
	return requiredQueryIntExtractTemplate(
		t.values[0], f.name, "Uint", uint8(bitSize), fmt.Sprintf("uint%d", bitSize)), nil
}

// requiredQueryStringExtractTemplate is the template code for extracting a string path parameter
const requiredQueryStringExtractTemplate = `
	%s := r.URL.Query()["%s"]
	if len(%s) != 1 {
		return d, fmt.Errorf("for query parameter '%s' expected 1 value, got '%s'", %s)
	}
	d.%s = %s[0]
`

// requiredQueryIntExtractTemplate generates code extracting & converting
// a query parameter with a (u)int* type
func requiredQueryIntExtractTemplate(param string, field string, parse string, bits uint8, typ string) string {
	v := safeVariableName(param)
	return fmt.Sprintf(`
	%s := r.URL.Query()["%s"]
	if len(%s) != 1 {
		return d, fmt.Errorf("for query parameter '%s' expected 1 value, got '%s'", %s)
	}
	%sConvert, err := strconv.Parse%s(%s[0], 10, %d)
	if err != nil {
		return d, err
	}
	d.%s = %s(%sConvert)
`, v, param, v, param, "%v", v, v, parse, v, bits, field, typ, v)
}

// requiredQueryBoolExtractTemplate generates code extracting & converting
// a query parameter with a bool type
func requiredQueryBoolExtractTemplate(param string, field string) string {
	v := safeVariableName(param)
	return fmt.Sprintf(`
	%s := r.URL.Query()["%s"]
	if len(%s) != 1 {
		return d, fmt.Errorf("for query parameter '%s' expected 1 value, got '%s'", %s)
	}
	%sConvert, err := strconv.ParseBool(%s[0])
	if err != nil {
		return d, err
	}
	d.%s = %sConvert
`, v, param, v, param, "%v", v, v, v, field, v)
}
