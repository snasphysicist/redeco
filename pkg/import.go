package redeco

import (
	"sort"
	"strings"
)

// importsCode generates the section of the code which imports packages
func importsCode(g *generation) string {
	sort.Sort(byPath(g.imports))
	iss := mapping(g.imports, iport.String)
	return strings.Join(iss, "\n")
}

// byPath wraps []iport to implement sort.Interface
// sorting alphanumerically by iport.path
type byPath []iport

// Len implements sort.Interface on byPath
func (b byPath) Len() int {
	return len(b)
}

// Less implements sort.Interface on byPath
func (b byPath) Less(i int, j int) bool {
	return b[i].path < b[j].path
}

// Swap implements sort.Interface on byPath
func (b byPath) Swap(i int, j int) {
	b[i], b[j] = b[j], b[i]
}
