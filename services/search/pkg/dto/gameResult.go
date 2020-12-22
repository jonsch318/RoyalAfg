package dto

type GameSearchResult struct {
	Games []GameResult
}

type GameResult struct {
	Name string `json:"name"`
	URL string `json:"url"`
	MaxPlayers int `json:"maxPlayers"`
}