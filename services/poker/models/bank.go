package models

type Bank interface {
	GetPlayerWallet(string) string
	GetMaxBet(string) string
}
