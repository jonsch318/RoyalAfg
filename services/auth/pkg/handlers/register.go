package handlers

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/responses"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/dto"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/services"
	validation "github.com/go-ozzo/ozzo-validation"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

// Register registers a new user
// swagger:route POST /account/register authentication registerUser
//
//	Register a new user
//
// This will register a new user with the provided details
//
//	Consumes:
//	-	application/json
//
//	Produces:
//	-	application/json
//
//	Schemes: http, https
//
// 	Responses:
// 		default: ErrorResponse
//		400: ErrorResponse
//		422: ValidationErrorResponse
//		500: ErrorResponse
//		200: UserResponse
//
func (h *User) Register(rw http.ResponseWriter, r *http.Request) {
	// Set content type header to json
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("X-Content-Type-Options", "nosniff")

	registerDto := &dto.RegisterDto{}
	err := FromJSON(registerDto, r.Body)
	if err != nil {
		h.l.Error("Decoding JSON", "error", err)
		responses.JSONError(rw, &responses.ErrorResponse{Error: "user could not be decoded"}, http.StatusBadRequest)
		return
	}

	user, token, err := h.Auth.Register(registerDto)

	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			h.l.Errorw("Error from grpc conn", "error", err, "status", st)
			switch st.Code() {
			case codes.InvalidArgument:
				h.l.Errorw("Validation", "error", st.Err())
				responses.JSONError(rw, &responses.ValidationError{Errors: st.Message()}, http.StatusUnprocessableEntity)
				return
			case codes.Internal:
				h.l.Errorw("UserService Call Internal", "error", st.Err())
				responses.JSONError(rw, &responses.ErrorResponse{Error: st.Message()}, http.StatusInternalServerError)
				return
			case codes.AlreadyExists:
				h.l.Errorw("Validation", "error", st.Err())
				responses.JSONError(rw, &responses.ValidationError{Errors: st.Message()}, http.StatusUnprocessableEntity)
				return
			}

			responses.JSONError(rw, &responses.ErrorResponse{Error: err.Error()}, http.StatusInternalServerError)
			return
		} else {
			h.l.Errorw("Error during RegisterService", "error", err)
			switch err.(type) {
			case *validation.Errors:
				responses.JSONError(rw, &responses.ValidationError{Errors: err}, http.StatusUnprocessableEntity)
				return
			default:
				responses.JSONError(rw, &responses.ErrorResponse{Error: "Something went wrong"}, http.StatusInternalServerError)
				return
			}
		}
	}

	cookie := services.GenerateCookie(token, registerDto.RememberMe)
	http.SetCookie(rw, cookie)
	err = ToJSON(dto.NewUserDTO(user), rw)
	if err != nil {
		h.l.Errorw("json serialization", "error", err)
		responses.JSONError(rw, &responses.ErrorResponse{Error: "Something went wrong"}, http.StatusInternalServerError)
		return
	}
}
