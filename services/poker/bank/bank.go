package bank

import (
	"sync"

	"github.com/Rhymond/go-money"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/currency"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/rabbit"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/showdown"
)

type Interface interface {
	RegisterLobby(string)
	GetMaxBet() *money.Money
	GetPlayerBet(id string) *money.Money
	GetPlayerWallet(id string) *money.Money
	GetPot() *money.Money
	HasZeroWallet(id string) bool
	PerformBet(id string) error
	PerformRaise(id string, amount *money.Money) (int, error)
	PerformAllIn(id string) (bool, error)
	PerformBlind(id string, amount *money.Money) error
	MustAllIn(id string) (bool, error)
	IsAllIn(id string) bool
	AddPlayer(player *models.Player)
	RemovePlayer(id string) error
	UpdatePublicPlayerBuyIn(p []models.PublicPlayer)
	ConcludeRound(winners []showdown.WinnerInfo, publicPlayers []models.PublicPlayer) []*money.Money
}

//Bank  handles the bets and wallets of players.
type Bank struct {
	lock         sync.RWMutex
	PlayerWallet map[string]*money.Money
	PlayerBets   map[string]*money.Money
	Pot          *money.Money
	MaxBet       *money.Money
	eventBus     *rabbit.RabbitMessageBroker
	LobbyId      string
}

//NewBank creates a new bank to handle the bets and wallets of players.
func NewBank(eventBus *rabbit.RabbitMessageBroker) *Bank {
	return &Bank{
		PlayerWallet: make(map[string]*money.Money),
		PlayerBets:   make(map[string]*money.Money),
		Pot:          currency.Zero(),
		MaxBet:       currency.Zero(),
		LobbyId:      "",
		eventBus:     eventBus,
	}
}

func (b *Bank) RegisterLobby(lobbyId string) {
	b.LobbyId = lobbyId
}

//GetMaxBet returns the highest bet in the current round
func (b *Bank) GetMaxBet() *money.Money {
	//A sync lock is not required because max bet is only be designed to be changed by the game routine and not the lobbies. So concurrent read and writes are not a concern.
	//convert arbitrary precision money value (stored in int64) to int
	return b.MaxBet
}

//GetPlayerBet gets the bet of a given player
func (b *Bank) GetPlayerBet(id string) *money.Money {
	//Have to lock because concurrent read and write are not possible with maps.
	b.lock.RLock()
	defer b.lock.RUnlock()
	t, ok := b.PlayerBets[id]
	if !ok {
		return currency.Zero()
	}
	return t
}

//GetPlayerWallet gets the current wallet for the given player
func (b *Bank) GetPlayerWallet(id string) *money.Money {
	//Have to lock because concurrent read and write are not possible with maps.
	b.lock.RLock()
	defer b.lock.RUnlock()
	t, ok := b.PlayerWallet[id]
	if !ok {
		return currency.Zero()
	}
	return t
}

//GetPot returns the current pot value.
func (b *Bank) GetPot() *money.Money {
	//Have to lock to remove concurrent read and writes.
	b.lock.RLock()
	defer b.lock.RUnlock()
	return b.Pot
}

//HasZeroWallet returns true if the player has zero money left in his wallet or no bank wallet could be found with this id.
func (b *Bank) HasZeroWallet(id string) bool {
	b.lock.RLock()
	defer b.lock.RUnlock()
	t, ok := b.PlayerWallet[id]
	if !ok {
		return true
	}
	return !t.IsPositive()
}
