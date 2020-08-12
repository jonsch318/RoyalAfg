package handlers

import "github.com/JohnnyS318/RoyalAfgInGo/shared/pkg/responses"

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
	Body responses.ErrorResponse
}

// ValidationError shows the failed validation requirements.
// Each form field that has missing requirements is listet under validationErrors
// swagger:response ValidationErrorResponse
type validationErrorWrapper struct {
	// The validation errors
	// in: body
	Body responses.ValidationError
}

// NoContentResponse is an empty object with no content
// swagger:response NoContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters registerUser
type registerUserParamsWrapper struct {

	// User to register and save
	// in: body
	// required: true
	Body RegisterUser
}

// swagger:parameters loginUser
type loginUserParamsWrapper struct {

	// Credentials verify the user
	// in: body
	// required: true
	Body LoginUser
}
