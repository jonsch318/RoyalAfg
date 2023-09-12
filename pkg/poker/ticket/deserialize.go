package ticket

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/jonsch318/royalafg/pkg/poker/models"
)

// Returns a deserialized token from a jwt claims map
func FromToken(claims jwt.MapClaims) *models.Token {
	buyIn, ok := claims["buyIn"].(float64)
	if !ok {
		return nil
	}
	return &models.Token{
		Username: claims["username"].(string),
		Id:       claims["id"].(string),
		BuyIn:    int(buyIn),
		LobbyId:  claims["lobbyId"].(string),
	}
}
