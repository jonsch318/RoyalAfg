package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/dtos"
)

type HistoryQueryDto struct {
	UserID string `json:"userId"`
	History []dtos.AccountHistoryEvent `history`
}

func (h *Account) QueryHistory(rw http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	q, ok := vals["userId"]
	if !ok {
		http.Error(rw, "a valid query is expected", http.StatusBadRequest)
		return
	}
	userID := q[0]

	i,_ := strconv.Atoi(vals["i"][0])

	history, err := h.historyReadModel.GetAccountHistory(userID, i, 25)

	if err != nil {
		log.Printf("Query error %v", err)
		http.Error(rw, err.Error(), http.StatusNotFound)
	}

	_ = json.NewEncoder(rw).Encode(&HistoryQueryDto{
		UserID:  userID,
		History: history,
	})

}
