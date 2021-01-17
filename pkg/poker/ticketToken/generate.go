package ticketToken

import (
	"github.com/form3tech-oss/jwt-go"
)

func GenerateTicketToken(username, id, lobbyId string, buyIn int, key string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"username": username,
		"id": id,
		"buyIn": buyIn,
		"lobbyId": lobbyId,
	})

	return token.SignedString([]byte(key))
}
