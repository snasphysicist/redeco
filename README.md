# redeco

__REquest DECOder__

Generate boilerplate for decoding HTTP requests
alongside your code instead of writing it by hand.

## Goals

### Core

To generate Go code that can automatically
decode the most common inputs to a HTTP request

- JSON
- path parameters
- query parameters

and the most common types of those inputs e.g.

- strings
- all Go `int*` primitives
- all Go `uint*` primitives
- all Go `float*` primitives
- bool

for client code using common request routers e.g. `chi`.

### Secondary

- Understandable generated source

## Non-Goals

- Highest possible performance
- Other body formats (e.g. XML)
- Deserialisation of every possible (custom) type

## Usage

Install.

```shell
go install github.com/snasphysicist/redeco@latest
```

Create a struct to hold all information from the request, in the same file as the handler.

```go
package endpoint

type target struct {
    body  string `json:"fromBody"`
    path  int16  `path:"fromPath"`
    query bool   `query:"fromQuery"`
}

type post(w http.ResponseWriter, r *http.Request) {
    // ...
}
```

Include a `go:generate` comment in the same file. The first argument to the command
is the handler name, the second is the target struct name.

```go
//go:generate redeco post target
```

Run `go generate`.

A new file will be generated in the same
package as the file containing the handler
with the generated decoding function.

```go
package endpoint

func postDecoder(r *http.Request) (target, error) {
    // generated decoding code
}
```
Finally, use the decoder in the the handler function.

```go
package endpoint

// ...

func post(w http.ResponseWriter, r *http.Request) {
    t, err := postDecoder(r)
    if err != nil {
        // ...
    }
    // ...
}
```