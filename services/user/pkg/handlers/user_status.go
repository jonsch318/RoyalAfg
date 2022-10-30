package handlers

import (
	"net/http"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/responses"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/utils"
	"github.com/JohnnyS318/RoyalAfgInGo/services/user/pkg/database"
)

type UserStatusResponse struct {
	OnlineStatus *database.OnlineStatus `json:"online"`
	Banned       byte                   `json:"banned"`
	Verified     byte                   `json:"verified"`
}

func (h *UserHandler) UserStatus(rw http.ResponseWriter, r *http.Request) {
	claims := mw.FromUserTokenContext(r.Context().Value("user"))

	//Get online status from redis as a routine
	statusChan := make(chan *database.OnlineStatus)
	go func() {

		status, err := h.statusDB.GetOnlineStatus(claims.ID)
		if err != nil {
			h.l.Errorw("Could not get online status", "error", err)
			statusChan <- nil
		}
		statusChan <- status
	}()

	//Get banned status from database

	user, err := h.db.FindById(claims.ID)

	if err != nil {
		h.l.Errorw("Could not query user", "error", err)
		responses.Error(rw, "user with the given id could not be found", http.StatusNotFound)
	}

	status := <-statusChan

	resp := &UserStatusResponse{
		Banned:       user.Banned,
		Verified:     user.Verified,
		OnlineStatus: status,
	}

	_ = utils.ToJSON(resp, rw)

}
