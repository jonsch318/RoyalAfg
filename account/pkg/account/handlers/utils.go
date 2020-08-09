package handlers

import (
	"encoding/json"
	"io"
)

// ToJSON encodes objects and writes them using the specified io.Writer
func ToJSON(i interface{}, w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(i)
}

// FromJSON decodes objects from an io.Reader
func FromJSON(i interface{}, r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(i)
}
