package test

import "net/http"

//go:generate redeco Post PostParameters

func Post(w http.ResponseWriter, r *http.Request) {

}

type PostParameters struct {
	A string `json:"a"`
	B int64  `path:"b"`
	C uint16 `query:"c"`
	D bool   `query:"d,optional"`
}
