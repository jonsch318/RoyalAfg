package models

type Bank interface {
	GetPlayerWallet(string) int
}
