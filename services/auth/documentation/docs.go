package documentation

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/dtos"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/responses"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/handlers"
)

// UserResponse represents a user
// swagger:response UserResponse
type userResponseWrapper struct {
	// The user
	// in: body
	Body handlers.UserDTO
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
	Body handlers.RegisterUser
}

// swagger:parameters loginUser
type loginUserParamsWrapper struct {

	// Credentials verify the user
	// in: body
	// required: true
	Body handlers.LoginUser
}

// swagger:parameters sessionUser
type sessionUser struct {

	// Credentials verify
	// in: body
	// required: true
	Body dtos.SessionUser
}