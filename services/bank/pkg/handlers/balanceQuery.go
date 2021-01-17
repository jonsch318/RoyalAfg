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
	userId := vals.Get("userId")

	balance, err := h.balanceReadModel.GetAccountBalance(userId)

	if err != nil {
		log.Printf("Query error: %v", err)
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}

	dto := &BalanceQueryDto{
		UserID: userId,
		Balance: balance,
	}

	_ = json.NewEncoder(rw).Encode(dto)

}