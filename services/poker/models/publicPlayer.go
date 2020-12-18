package models

type PublicPlayer struct {
	Username string  `json:"username" mapstructure:"username"`
	ID       string  `json:"id" mapstructure:"id"`
	BuyIn    float32 `json:"buyIn" mapstructure:"buyIn"`
}

func (player *Player) ToPublic() *PublicPlayer {
	return &PublicPlayer{
		ID:       player.ID,
		Username: player.Username,
	}
}

func (player *Player) ToPublicWithWallet(b Bank) *PublicPlayer {
	p := player.ToPublic()
	p.BuyIn = float32(b.GetPlayerWallet(p.ID)) / 100
	return p
}
