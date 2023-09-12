package logic

import (
	"github.com/Rhymond/go-money"
	"github.com/jonsch318/royalafg/pkg/models"
)

// ProcessGame processes a game and returns the resulting transaction. Money is the absolute amount and if true it is a withdrawl (user lost money) else a deposit (user won money).
func ProcessGame(game *models.SlotGame, factor uint) (*money.Money, bool) {
	win := money.New(game.Win, "EUR")
	win.Subtract(money.New(int64(factor), "EUR"))
	return win.Absolute(), win.IsNegative()
}
