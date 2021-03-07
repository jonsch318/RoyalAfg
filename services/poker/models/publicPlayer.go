package models

import (
	"fmt"
)

type PublicPlayer struct {
	Username string  `json:"username" mapstructure:"username"`
	ID       string  `json:"id" mapstructure:"id"`
	BuyIn    string `json:"buyIn" mapstructure:"buyIn"`
}

func (p *PublicPlayer) String() string {
	return fmt.Sprintf("Player [%v] has [%v]", p.Username, p.BuyIn)
}

func (p *PublicPlayer) SetBuyIn(buyIn string) {
	p.BuyIn = buyIn
}

func (p *Player) ToPublic() *PublicPlayer {
	return &PublicPlayer{
		ID:       p.ID,
		Username: p.Username,
	}
}


func (p *Player) ToPublicWithWallet(b Bank) *PublicPlayer {
	public := p.ToPublic()
	public.SetBuyIn(b.GetPlayerWallet(public.ID))
	return public
}
