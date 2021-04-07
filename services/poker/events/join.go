package events

import (
	"errors"

	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"

	"github.com/mitchellh/mapstructure"
)

type JoinEvent struct {
	Token string `json:"token"`
}

func ToJoinEvent(raw *models.Event) (*JoinEvent, error) {

	if !ValidateEventName(JOIN, raw.Name) {
		return nil, errors.New(REQUIRED_EVENT_NAME_MISSING)
	}

	event := new(JoinEvent)
	err := mapstructure.Decode(raw.Data.(map[string]interface{}), event)
	return event, err
}

type LobbyInfo struct {
	LobbyID          string `json:"lobbyId"`
	MaxBuyIn         int    `json:"maxBuyIn"`
	MinBuyIn         int    `json:"minBuyIn"`
	Blind            int    `json:"blind"`
	MinToStart       int    `json:"minPlayersToStart"`
	PlayerCount      int    `json:"playerCount"`
	GameStartTimeout int    `json:"gameStartTimeout"`
	GameStarted      bool   `json:"gameStarted"`
}

//NewLobbyInfoEvent
func NewLobbyInfoEvent(lobbyId string, players, minToStart, maxBuyIn, minBuyIn, blind, gameStartTimeout int, gameStarted bool) *models.Event {
	return models.NewEvent(LOBBY_INFO, &LobbyInfo{
		LobbyID:          lobbyId,
		MaxBuyIn:         maxBuyIn,
		MinBuyIn:         minBuyIn,
		Blind:            blind,
		MinToStart:       minToStart,
		PlayerCount:      players,
		GameStarted:      gameStarted,
		GameStartTimeout: gameStartTimeout,
	})
}

type JoinSuccess struct {
	Players     []models.PublicPlayer `json:"players"`
	BuyIn       string                `json:"buyin"`
	Position    int                   `json:"position"`
	GameStarted bool                  `json:"gameStarted"`
}

func NewJoinSuccessEvent(players []models.PublicPlayer, position int, buyIn string, gameStarted bool) *models.Event {
	return models.NewEvent(JOIN_SUCCESS, &JoinSuccess{
		Players:     players,
		Position:    position,
		BuyIn:       buyIn,
		GameStarted: gameStarted,
	})
}

//PlayerLeavesEvent messages that a player left
type PlayerLeavesEvent struct {
	Player      *models.PublicPlayer `json:"player"`
	Index       int                  `json:"index"`
	PlayerCount int                  `json:"playerCount"`
	GameStarted bool                 `json:"gameStarted"`
}

func NewPlayerLeavesEvent(player *models.PublicPlayer, i, playerCount int, gameStarted bool) *models.Event {
	return models.NewEvent(PLAYER_LEAVE, &PlayerLeavesEvent{
		Player:      player,
		Index:       i,
		PlayerCount: playerCount,
		GameStarted: gameStarted,
	})
}

//NewPlayerJoinEvent is the event for all other players to receive when another player joins. The joining player receives a join success message
type PlayerJoinEvent struct {
	Player      *models.PublicPlayer `json:"player"`
	Index       int                  `json:"index"`
	PlayerCount int                  `json:"playerCount"`
	GameStarted bool                 `json:"gameStarted"`
}

//NewPlayerJoinEvent creates a player join event for all other players to receive. The joining player receives a join success message
func NewPlayerJoinEvent(player *models.PublicPlayer, index, playerCount int, gameStarted bool) *models.Event {
	return models.NewEvent(PLAYER_JOIN, &PlayerJoinEvent{
		Player:      player,
		Index:       index,
		PlayerCount: playerCount,
		GameStarted: gameStarted,
	})
}
