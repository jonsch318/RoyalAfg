package bank

import (
	"sync"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/bank"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/rabbit"
)

//Bank  handles the bets and wallets of players.
type Bank struct {
	lock         sync.RWMutex
	PlayerWallet map[string]int
	PlayerBets   map[string]int
	Pot          int
	MaxBet       int
	eventBus *rabbit.RabbitMessageBroker
	eventQueue []bank.Command
	LobbyId string
}

//NewBank creates a new bank to handle the bets and wallets of players.
func NewBank() *Bank {
	return &Bank{
		PlayerWallet: make(map[string]int),
		PlayerBets:   make(map[string]int),
		eventQueue: make([]bank.Command, 0),
	}
}

func (b *Bank) RegisterLobby(lobbyId string) {
	b.LobbyId = lobbyId
}

//GetMaxBet returns the highest bet in the current round
func (b *Bank) GetMaxBet() int {
	//A sync lock is not required because max bet is only be designed to be changed by the game routine and not the lobbies. So concurent read and writes are not possible.
	return b.MaxBet
}

//GetPlayerBet gets the bet of a given player
func (b *Bank) GetPlayerBet(id string) int {
	//Have to lock because concurrent read and write are not possible with maps.
	b.lock.RLock()
	defer b.lock.RUnlock()
	t, ok := b.PlayerBets[id]
	if !ok {
		return -1
	}
	return t
}

//GetPlayerWallet gets the current wallet for the given player
func (b *Bank) GetPlayerWallet(id string) int {
	//Have to lock because concurrent read and write are not possible with maps.
	b.lock.RLock()
	defer b.lock.RUnlock()
	t, ok := b.PlayerWallet[id]
	if !ok {
		return -1
	}
	return t
}
