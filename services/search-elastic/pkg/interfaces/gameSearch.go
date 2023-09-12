package interfaces

import (
	"github.com/jonsch318/royalafg/services/search-elastic/pkg/dto"
)

type GameSearch interface {
	SearchGames(query string) []dto.GameResult
}
