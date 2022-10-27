package utils

import (
	"net/http"
)

func RespondWithError(rw http.ResponseWriter, code int, message string) error {
	rw.WriteHeader(code)
	_, err := rw.Write([]byte(message))
	return err
}
