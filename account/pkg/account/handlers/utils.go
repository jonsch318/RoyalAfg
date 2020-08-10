package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/schema"
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

// FromFormURLEncoded decodes objects form a parsed form
func FromFormURLEncoded(i interface{}, r map[string][]string) error {
	decoder := schema.NewDecoder()
	return decoder.Decode(i, r)
}

// FromFormURLEncodedRequest decodes objects from a form url encoded request
func FromFormURLEncodedRequest(i interface{}, r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}
	return FromFormURLEncoded(i, r.Form)
}
