package handlers

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/auth"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/dtos"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/responses"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/utils"
)

// UpdateUserResponse is the response of a update user command.
type UpdateUserResponse struct {
	Token string     `json:"session"`
	User  *dtos.User `json:"user"`
}

func (h *UserHandler) UpdateUser(rw http.ResponseWriter, r *http.Request) {
	//state whether the session token has to be regenerated with the new information
	reGenToken := false
	claims := mw.FromUserTokenContext(r.Context().Value("user"))

	//Validate CSRF
	if err := mw.ValidateCSRF(r); err != nil {
		h.l.Errorw("could not validate csrf token", "error", err)
		responses.Error(rw, "wrong format decoding failed", http.StatusForbidden)
		return
	}

	//Decode DTO
	dto := new(dtos.User)
	err := utils.FromJSON(dto, r.Body)
	if err != nil {
		h.l.Errorw("DTO deserialization error", "error", err)
		responses.JSONError(rw, &responses.ErrorResponse{Error: "wrong format. decoding failed."}, http.StatusBadRequest)
		return
	}

	//We dont let the user replace all the information. Only certain (email, filename) fields are changeable.
	user, err := h.db.FindById(claims.ID)
	if err != nil {
		h.l.Errorw("User query", "error", err)
		responses.Error(rw, "user with the given id could not be found", http.StatusNotFound)
		return
	}

	//Check if email should be changed
	if err = validation.Validate(dto.Email, validation.Required, is.Email); err == nil {
		user.Email = dto.Email
	}

	//Check if fullname should be changed
	if err = validation.Validate(dto.FullName, validation.Required, validation.Length(1, 100)); err == nil {
		user.FullName = dto.FullName
		reGenToken = true
	}	

	//Check if the username should be changed
	if err = validation.Validate(dto.Username, validation.Required, validation.Length(4, 100)); err == nil {
		user.Username = dto.Username
		reGenToken = true
	}

	//Update user on database
	err = h.db.UpdateUser(user)
	if err != nil {
		h.l.Errorw("Error during user update", "error", err)
		responses.Error(rw, "user could not be saved to the database", http.StatusUnprocessableEntity)
		// Most probably the username or email ist already in use. You would normally check this before doing the update command and sending specific error codes for each error. (Normal REST API Behaviour)
		return
	}

	h.l.Debugf("Updated user [%v]", dto.ID)

	//Decode DTO
	res := dtos.NewUser(user)

	//Regenerate Token if necessary
	token := ""
	if reGenToken {
		//Session Token and Cookie have to be regenerated because of the new user information.
		token, err = auth.GetJwt(user)
		if err != nil {
			h.l.Errorw("jwt could not be created", "error", err)
			responses.Error(rw, "Something went wrong", http.StatusInternalServerError)
			return
		}

		cookie := auth.GenerateCookie(token, true)
		http.SetCookie(rw, cookie)
		h.l.Debugf("Updated cookie for user")
		return
	}

	//Send response
	_ = utils.ToJSON(&UpdateUserResponse{
		Token: token,
		User:  res,
	}, rw)
}
