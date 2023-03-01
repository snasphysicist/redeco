package redeco

import "fmt"

// safeVariableName renames variables that are known
// to be used in template code to avoid clashes
func safeVariableName(v string) string {
	switch v {
	case "b", "d":
		return fmt.Sprintf("%s_", v)
	}
	return v
}
