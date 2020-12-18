package lobby

import (
	"math/rand"
	"strings"
	"time"
)

const idLength = 5

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

//GenerateLobbyID generates a new lobby id out of random letters
func GenerateLobbyID() string {
	rand.Seed(time.Now().UnixNano())
	sb := strings.Builder{}
	sb.Grow(idLength)
	for i := 0; i < idLength; i++ {
		sb.WriteByte(letterBytes[rand.Intn(len(letterBytes))])
	}
	return sb.String()
}
