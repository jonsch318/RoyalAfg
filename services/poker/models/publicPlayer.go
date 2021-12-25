package models

import (
	"fmt"
)

type PublicPlayer struct {
	Username string  `json:"username" mapstructure:"username"`
	ID       string  `json:"id" mapstructure:"id"`
	BuyIn    string  `json:"buyIn" mapstructure:"buyIn"`
	BuyInNum float64 `json:"buyInNum" mapstructure:"buyInNum"`
}

func (p *PublicPlayer) String() string {
	return fmt.Sprintf("Player [%v] has [%v]", p.Username, p.BuyIn)
}

func (p *PublicPlayer) SetBuyIn(buyIn string, buyInNum float64) {
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
	public.SetBuyIn(b.GetPlayerWallet(public.ID).Display(), b.GetPlayerWallet(public.ID).AsMajorUnits())
	return public
}
