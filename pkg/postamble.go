package redeco

// commentCode returns comments that should be added to the end of the generated file
func commentCode() string {
	return generatedCodeComment
}

// generatedCodeComment is a standard format string used to indicate that
// the file contains generated code and should not be e.g. linted.
const generatedCodeComment = "// Code generated by github.com/snasphysicist/redeco; DO NOT EDIT."
