package handlers

import (
	"encoding/json"
	"net/http"

	ycq "github.com/jetbasrawi/go.cqrs"

	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/commands"
)

type AccountDto struct {
	UserID string `json:"userId"`
}

type TransactionDto struct {
	UserID string `json:"userId"`
	Amount int `json:"amount"`
}

func (h Account) Create(rw http.ResponseWriter, r *http.Request) {

	var dto AccountDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.dispatcher.Dispatch(ycq.NewCommandMessage(dto.UserID, &commands.CreateBankAccount{}))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (h Account) Deposit(rw http.ResponseWriter, r *http.Request) {
	var dto TransactionDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.dispatcher.Dispatch(ycq.NewCommandMessage(dto.UserID, &commands.Deposit{Amount: dto.Amount}))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (h Account) Withdraw(rw http.ResponseWriter, r *http.Request) {
	var dto TransactionDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.dispatcher.Dispatch(ycq.NewCommandMessage(dto.UserID, &commands.Withdraw{Amount: dto.Amount}))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}