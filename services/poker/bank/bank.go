package bank

import (
	"sync"

	"github.com/Rhymond/go-money"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/bank"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/rabbit"
)

//Bank  handles the bets and wallets of players.
type Bank struct {
	lock         sync.RWMutex
	PlayerWallet map[string]*money.Money
	PlayerBets   map[string]*money.Money
	Pot          *money.Money
	MaxBet       *money.Money
	eventBus     *rabbit.RabbitMessageBroker
	eventQueue   []bank.Command
	LobbyId      string
}

//NewBank creates a new bank to handle the bets and wallets of players.
func NewBank(eventBus *rabbit.RabbitMessageBroker) *Bank {
	return &Bank{
		PlayerWallet: make(map[string]*money.Money),
		PlayerBets:   make(map[string]*money.Money),
		eventBus:     eventBus,
		eventQueue:   make([]bank.Command, 0),
	}
}

func (b *Bank) RegisterLobby(lobbyId string) {
	b.LobbyId = lobbyId
}

//GetMaxBet returns the highest bet in the current round
func (b *Bank) GetMaxBet() string {
	//A sync lock is not required because max bet is only be designed to be changed by the game routine and not the lobbies. So concurrent read and writes are not a concern.
	//convert arbitrary precision money value (stored in int64) to int
	return b.MaxBet.Display()
}

//GetPlayerBet gets the bet of a given player
func (b *Bank) GetPlayerBet(id string) string {
	//Have to lock because concurrent read and write are not possible with maps.
	b.lock.RLock()
	defer b.lock.RUnlock()
	t, ok := b.PlayerBets[id]
	if !ok {
		return ""
	}
	return t.Display()
}

//GetPlayerWallet gets the current wallet for the given player
func (b *Bank) GetPlayerWallet(id string) string {
	//Have to lock because concurrent read and write are not possible with maps.
	b.lock.RLock()
	defer b.lock.RUnlock()
	t, ok := b.PlayerWallet[id]
	if !ok {
		return ""
	}
	return t.Display()
}
