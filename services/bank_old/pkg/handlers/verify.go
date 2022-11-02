package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Rhymond/go-money"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/currency"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/dtos"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/responses"
)

// VerifyAmount validates.
// swagger:response VerifyAmountResponse
type verifyAmountWrapper struct {
	// The verification result
	// in: body
	Body dtos.VerifyAmount
}

// ErrorResponse is a generic error response
// swagger:response ErrorResponse
type errorResponseWrapper struct {
	// The error
	// in: body
	Body responses.ErrorResponse
}

// VerifyAmount verifies that the amount can be transacted by the user.
// swagger:route POST /api/bank/verifyAmount balance verifyAmount
//
// VerifyAmount verifies the amount against the given user.
//
// This will check the balance of the user and compare the given amount to it.
//
//	Consumes:
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
//	500: ErrorResponse
//	200: VerifyAmountResponse
//
func (h *Account) VerifyAmount(rw http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()

	userId := values.Get("userId")
	amount, err := strconv.Atoi(values.Get("amount"))

	if err != nil || amount < 0 {
		http.Error(rw, "the amount has to be a positive number", http.StatusBadRequest)
	}

	balance, err := h.balanceReadModel.GetAccountBalance(userId)

	if err != nil {
		log.Logger.Errorw("Query error", "error", err)
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}
	res, err := money.New(int64(amount), currency.Code).LessThanOrEqual(balance)
	if err != nil {
		log.Logger.Errorw("verify comparison failed", "error", err)
		res = false
	}
	_ = json.NewEncoder(rw).Encode(dtos.VerifyAmount{
		VerificationResult: res,
	})

}
