package handlers

import (
	"net/http"
	"strconv"

	"github.com/spf13/viper"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/poker/pokerConfig"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/poker/ticketToken"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/utils"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker-matchmaker/pkg/validation"
)

//TicketResponse is the successful response of a ticket request
type TicketResponse struct {
	Address string `json:"address"`
	Token   string `json:"token"`
}

//GetTicketWithParams requests a ticket with lobby params
func (h *Ticket) GetTicketWithParams(rw http.ResponseWriter, r *http.Request) {

	//CSRF Validation could not be accomplished because the frontend requests tickets server side.
	if err := mw.ValidateCSRF(r); err != nil {
		h.logger.Errorw("could not validate csrf token. Continuing with request....", "error", err)
		//responses.JSONError(rw, &responses.ErrorResponse{Error: "wrong format decoding failed"}, http.StatusForbidden)
		//return
	}

	//Get user claims
	claims := mw.FromUserTokenContext(r.Context().Value("user"))

	vals := r.URL.Query()
	class, err := strconv.Atoi(vals.Get("class"))
	if err != nil {
		h.logger.Errorw("Invalid Class", "error", err)
		http.Error(rw, "Either a valid class or a lobby class has to be given", http.StatusBadRequest)
		return
	}

	buyIn, err := strconv.Atoi(vals.Get("buyIn"))
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

	//Get ticket from manager
	res, err := h.manager.RequestTicket(class)
	if err != nil {
		h.logger.Errorw("Error during ticket request", "error", err)
		http.Error(rw, "something went wrong during a lobby search", http.StatusInternalServerError)
		return
	}

	//Generate token
	h.logger.Infow("Generate Ticket", "username", claims.Username, "id", claims.ID, "lobbyId", res.LobbyId, "buyIn", buyIn)
	token, err := ticketToken.GenerateTicketToken(claims.Username, claims.ID, res.LobbyId, buyIn, viper.GetString(pokerConfig.MatchMakerJWTKey))

	//Send response
	_ = utils.ToJSON(&TicketResponse{
		Address: res.Address,
		Token:   token,
	}, rw)
}



