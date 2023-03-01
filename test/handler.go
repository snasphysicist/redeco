package test

import (
	"encoding/json"
	"net/http"
)

//go:generate redeco Post PostParameters

// Post is a handler for which to generate a decoder for testing
func Post(w http.ResponseWriter, r *http.Request) {
	p, err := PostDecoder(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		mustWrite(w, []byte(err.Error()))
		return
	}
	bs, err := json.Marshal(response(p))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		mustWrite(w, []byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	mustWrite(w, bs)
}

// PostParameters is the target struct for generating a decoder
type PostParameters struct {
	A string `json:"a"`
	B int64  `path:"b"`
	C uint16 `query:"c"`
	D bool   `query:"d,optional"`
}

// mustWrite writes the given bytes to the writer and panics on error
func mustWrite(w http.ResponseWriter, bs []byte) {
	_, err := w.Write(bs)
	if err != nil {
		panic(err)
	}
}

// response allows the information decoded to be
// sent back in the response as JSON for checking
type response struct {
	A string `json:"a"`
	B int64  `json:"b"`
	C uint16 `json:"c"`
	D bool   `json:"d"`
}
