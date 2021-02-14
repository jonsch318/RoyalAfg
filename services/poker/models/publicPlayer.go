package models

type PublicPlayer struct {
	Username string  `json:"username" mapstructure:"username"`
	ID       string  `json:"id" mapstructure:"id"`
	BuyIn    string `json:"buyIn" mapstructure:"buyIn"`
}

func (player *Player) ToPublic() *PublicPlayer {
	return &PublicPlayer{
		ID:       player.ID,
		Username: player.Username,
	}
}

func (player *Player) ToPublicWithWallet(b Bank) *PublicPlayer {
	p := player.ToPublic()
	p.BuyIn = b.GetPlayerWallet(p.ID)
	return p
}
