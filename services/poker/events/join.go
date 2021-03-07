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
	LobbyID     string `json:"lobbyId"`
	MaxBuyIn    int    `json:"maxBuyIn"`
	MinBuyIn    int    `json:"minBuyIn"`
	Blind       int    `json:"blind"`
	MinToStart  int    `json:"minPlayersToStart"`
	PlayerCount int    `json:"playerCount"`
	GameStartTimeout int `json:"gameStartTimeout"`
	GameStarted bool   `json:"gameStarted"`
}

//NewLobbyInfoEvent
func NewLobbyInfoEvent(lobbyId string, players, minToStart, maxBuyIn, minBuyIn, blind, gameStartTimeout int, gameStarted bool ) *models.Event {
	return models.NewEvent(LOBBY_INFO, &LobbyInfo{
		LobbyID:     lobbyId,
		MaxBuyIn:    maxBuyIn,
		MinBuyIn:    minBuyIn,
		Blind:       blind,
		MinToStart:  minToStart,
		PlayerCount: players,
		GameStarted: gameStarted,
		GameStartTimeout: gameStartTimeout,
	})
}

type JoinSuccess struct {
	Players  []models.PublicPlayer `json:"players"`
	Wallet   string                `json:"wallet"`
	Position int                   `json:"position"`
}

func NewJoinSuccessEvent(players []models.PublicPlayer, position int, wallet string) *models.Event {
	return models.NewEvent(JOIN_SUCCESS, &JoinSuccess{
		Players:  players,
		Position: position,
		Wallet:   wallet,
	})
}

//PlayerLeavesEvent messages that a player left
type PlayerLeavesEvent struct {
	Player *models.PublicPlayer `json:"player"`
	Index  int                  `json:"index"`
}

func NewPlayerLeavesEvent(player *models.PublicPlayer, i int) *models.Event {
	return models.NewEvent(PLAYER_LEAVE, &PlayerLeavesEvent{
		Player: player,
		Index:  i,
	})
}

//NewPlayerJoinEvent is the event for all other players to receive when another player joins. The joining player receives a join success message
type PlayerJoinEvent struct {
	Player      *models.PublicPlayer `json:"player"`
	Index       int                  `json:"index"`
	PlayerCount int                  `json:"playerCount"`
}

//NewPlayerJoinEvent creates a player join event for all other players to receive. The joining player receives a join success message
func NewPlayerJoinEvent(player *models.PublicPlayer, index, playerCount int) *models.Event {
	return models.NewEvent(PLAYER_JOIN, &PlayerJoinEvent{
		Player:      player,
		Index:       index,
		PlayerCount: playerCount,
	})
}
