package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jonsch318/royalafg/pkg/dtos"
	"github.com/jonsch318/royalafg/pkg/log"
	"github.com/jonsch318/royalafg/pkg/mw"
)

type HistoryQueryDto struct {
	UserID  string                     `json:"userId"`
	History []dtos.AccountHistoryEvent `json:"history"`
}

// HistoryQuery shows the latest transaction history.
// swagger:response HistoryQueryResponse
type historyQueryWrapper struct {
	// The transaction history
	// in: body
	Body HistoryQueryDto
}

// QueryHistory returns the transaction history for the authenticated user.
// swagger:route POST /api/bank/balance authentication loginUser
//
// Query the users history
//
//	The transaction history of a user
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
//	200: HistoryQueryResponse
func (h *Account) QueryHistory(rw http.ResponseWriter, r *http.Request) {
	//vals := r.URL.Query()
	//i, _ := strconv.Atoi(vals["i"][0])

	i := 0

	claims := mw.FromUserTokenContext(r.Context().Value("user"))
	history, err := h.historyReadModel.GetAccountHistory(claims.ID, i, 200)

	if err != nil {
		log.Logger.Errorw("Query error", "error", err, "ID", claims.ID)
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}

	log.Logger.Debugw("History Printout", "history", history)

	_ = json.NewEncoder(rw).Encode(&HistoryQueryDto{
		UserID:  claims.ID,
		History: history,
	})

}
