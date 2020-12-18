package models

import "fmt"

type Player struct {
	ID       string `json:"id" mapstructure:"buyin"`
	Username string `json:"username" mapstructure:"buyin"`
	BuyIn    int    `json:"buyin" mapstructure:"buyin"`
	Out      chan []byte
	In       chan *Event
	Active   bool
	Close    chan bool
}

func NewPlayer(username, id string, buyin int, in chan *Event, out chan []byte, close chan bool) *Player {
	return &Player{
		ID:       id,
		Username: username,
		BuyIn:    buyin,
		Out:      out,
		In:       in,
		Active:   false,
		Close:    close,
	}
}

func (p *Player) String() string {
	return fmt.Sprintf("Player [%v]: %v has %d", p.ID, p.Username, p.BuyIn)
}
