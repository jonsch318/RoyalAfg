package round

import (
	"github.com/Rhymond/go-money"

	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/bank"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	moneyUtils "github.com/JohnnyS318/RoyalAfgInGo/services/poker/money"
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
	SmallBlind      *money.Money
	bigBlindIndex   int
	smallBlindIndex int
	wg              sync.WaitGroup
}

//NewHand creates a new hand and sets the dealer to the next
func NewHand(b *bank.Bank, smallBlind int) *Round {

	return &Round{
		Bank:       b,
		cardGen:    utils.NewCardSelector(),
		SmallBlind: moneyUtils.ConvertToIMoney(smallBlind),
	}
}

func (r *Round) WhileNotEnded(f func()) {
	if !r.Ended {
		f()
	}
}
