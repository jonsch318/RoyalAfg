package lobby

import (
	"errors"
	"log"
	"strconv"
	"sync"
	"time"

	sdk "agones.dev/agones/sdks/go"

	pokerModels "github.com/JohnnyS318/RoyalAfgInGo/pkg/poker/models"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/bank"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/events"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/round"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/serviceconfig"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/utils"

	"github.com/spf13/viper"
)

//Lobby is the Parent structure for a poker lobby it handles the player joins and removals and the game starts.
type Lobby struct {
	pokerModels.LobbyBase
	lock          sync.RWMutex
	Players       []models.Player
	PublicPlayers []models.PublicPlayer `json:"players"`
	GameStarted   bool
	ToBeRemoved   []int
	ToBeAdded     []*models.Player
	PlayerQueue   []*models.Player
	Bank          *bank.Bank
	dealer        int
	round         *round.Round
	c             chan bool
	sdk           *sdk.SDK
}

//NewLobby creates a new lobby object
func NewLobby(bank *bank.Bank, sdk *sdk.SDK) *Lobby {
	return &Lobby{
		Players:     make([]models.Player, 0),
		ToBeRemoved: make([]int, 0),
		PlayerQueue: make([]*models.Player, 0),
		Bank:        bank,
		dealer:      -1,
		c:           make(chan bool, 1),
		sdk:         sdk,
	}
}

func (l *Lobby) RegisterLobbyValue(class *pokerModels.Class, classIndex int, id string) {
	l.Class = class
	l.ClassIndex = classIndex
	l.round = round.NewHand(l.Bank, class.Blind)
	l.LobbyID = id
}

//TotalPlayerCount returns the total player count in queue and already joined
func (l *Lobby) TotalPlayerCount() int {
	l.lock.RLock()
	c := len(l.Players)
	c += len(l.PlayerQueue)
	c += len(l.ToBeAdded)
	l.lock.RUnlock()
	return c
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
	return len(l.ToBeAdded) > 0
}

//HasToBeRemoved determines whether there are any pending removals
func (l *Lobby) HasToBeRemoved() bool {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return len(l.ToBeRemoved) > 0
}

//Join adds the player to the lobby and starts the game if the minimum player count is exceeded
func (l *Lobby) Join(player *models.Player) {

	err := l.sdk.Allocate()
	if err != nil {
		log.Fatalf("Error during Allocation: %s", err)
	}
	le := len(l.Players) + len(l.ToBeAdded)
	if le <= 10 {
		l.lock.Lock()
		l.ToBeAdded = append(l.ToBeAdded, player)
		l.lock.Unlock()
		gameStarted := l.GetGameStarted()
		if gameStarted {
			return
		}

		l.EmptyToBeAdded()

		utils.SendToPlayer(player, events.NewJoinSuccessEvent(
			l.LobbyID,
			l.PublicPlayers,
			l.GetGameStarted(),
			0,
			len(l.Players)+len(l.ToBeAdded)+len(l.PlayerQueue), l.Class.Max, l.Class.Min, l.Class.Blind))

		if len(l.Players) >= viper.GetInt(serviceconfig.PlayersRequiredForStart) && !gameStarted {
			l.Start()
		}
	} else {
		l.EnqueuePlayer(player)
	}
}

//EmptyToBeAdded adds all pending players to the lobby
func (l *Lobby) EmptyToBeAdded() {
	l.lock.Lock()
	for i := range l.ToBeAdded {
		if len(l.Players) < 10 {
			for j := range l.Players {
				if l.Players[j].ID == l.ToBeAdded[i].ID {
					continue
				}
			}
			player := l.ToBeAdded[i]
			public := l.ToBeAdded[i].ToPublic()
			public.BuyIn = l.ToBeAdded[i].BuyIn.Display()
			if len(l.Players) > 0 {
				utils.SendToAll(l.Players, events.NewPlayerJoinEvent(public, len(l.Players)-1))
			}
			j := len(l.Players)
			l.Players = append(l.Players, *player)
			l.PlayerCount++
			l.PublicPlayers = append(l.PublicPlayers, *public)

			go func() {
				<-l.Players[j].Close
				_ = l.RemovePlayerByID(l.Players[j].ID)
			}()
			log.Printf("Adding now")

			l.Bank.AddPlayer(player)
			l.Bank.UpdatePublicPlayerBuyIn(l.PublicPlayers)

			time.Sleep(100)

			l.lock.Unlock()
			utils.SendToPlayer(player, events.NewJoinSuccessEvent(l.LobbyID, l.PublicPlayers, l.GetGameStarted(), 0, len(l.Players)+len(l.ToBeAdded)+len(l.PlayerQueue), l.Class.Max, l.Class.Min, l.Class.Blind))
			//l.ToBeAdded[i].Out <- events.NewJoinSuccessEvent(l.LobbyID, l.PublicPlayers, l.GetGameStarted(), 0, len(l.Players)+len(l.ToBeAdded)+len(l.PlayerQueue), l.MaxBuyIn, l.MinBuyIn, l.SmallBlind).ToRaw()

			log.Printf("Player joined lobby [%v] count: %v", l.LobbyID, len(l.Players))

			l.SetPlayerCountLabel()

		} else {
			l.EnqueuePlayer(l.ToBeAdded[i])
			l.lock.Unlock()
		}
	}
	l.ToBeAdded = nil

}

func (l *Lobby) SetPlayerCountLabel() {
	log.Printf("Players %v", l.TotalPlayerCount())
	err := l.sdk.SetLabel("players", strconv.Itoa(l.TotalPlayerCount()))
	if err != nil {
		log.Printf("error during player label set %v", err.Error())
	}
}

//RemovePlayerByID removes the given player identified by his id
func (l *Lobby) RemovePlayerByID(id string) error {

	i := l.FindPlayerByID(id)

	if i < 0 {
		return errors.New("The player is not in the lobby")
	}

	l.lock.Lock()
	l.ToBeRemoved = append(l.ToBeRemoved, i)
	l.lock.Unlock()
	_ = l.round.Fold(id)
	if !l.GetGameStarted() {
		l.RemoveAfterGame()
	}

	return nil
}

//RemoveAfterGame removes the left players from the lobby after a game has finished. During a game the player is counted as folded.
func (l *Lobby) RemoveAfterGame() {
	l.lock.Lock()
	defer l.lock.Unlock()
	for _, i := range l.ToBeRemoved {
		if len(l.Players) > i {
			_ = l.Bank.RemovePlayer(l.Players[i].ID)
			if len(l.PlayerQueue) > 0 {
				player, ok := l.DequeuePlayer(i)
				if ok {
					l.Join(player)
				}
			} else {
				l.Players = append(l.Players[:i], l.Players[i+1:]...)
				l.PlayerCount--
				l.PublicPlayers = append(l.PublicPlayers[:i], l.PublicPlayers[i+1:]...)
			}

		}
	}
	log.Printf("Updated Playerlist in RoundId: %v", l.Players)

	log.Printf("Check if player count is 0")
	l.SetPlayerCountLabel()
	if len(l.Players) < 1 {
		err := l.sdk.Shutdown()
		if err != nil {
			log.Fatal("Error shutting down")
		}
	}
}
