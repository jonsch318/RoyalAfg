package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/dtos"
)

type HistoryQueryDto struct {
	UserID string `json:"userId"`
	History []dtos.AccountHistoryEvent `json:"history"`
}

func (h *Account) QueryHistory(rw http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	i,_ := strconv.Atoi(vals["i"][0])

	claims := mw.FromUserTokenContext(r.Context().Value("user"))

	history, err := h.historyReadModel.GetAccountHistory(claims.ID, i, 25)

	if err != nil {
		log.Printf("Query error %v", err)
		http.Error(rw, err.Error(), http.StatusNotFound)
	}

	_ = json.NewEncoder(rw).Encode(&HistoryQueryDto{
		UserID:  claims.ID,
		History: history,
	})

}
