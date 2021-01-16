package commands

type CreateBankAccount struct {
}

type Deposit struct {
	Amount int
	//Identification of Game
	GameID string
	//Identification in which round
	RoundID string
}

type Withdraw struct {
	Amount int
	//Identification of Game
	GameID string
	//Identification in which round
	RoundID string
}