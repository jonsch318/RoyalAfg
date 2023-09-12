package logic

import (
	"crypto/ecdsa"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jonsch318/royalafg/pkg/models"
	"github.com/jonsch318/royalafg/services/slot/pkg/crypto"
	"github.com/jonsch318/royalafg/services/slot/pkg/database"
	"github.com/jonsch318/royalafg/services/slot/pkg/statistics"
)

type GameProvider struct {
	buffer *database.GameBuffer
	db     *database.GameDatabase
	rng    *crypto.VRFNumberGenerator
}

func NewGameProvider(buffer *database.GameBuffer, db *database.GameDatabase, rng *crypto.VRFNumberGenerator) *GameProvider {
	return &GameProvider{
		buffer: buffer,
		db:     db,
		rng:    rng,
	}
}

func (g *GameProvider) GetGame(id string) (*models.SlotGame, error) {
	//return g.db.GetGame(id)
	return nil, nil
}

func toBase64(val []byte) string {
	return base64.StdEncoding.EncodeToString(val)
}

func (g *GameProvider) GetPublicKey() *ecdsa.PublicKey {
	return g.rng.GetPublicKey()
}

func (g *GameProvider) NewGame(factor uint) (*models.SlotGame, error) {
	//Create new gameId
	gameId := time.Now().Format(time.RFC3339Nano) + "-" + uuid.New().String()

	// wait between 0-100ms randomly
	time.Sleep(time.Duration(g.rng.GenerateNumber(50)) * time.Millisecond)

	//random vrf number
	gameTime, alpha, beta, proof, err := g.rng.Generate()

	if err != nil {
		return nil, err
	}

	numbers := statistics.ParseNumber(beta)

	if numbers == nil {
		return nil, fmt.Errorf("could not parse numbers")
	}

	gameResult := statistics.EvaluateGame(numbers, factor)

	game := models.NewSlotGame(gameId, numbers, gameResult.Amount(), toBase64(proof), toBase64(alpha), toBase64(beta), gameTime)

	return game, nil
	//Save game to buffer

}

func (g *GameProvider) SaveGame(game *models.SlotGame) error {
	return g.buffer.BufferGame(game)
}
