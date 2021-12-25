package models

import "github.com/Rhymond/go-money"

type Bank interface {
	GetPlayerWallet(string) *money.Money
	GetMaxBet(string) *money.Money
}
