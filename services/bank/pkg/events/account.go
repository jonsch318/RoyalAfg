package events

type AccountCreated struct {
	ID string
	Balance int
}

type Withdrawn struct {
	ID     string
	Amount int
}

type Deposited struct {
	ID     string
	Amount int
}