package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
)

type BalanceQueryDto struct {
	UserID string `json:"userId"`
	Balance int `json:"balance"`
}

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
		UserID: claims.ID,
		Balance: balance,
	}

	_ = json.NewEncoder(rw).Encode(dto)

}