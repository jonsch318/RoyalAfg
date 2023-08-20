package interfaces

import (
	"github.com/JohnnyS318/RoyalAfgInGo/services/search-elastic/pkg/dto"
)

type GameSearch interface {
	SearchGames(query string) []dto.GameResult
}
