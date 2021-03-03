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

type JoinSuccess struct {
	Players     []models.PublicPlayer `json:"players" mapstructure:"players"`
	Wallet string	`json:"wallet"`
	LobbyID     string                `json:"lobbyId" mapstructure:"lobbyId"`
	MaxBuyIn    int                   `json:"maxBuyIn"`
	MinBuyIn    int                   `json:"minBuyIn"`
	BigBlind    int                   `json:"bigBlind"`
	Position    int                   `json:"position" mapstructure:"position"`
	GameState   byte                  `json:"gameState" mapstructure:"gameState"`
	GameStarted bool                  `json:"gameStarted" mapstructure:"gameStarted"`
}

func NewJoinSuccessEvent(lobbyId string, players []models.PublicPlayer, gameStarted bool, gameState byte, position, maxBuyIn, minBuyIn, bigBlind int, wallet string) *models.Event {
	return models.NewEvent(JOIN_SUCCESS, &JoinSuccess{
		LobbyID:     lobbyId,
		Players:     players,
		GameStarted: gameStarted,
		MaxBuyIn:    maxBuyIn,
		MinBuyIn:    minBuyIn,
		BigBlind:    bigBlind,
		Position:    position,
		GameState: gameState,
		Wallet: wallet,
	})
}

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

type PlayerJoinEvent struct {
	Player *models.PublicPlayer `json:"player"`
	Index  int                  `json:"index"`
}

func NewPlayerJoinEvent(player *models.PublicPlayer, index int) *models.Event {
	return models.NewEvent(PLAYER_JOIN, &PlayerJoinEvent{
		Player: player,
		Index:  index,
	})
}
