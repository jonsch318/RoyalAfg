package commands

type CreateAccount struct {
	UserId string
}

type Deposit struct {
	Amount int
}

type Purchase struct {
	Amount  int
	GameId  string
	Details interface{}
}
