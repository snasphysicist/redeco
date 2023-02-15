package redeco

import "fmt"

// generateFunction generates the go source code for the decoding function
func generateFunction(g *generation) (string, error) {
	b, err := bodyDeserialiseCode(g)
	if err != nil {
		return "", err
	}
	p, err := pathDeserialiseCode(g)
	if err != nil {
		return "", err
	}
	q, err := queryDeserialiseCode(g)
	if err != nil {
		return "", err
	}
	g.newImport(iport{path: "net/http"})
	return fmt.Sprintf(
		functionTemplate,
		g.o.handler,
		g.o.target,
		g.o.target,
		b,
		p,
		q,
	), nil
}

// functionTemplate is the wrapper around the decoding function
const functionTemplate = `
func %sDecoder(r *http.Request) (%s, error) {
	var d %s
	var err error
%s%s%s
	return d, err
}
`
