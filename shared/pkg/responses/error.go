package responses

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse is a generic response containing an error
type ErrorResponse struct {
	// The error
	Error string `json:"error"`
}

// ValidationError shows the failed validation requirements.
// Each form property that has missing requirements is listet under Errors (validationErrors)
type ValidationError struct {
	// The missing requirements
	Errors interface{} `json:"validationErrors"`
}

func (myErr *ValidationError) Error() string {
	b, err := json.Marshal(myErr.Errors)

	if err != nil {
		return "Could not encode validation errors"
	}
	return string(b)
}

// JSONError writes the given object as an error with the given code to the response writer
func JSONError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}
