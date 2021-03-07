package models

import (
	"github.com/form3tech-oss/jwt-go"
)

type Token struct {
	Username string
	Id string
	BuyIn int
	LobbyId string
}

func FromToken(claims jwt.MapClaims) *Token {
	buyIn, ok := claims["buyIn"].(float64)
	if !ok {
		return nil
	}
	return &Token{
		Username: claims["username"].(string),
		Id:       claims["id"].(string),
		BuyIn: int(buyIn),
		LobbyId:  claims["lobbyId"].(string),
	}
}