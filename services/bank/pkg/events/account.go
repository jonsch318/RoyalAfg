package events

type AccountCreated struct {
	ID     string
	UserId string
}

type Deposited struct {
	ID     string
	Amount int
}

type Purchased struct {
	ID      string
	Amount  int
	GameID  string
	Details interface{}
}
