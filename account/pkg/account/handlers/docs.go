// Package classification RoyalAfg Auth API
//
// Documentation for RoyalAfg Auth API
//
//	Schemes:http, https
//	BasePath: /
//	Version: 0.0.1
// 	Host: royalafg.games
//	Contact: jonas.max.schneider@gmail.com
//	License: MIT http://opensource.org/license/MIT
//
//	Consumes:
//	- 	application/json
//	-	application/x-www-form-urlencoded
//
//	Produces:
//	- application/json
//
//
// swagger:meta
package handlers

// UserResponse represents a user
// swagger:response UserResponse
type userResponseWrapper struct {
	// The user
	// in: body
	Body UserDTO
}

// ErrorResponse is a generic error response
// swagger:response ErrorResponse
type errorResponseWrapper struct {
	// The error
	// in: body
	Body ErrorResponse
}

// ValidationError shows the failed validation requirements.
// Each form field that has missing requirements is listet under validationErrors
// swagger:response ValidationErrorResponse
type validationErrorWrapper struct {
	// The validation errors
	// in: body
	Body ValidationError
}

// NoContentResponse is an empty object with no content
// swagger:response NoContentResponse
type noContentResponseWrapper struct {
}
