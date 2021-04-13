package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/poker/ticket"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/utils"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker-matchmaker/pkg/validation"
)

//GetTicketWithID requests a ticket with lobby id
func (h *Ticket) GetTicketWithID(rw http.ResponseWriter, r *http.Request) {
	//CSRF Validation could not be accomplished because the frontend requests tickets server side.
	if err := mw.ValidateCSRF(r); err != nil {
		h.logger.Errorw("could not validate csrf token", "error", err)
		//responses.JSONError(rw, &responses.ErrorResponse{Error: "wrong format decoding failed"}, http.StatusForbidden)
		//return
	}

	//Get user claims
	claims := mw.FromUserTokenContext(r.Context().Value("user"))

	//get information to verify before ticket response
	buyIn, err := strconv.Atoi(r.URL.Query().Get("buyIn"))
	if err != nil {
		h.logger.Errorw("Invalid BuyIn", "error", err)
		http.Error(rw, "the buyIn has to be valid", http.StatusBadRequest)
		return
	}

	if viper.GetBool("include_bank_service_validation") {
		if err = validation.VerifyBuyIn(claims.ID, buyIn); err != nil {
			h.logger.Errorw("Error during bank service validation", "error", err)
			http.Error(rw, "the buyIn has to be lower that the users wallet", http.StatusUnprocessableEntity)
			return
		}
	}

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok || id == "" {
		http.Error(rw, "Either a valid class or a lobby class has to be given", http.StatusBadRequest)
		return
	}

	//Get a ticket
	res, err := h.manager.Connect(id)
	if err != nil {
		h.logger.Errorw("error during connection", "error", err)
		http.Error(rw, "a lobby with the given id is not found", http.StatusNotFound)
		return
	}

	//Generate token
	h.logger.Infof("Creating token for [%v;%v] to join lobby[%v] with %v", claims.Username, claims.ID, res.LobbyId, buyIn)
	token, err := ticket.GenerateTicketToken(claims.Username, claims.ID, res.LobbyId, buyIn, viper.GetString(config.MatchMakerJWTKey))

	//Send response
	_ = utils.ToJSON(&TicketResponse{
		Address: res.Address,
		Token:   token,
	}, rw)

}