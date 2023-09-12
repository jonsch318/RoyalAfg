package handlers

import (
	"go.uber.org/zap"

	"github.com/jonsch318/royalafg/services/search-elastic/pkg/interfaces"
)

type Game struct {
	logger        *zap.SugaredLogger
	searchService interfaces.GameSearch
}

func NewGameHandler(logger *zap.SugaredLogger, searchService interfaces.GameSearch) *Game {
	return &Game{
		logger:        logger,
		searchService: searchService,
	}
}
