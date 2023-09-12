package models

import (
	"fmt"

	"github.com/Rhymond/go-money"

	moneyUtils "github.com/jonsch318/royalafg/services/poker/money"
)

type Player struct {
	//ID the id of the player. should be unique
	ID string `json:"id" mapstructure:"buyin"`
	//Username the username
	Username string `json:"username" mapstructure:"buyin"`
	//BuyIn buy in
	BuyIn *money.Money `json:"buyin" mapstructure:"buyin"`
	//Out channel to the outside. Will be send over websocket.
	Out chan []byte
	//In channel is the incoming channel over the websocket conn.
	In chan *Event
	//Active active
	Active bool
	//Close channel will be closed when player leaves
	Close chan bool
	//Left determins if the player already quitted but is in an active round
	Left bool
}

func NewPlayer(username, id string, buyin int, in chan *Event, out chan []byte, close chan bool) *Player {
	return &Player{
		ID:       id,
		Username: username,
		BuyIn:    moneyUtils.ConvertToIMoney(buyin),
		Out:      out,
		In:       in,
		Active:   false,
		Close:    close,
		Left:     false,
	}
}

func (p *Player) String() string {
	return fmt.Sprintf("Player [%s]: %vs has %s", p.ID, p.Username, p.BuyIn.Display())
}
