package round

import (
	"github.com/Rhymond/go-money"

	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/bank"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	moneyUtils "github.com/JohnnyS318/RoyalAfgInGo/services/poker/money"
)

//Round is one game of a session. it results in everybody but one folding or a showdown
//Probably this could be optimised into multiple structs.
type Round struct {
	//Players includes all the Players who have started this hand. After a fold the player is still included
	Players         []models.Player
	PublicPlayers   []models.PublicPlayer
	Bank            bank.Interface
	Board           [5]models.Card
	HoleCards       map[string][2]models.Card
	SmallBlind      *money.Money
	Dealer          int
	bigBlindIndex   int
	smallBlindIndex int
	InCount         byte
	Ended           bool
}

//NewHand creates a new hand and sets the dealer to the next
func NewHand(b bank.Interface, smallBlind int) *Round {

	return &Round{
		Bank:      b,
		SmallBlind: moneyUtils.ConvertToIMoney(smallBlind),
	}
}

func (r *Round) WhileNotEnded(f func()) {
	if !r.Ended {
		f()
	}
}
