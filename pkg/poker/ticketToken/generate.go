package ticketToken

import (
	"time"

	"github.com/form3tech-oss/jwt-go"
)

func GenerateTicketToken(username, id, lobbyId string, buyIn int, key string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"id":       id,
		"buyIn":    buyIn,
		"lobbyId":  lobbyId,
		"exp":      time.Now().Add(1 * time.Minute).Unix(),
	})

	return token.SignedString([]byte(key))
}
