package redeco

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// fileWithNameAndPackage searches for a file with the given name
// including extension which is in the given package
func fileWithNameAndPackage(name string, pkg string) (string, error) {
	withName := filter(goFiles(listFiles()), func(s string) bool { return strings.HasSuffix(s, name) })
	for _, path := range withName {
		i, err := inPackage(path, pkg)
		if err != nil {
			return "", err
		}
		if i {
			return path, nil
		}
	}
	return "", fmt.Errorf("failed to find file '%s' in package '%s'", name, pkg)
}

// goFiles filters out no go file paths from the input
func goFiles(files []string) []string {
	return filter(files, func(s string) bool { return strings.HasSuffix(s, ".go") })
}

// listFiles lists all files under the current directory
func listFiles() []string {
	pc := &pathCollector{}
	root, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}
	err = filepath.Walk(root, pc.gather)
	if err != nil {
		panic(err)
	}
	return pc.paths
}

// pathCollector implements an os.WalkFunc to gather all walked file's paths
type pathCollector struct {
	paths []string
}

// gather implements os.WalkFunc for pathCollector
func (p *pathCollector) gather(path string, _ os.FileInfo, _ error) error {
	p.paths = append(p.paths, path)
	return nil
}

// inPackage returns true if the go source file at the path is in this package
func inPackage(path string, targetPkg string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()
	src, err := io.ReadAll(f)
	if err != nil {
		return false, err
	}
	pkg, err := findPackage(string(src))
	if err != nil {
		return false, err
	}
	return pkg == targetPkg, nil
}
