package utils

import (
	"errors"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
)

func SearchByID(players []models.Player, id string) (*models.Player, int, error) {
	for i, n := range players {
		if n.ID == id {
			return &n, i, nil
		}
	}
	return nil, -1, errors.New("Player not in hand, session or lobby")
}
