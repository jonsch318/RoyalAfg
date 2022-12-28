package events

import (
	"time"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/bank/system"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/payment"
	"github.com/Rhymond/go-money"
)

type AccountEventParser struct{}

func (a *AccountEventParser) Parse(event *esdb.RecordedEvent) (*system.IEvent[any], error) {
	return nil, nil
}

type AccountCreated struct {
	HolderId   string
	HolderName string
}

type AccountDeleted struct {
}

type Deposited struct {
	Amount      *money.Money
	GameId      string
	RoundId     string
	Time        time.Time
	PaymentInfo payment.PaymentInfo
}

// Withdrawal is the event for a withdrawal.
type Withdrawn struct {
	Amount      *money.Money
	GameId      string
	RoundId     string
	Time        time.Time
	PaymentInfo payment.PaymentInfo
}

// Rolled back an
type RolledBack struct {
	Amount                  *money.Money // Amount to rollback
	GameId                  string       // The game that the rollback is for if any
	RoundId                 string       // The round that was rolled back if any
	Time                    time.Time    // Time of the rollback
	Withdraw                bool         // true = rollback (withdraw), false = refund (deposit)
	Reason                  string       //Reason for the rollback
	RolledBackTransactionId string       // The transaction id of the transaction that was rolled back
	PaymentInfo             payment.PaymentInfo
}

type TransactionSend struct {
	Amount             *money.Money
	GameId             string
	RoundId            string
	Time               time.Time
	ReceivingUserId    string
	ReceivingAccountId string
	PaymentInfo        payment.PaymentInfo
}

type TransactionReceived struct {
	Amount           *money.Money
	GameId           string
	RoundId          string
	Time             time.Time
	SendingUserId    string
	SendingAccountId string
	PaymentInfo      payment.PaymentInfo
}
