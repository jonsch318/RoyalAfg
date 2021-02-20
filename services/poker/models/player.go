package models

import (
	"fmt"

	"github.com/Rhymond/go-money"

	moneyUtils "github.com/JohnnyS318/RoyalAfgInGo/services/poker/money"
)

type Player struct {
	ID       string `json:"id" mapstructure:"buyin"`
	Username string `json:"username" mapstructure:"buyin"`
	BuyIn    *money.Money    `json:"buyin" mapstructure:"buyin"`
	Out      chan []byte
	In       chan *Event
	Active   bool
	Close    chan bool
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
	}
}

func (p *Player) String() string {
	return fmt.Sprintf("Player [%s]: %vs has %s", p.ID, p.Username, p.BuyIn.Display())
}
