package handlers

import (
	"net/http"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/responses"
)

func (h *UserHandler) DeleteUser(rw http.ResponseWriter, r *http.Request)  {

	claims := mw.FromUserTokenContext(r.Context().Value("user"))

	if err := mw.ValidateCSRF(r); err != nil {
		h.l.Errorw("could not validate csrf token", "error", err)
		responses.Error(rw, "wrong format decoding failed", http.StatusForbidden)
		return
	}

	user, err := h.db.FindById(claims.ID)

	if err != nil {
		h.l.Errorw("error during user search", "error", err)
		responses.Error(rw, "logged in user not found", http.StatusNotFound)
		return
	}

	err = h.db.DeleteUser(user)

	//TODO: Delete from bank ().
	//Message through rabbitmq user deleted.

	if err != nil {
		h.l.Errorw("error during deletion", "error", err)
		responses.Error(rw, "unable to delete user", http.StatusInternalServerError)
	}
	h.l.Debugf("user deleted %v", claims.ID)

	rw.WriteHeader(http.StatusOK)
}