package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jonsch318/royalafg/pkg/dtos"
	"github.com/jonsch318/royalafg/pkg/mw"
	"github.com/jonsch318/royalafg/pkg/responses"
)

type BalanceQueryDto struct {
	Balance *dtos.CurrencyDto `json:"balance"`
	UserID  string            `json:"userId"`
}

// BalanceQueryResponse shows the latest current account balance.
// swagger:response BalanceQueryResponse
type balanceQueryWrapper struct {
	// The balance of the user
	// in: body
	Body HistoryQueryDto
}

// ValidationError shows the failed validation requirements.
// Each form field that has missing requirements is listet under validationErrors
// swagger:response ValidationErrorResponse
type validationErrorWrapper struct {
	// The validation errors
	// in: body
	Body responses.ValidationError
}

// QueryBalance returns the current balance for the authenticated user.
// swagger:route POST /api/bank/balance authentication loginUser
//
// Query the users balance
//
//	The balance of a user
//
//	Consumes:
//
//	Produces:
//	-	application/json
//
//	Security:
//	-	api_key
//
//	Schemes: http, https
//
//	Responses:
//	default: ErrorResponse
//	400: ErrorResponse
//	422: ValidationErrorResponse
//	404: ErrorResponse
//	401: ErrorResponse
//	403: ErrorResponse
//	500: ErrorResponse
//	200: BalanceQueryResponse
func (h Account) QueryBalance(rw http.ResponseWriter, r *http.Request) {
	claims := mw.FromUserTokenContext(r.Context().Value("user"))

	log.Printf("User balance for user %s", claims.ID)
	balance, err := h.balanceReadModel.GetAccountBalance(claims.ID)

	if err != nil {
		log.Printf("Query error: %v", err)
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}

	dto := &BalanceQueryDto{
		UserID:  claims.ID,
		Balance: dtos.FromMoney(balance),
	}

	_ = json.NewEncoder(rw).Encode(dto)

}
