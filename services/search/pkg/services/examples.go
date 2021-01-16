package services

import (
	"github.com/JohnnyS318/RoyalAfgInGo/services/search/pkg/dto"
)

func LoadExampleDbIndexes() []string {
	return []string{
		"poker",
		"pacman",

	}
}

func LoadExampleDb() map[string]dto.GameResult{
	m := make(map[string]dto.GameResult, 2)
	m["poker"] = dto.GameResult{Name: "Poker", URL: "/games/poker", MaxPlayers: 10}
	m["pacman"] = dto.GameResult{Name: "Pacman", URL: "/games/pacman", MaxPlayers: 1}
	return m
}
