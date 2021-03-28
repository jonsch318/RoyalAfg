package dtos

import "time"

//VeriyAmount is a result of a verification
type VerifyAmount struct {
	VerificationResult bool `json:"result"`
}

func NewVerifyAmount(res bool) *VerifyAmount {
	return &VerifyAmount{
		VerificationResult: res,
	}
}

//AccountHistoryEvent is a transaction event
type AccountHistoryEvent struct {
	Amount  *CurrencyDto `json:"amount"`
	Type    string       `json:"type"`
	Time    time.Time    `json:"time"`
	Game    string       `json:"gameId"`
	LobbyID string       `json:"roundId"`
}
