package events

import (
	"errors"

	"github.com/Rhymond/go-money"

	"github.com/mitchellh/mapstructure"

	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	moneyUtils "github.com/JohnnyS318/RoyalAfgInGo/services/poker/money"
)

//FOLD describes the action of a player quiting this hand
const FOLD = 1

//BET describes the action of a player betting the same amount as the highes bet and therefore go along or callling the hand
const BET = 2

//RAISE raises sets the highest bet a certain amount
const RAISE = 3

//CHECK action pushes the action requirement to the next player
const CHECK = 4

const ALL_IN = 5

//Action describes a action a player can make one a normal hand stage
type Action struct {
	Action  int `json:"action" mapstructure:"action"`
	Payload *money.Money `json:"payload" mapstructure:"payload"`
}


//Action describes a action a player can make one a normal hand stage
type ActionDTO struct {
	Action  int `json:"action" mapstructure:"action"`
	Payload int `json:"payload" mapstructure:"payload"`
}

func ToActionDTO(raw *models.Event) (*ActionDTO, error) {

	if !ValidateEventName(PLAYER_ACTION, raw.Name) {
		return nil, errors.New(REQUIRED_EVENT_NAME_MISSING)
	}

	event := new(ActionDTO)
	err := mapstructure.Decode(raw.Data.(map[string]interface{}), event)
	return event, err
}

func ToAction(raw *models.Event) (*Action, error) {
	event, err := ToActionDTO(raw)
	if err != nil {
		return nil, err
	}
	return &Action{
		Action:  event.Action,
		Payload: moneyUtils.ConvertToIMoney(event.Payload),
	}, nil
}

//WaitForActionEvent encodes all possible actions the user can perform.
type WaitForActionEvent struct {
	Position        int  `json:"position" mapstructure:"position"`
	PossibleActions byte `json:"possibleActions" mapstructure:"possibleActions"`
}

// NewWaitForAction is an event that the server is waiting for an action from a given player. The possible actions range from 0001 = Fold | 0010=Bet | 0100=Raise | 1000=Check to 1111=All
func NewWaitForActionEvent(position int, possibleActions byte) *models.Event {
	return models.NewEvent(WAIT_FOR_PLAYER_ACTION, &WaitForActionEvent{Position: position, PossibleActions: possibleActions})
}

type ActionProcessedEvent struct {
	Action int `json:"action" mapstructure:"action"`
	//Player   *models.PublicPlayer `json:"player" mapstructure:"player"`
	Position    int     `json:"position" mapstructure:"position"`
	Amount      string `json:"amount" mapstructure:"amount"`
	TotalAmount string `json:"totalAmount" mapstructure:"totalAmount"`
	BuyIn       string `json:"buyIn" mapstructure:"buyIn"`
}

func NewActionProcessedEvent(action, position int, amount, totalAmount, wallet string) *models.Event {
	return models.NewEvent(ACTION_PROCESSED, &ActionProcessedEvent{
		Action:      action,
		Position:    position,
		Amount:      amount,
		TotalAmount: amount,
		BuyIn:       amount,
		//Player:   player.ToPublic(),
	})
}
