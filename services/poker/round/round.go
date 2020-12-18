package round

import (
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/bank"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/utils"
	"sync"
)

//Round is one game of a session. it results in everybody but one folding or a showdown
type Round struct {
	//Players includes all the Players who have started this hand. After a fold the player is still included
	Players         []models.Player
	PublicPlayers   []models.PublicPlayer
	Bank            *bank.Bank
	Board           [5]models.Card
	HoleCards       map[string][2]models.Card
	InCount         byte
	Dealer          int
	Ended           bool
	cardGen         *utils.CardGenerator
	EndCallback     func(int)
	SmallBlind      int
	bigBlindIndex   int
	smallBlindIndex int
	wg              sync.WaitGroup
}

//NewHand creates a new hand and sets the dealer to the next
func NewHand(bank *bank.Bank, smallBlind int) *Round {

	return &Round{
		Bank:       bank,
		cardGen:    utils.NewCardSelector(),
		SmallBlind: smallBlind,
	}
}

func (h *Round) WhileNotEnded(f func()) {
	if !h.Ended {
		f()
	}
}
