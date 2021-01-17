package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/dtos"
)

func (h *Account) VerifyAmount(rw http.ResponseWriter, r *http.Request)  {
	vals := r.URL.Query()

	userId := vals.Get("userId")
	amount, err := strconv.Atoi(vals.Get("amount"))

	if err != nil || amount < 0 {
		http.Error(rw, "the amount has to be a positive number", http.StatusBadRequest)
	}

	balance, err := h.balanceReadModel.GetAccountBalance(userId)

	if err != nil {
		log.Printf("Query error: %v", err)
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}

	_ = json.NewEncoder(rw).Encode(dtos.VerifyAmount{
		VerificationResult: amount <= balance,
	})

}