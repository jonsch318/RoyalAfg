package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Rhymond/go-money"
	ycq "github.com/jetbasrawi/go.cqrs"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/currency"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/responses"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/commands"
)

type TransactionDto struct {
	Amount int `json:"amount"`
}

// NoContentResponse is an empty object with no content
// swagger:response NoContentResponse
type noContentResponseWrapper struct {
}

// Deposit deposits the specified amount to the user.
// swagger:route POST /api/bank/deposit transaction bankTransaction
//
// deposits the specified amount to the user.
//
//	Consumes:
//	-	application/json
//
// 	Produces:
//	-	application/json
//
//	Security:
//	-	api_key
//
//	Schemes: http, https
//
// 	Responses:
//	default: ErrorResponse
//	400: ErrorResponse
//	404: ErrorResponse
//	401: ErrorResponse
//	403: ErrorResponse
//	500: ErrorResponse
//	200: NoContentResponse
//
func (h Account) Deposit(rw http.ResponseWriter, r *http.Request) {

	if err := mw.ValidateCSRF(r); err != nil {
		log.Logger.Errorw("could not validate csrf token", "error", err)
		responses.Error(rw, "wrong format decoding failed", http.StatusForbidden)
		return
	}

	claims := mw.FromUserTokenContext(r.Context().Value("user"))
	var dto TransactionDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		log.Logger.Errorw("Dto deserialization error", "error", err)
		responses.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	log.Logger.Debugf("Depositing %v €", dto.Amount)

	err = h.dispatcher.Dispatch(ycq.NewCommandMessage(claims.ID, &commands.Deposit{
		Amount:  money.New(int64(dto.Amount), currency.Code),
		GameId:  "",
		RoundId: "",
		Time:    time.Now(),
	}))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (h Account) Withdraw(rw http.ResponseWriter, r *http.Request) {

	if err := mw.ValidateCSRF(r); err != nil {
		log.Logger.Errorw("could not validate csrf token", "error", err)
		responses.JSONError(rw, &responses.ErrorResponse{Error: "wrong format decoding failed"}, http.StatusForbidden)
		return
	}

	claims := mw.FromUserTokenContext(r.Context().Value("user"))

	var dto TransactionDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.dispatcher.Dispatch(ycq.NewCommandMessage(claims.ID, &commands.Withdraw{
		Amount:  money.New(int64(dto.Amount), currency.Code),
		GameId:  "",
		RoundId: "",
		Time:    time.Now(),
	}))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}
