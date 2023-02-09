package redeco

import (
	"fmt"
	"log"
	"os"
)

// parseArguments reads the provided arguments into generation options
func parseArguments() (options, error) {
	arguments := os.Args[1:]
	if len(arguments) != 2 {
		return options{}, fmt.Errorf(
			"must be called with two arguments, called with %d", len(arguments))
	}
	o := options{handler: arguments[0], target: arguments[1]}
	log.Printf("Loaded options %#v", o)
	return o, nil
}

// options are the options for code generation
type options struct {
	handler string // name of the handler function we are generating code for
	target  string // name of the struct the request should be deserialised into
}
