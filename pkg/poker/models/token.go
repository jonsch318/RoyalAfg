package models

//Token is the model for a ticket token
type Token struct {
	Username string
	Id       string
	BuyIn    int
	LobbyId  string
}

