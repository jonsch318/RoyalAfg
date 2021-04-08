package lobby

import (
	"sync"

	sdk "agones.dev/agones/sdks/go"

	pokerModels "github.com/JohnnyS318/RoyalAfgInGo/pkg/poker/models"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/bank"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/queue"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/round"
)

//Lobby is the Parent structure for a poker lobby it handles the player joins and removals and the game starts.
type Lobby struct {
	pokerModels.LobbyBase
	lock          sync.RWMutex
	Players       []models.Player
	PublicPlayers []models.PublicPlayer
	PlayerQueue   *queue.PlayerQueue
	RemovalQueue  queue.IQueue
	Bank          bank.Interface
	round         round.Interface
	sdk           *sdk.SDK
	dealer        int
	c             chan bool
	GameStarted   bool
}

//NewLobby creates a new lobby object
func NewLobby(b bank.Interface, sdk *sdk.SDK) *Lobby {
	return &Lobby{
		Players:       make([]models.Player, 0),
		PublicPlayers: make([]models.PublicPlayer, 0),
		PlayerQueue:   queue.NewPlayer(),
		RemovalQueue:  queue.New(),
		Bank:          b,
		dealer:        -1,
		c:             make(chan bool, 1),
		sdk:           sdk,
	}
}

func (l *Lobby) RegisterLobbyValue(class *pokerModels.Class, classIndex int, id string) {
	l.Class = class
	l.ClassIndex = classIndex
	l.round = round.NewRound(l.Bank, class.Blind)
	l.LobbyID = id
}

//GetGameStarted determines whether a game has already in this lobby started.
func (l *Lobby) GetGameStarted() bool {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.GameStarted
}

//HasToBeAdded determines whether there are any pending lobby joins
func (l *Lobby) HasToBeAdded() bool {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.PlayerQueue.Length() > 0
}

//HasToBeRemoved determines whether there are any pending removals
func (l *Lobby) HasToBeRemoved() bool {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.RemovalQueue.Length() != 0
}
