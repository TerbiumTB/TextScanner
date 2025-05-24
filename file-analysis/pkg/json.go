package pkg

import (
	"encoding/json"
	"io"
)

func ToJSON(i any, w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(i)
}


func FromJSON(i any, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}
