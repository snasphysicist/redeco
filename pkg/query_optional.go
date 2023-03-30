package redeco

import (
	"fmt"
	"strconv"
	"strings"
)

// optionalQueryExtractCode generates code to extract and convert
// a query parameter associated with the provided field & tag
func optionalQueryExtractCode(g *generation, f field, t tag) (string, error) {
	if strings.HasPrefix(f.typ, "int") {
		return optionalQueryIntExtractCode(g, f, t)
	}
	if strings.HasPrefix(f.typ, "uint") {
		return optionalQueryUintExtractCode(g, f, t)
	}
	switch f.typ {
	case "string":
		return optionalQueryStringExtractTemplate(t.values[0], f.name), nil
	case "bool":
		attachConversionImports(g)
		return optionalQueryBoolExtractTemplate(t.values[0], f.name), nil
	}
	return "", fmt.Errorf("don't know how to handle type '%s'", f.typ)
}

// optionalQueryIntExtractCode generates the code for reading & converting an int(X) optional query parameter
func optionalQueryIntExtractCode(g *generation, f field, t tag) (string, error) {
	attachConversionImports(g)
	if f.typ == "int" {
		return optionalQueryIntExtractTemplate(t.values[0], f.name, "Int", 64, "int"), nil
	}
	bitSize, err := strconv.ParseInt(strings.ReplaceAll(f.typ, "int", ""), 10, 8)
	if err != nil {
		return "", err
	}
	return optionalQueryIntExtractTemplate(
		t.values[0], f.name, "Int", uint8(bitSize), fmt.Sprintf("int%d", bitSize)), nil
}

// optionalQueryUintExtractCode generates the code for reading & converting a uint(X) optional query parameter
func optionalQueryUintExtractCode(g *generation, f field, t tag) (string, error) {
	attachConversionImports(g)
	if f.typ == "uint" {
		return optionalQueryIntExtractTemplate(t.values[0], f.name, "Uint", 64, "uint"), nil
	}
	bitSize, err := strconv.ParseInt(strings.ReplaceAll(f.typ, "uint", ""), 10, 8)
	if err != nil {
		return "", err
	}
	return optionalQueryIntExtractTemplate(
		t.values[0], f.name, "Uint", uint8(bitSize), fmt.Sprintf("uint%d", bitSize)), nil
}

// optionalQueryStringExtractTemplate generates code extracting an optional string type query parameter
func optionalQueryStringExtractTemplate(param string, field string) string {
	v := safeVariableName(param)
	return fmt.Sprintf(`
	%s := r.URL.Query()["%s"]
	if len(%s) > 1 {
		return d, fmt.Errorf("for query parameter '%s' expected 0 or 1 value, got '%s'", %s)
	}
	if len(%s) == 1 {
		d.%s = %s[0]
	}
`, v, param, v, param, "%v", v, v, field, v)
}

// optionalQueryBoolExtractTemplate generates code extracting & converting
// a query parameter with a bool type
func optionalQueryBoolExtractTemplate(param string, field string) string {
	v := safeVariableName(param)
	return fmt.Sprintf(`
	%s := r.URL.Query()["%s"]
	if len(%s) > 1 {
		return d, fmt.Errorf("for query parameter '%s' expected 0 or 1 value, got '%s'", %s)
	}
	if len(%s) == 1 {
		%sConvert, err := strconv.ParseBool(%s[0])
		if err != nil {
			return d, err
		}
		d.%s = %sConvert
	}
`, v, param, v, param, "%v", v, v, v, v, field, v)
}

// optionalQueryIntExtractTemplate generates code extracting & converting
// a query parameter with a (u)int type
func optionalQueryIntExtractTemplate(param string, field string, fn string, bits uint8, cast string) string {
	v := safeVariableName(param)
	return fmt.Sprintf(`
	%s := r.URL.Query()["%s"]
	if len(%s) > 1 {
		return d, fmt.Errorf("for query parameter '%s' expected 0 or 1 value, got '%s'", %s)
	}
	if len(%s) == 1 {
		%sConvert, err := strconv.Parse%s(%s[0], 10, %d)
		if err != nil {
			return d, err
		}
		d.%s = %s(%sConvert)
	}
`, v, param, v, param, "%v", v, v, v, fn, v, bits, field, cast, v)
}
