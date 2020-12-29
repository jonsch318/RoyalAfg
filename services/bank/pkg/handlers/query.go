package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type BalanceQueryDto struct {
	UserID string `json:"userId"`
	Balance int `json:"balance"`
}

func (h Account) QueryBalance(rw http.ResponseWriter, r *http.Request) {

	vals := r.URL.Query()
	q, ok := vals["userId"]
	if !ok {
		http.Error(rw, "a valid query is expected", http.StatusBadRequest)
		return
	}

	userID := q[0]

	balance, err := h.balanceReadModel.GetAccountBalance(userID)

	if err != nil {
		log.Printf("Query error: %v", err)
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}

	dto := &BalanceQueryDto{
		UserID: userID,
		Balance: balance,
	}

	_ = json.NewEncoder(rw).Encode(dto)

}