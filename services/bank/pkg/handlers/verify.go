package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Rhymond/go-money"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/currency"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/dtos"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
)

func (h *Account) VerifyAmount(rw http.ResponseWriter, r *http.Request)  {
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