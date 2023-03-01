package test

import "encoding/json"
import "fmt"
import chi "github.com/go-chi/chi/v5"
import "io"
import "net/http"
import "strconv"

func PostDecoder(r *http.Request) (PostParameters, error) {
	var d PostParameters
	var err error

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return d, err
	}
	err = json.Unmarshal(b, &d)
	if err != nil {
		return d, err
	}

	b_ := chi.URLParam(r, "b")
	b_Convert, err := strconv.ParseInt(b_, 10, 64)
	if err != nil {
		return d, err
	}
	d.B = int64(b_Convert)

	c := r.URL.Query()["c"]
	if len(c) != 1 {
		return d, fmt.Errorf("for query parameter 'c' expected 1 value, got '%v'", c)
	}
	cConvert, err := strconv.ParseUint(c[0], 10, 16)
	if err != nil {
		return d, err
	}
	d.C = uint16(cConvert)


	d_ := r.URL.Query()["d"]
	if len(d_) > 1 {
		return d, fmt.Errorf("for query parameter 'd' expected 0 or 1 value, got '%v'", d_)
	}
	if len(d_) == 1 {
		d_Convert, err := strconv.ParseBool(d_[0])
		if err != nil {
			return d, err
		}
		d.D = d_Convert
	}

	return d, err
}
