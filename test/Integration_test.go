package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestExampleHandlerDeserialisesAsExpected(t *testing.T) {
	shutdown := StartServer()
	defer shutdown()
	res := makeRequestToTestHandler()
	expectResponse(t, res, response{A: a, B: b, C: c, D: false})
}

// expectResponse fails the test if the content of actual's body
// does not match the expected response
func expectResponse(t *testing.T, actual *http.Response, expected response) {
	bs, err := io.ReadAll(actual.Body)
	if err != nil {
		t.Errorf("Failed to read response body: %s", err)
	}
	if actual.StatusCode != http.StatusOK {
		t.Errorf("Response indicated failure: %d", actual.StatusCode)
	}
	var actualRes response
	err = json.Unmarshal(bs, &actualRes)
	if err != nil {
		t.Errorf("Failed to deserialise response body: %s", err)
	}
	if !reflect.DeepEqual(expected, actualRes) {
		t.Errorf("Actual response %#v != expected response %#v", actualRes, expected)
	}
}

// makeRequestToTestHandler builds & makes the test request to the test server
// and returns the response, panicing on any error
func makeRequestToTestHandler() *http.Response {
	bs, err := json.Marshal(struct {
		A string `json:"a"`
	}{A: a})
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("http://localhost:%d/records/%d?c=%d", port, b, c),
		bytes.NewReader(bs),
	)
	if err != nil {
		panic(err)
	}
	res, err := (&http.Client{}).Do(req)
	if err != nil {
		panic(err)
	}
	return res
}

// a is the value for parameter a, a string value in the request JSON body
const a = "body-string"

// b is the value for parameter b, a path parameter
const b = int64(-14848)

// c is the value for parameter c, a required query parameter
const c = uint16(16)
