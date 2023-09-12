package helpers

import (
	"time"

	"github.com/Rhymond/go-money"
	"github.com/jonsch318/royalafg/services/bank/pkg/events"
)

type GeneralTransaction struct {
	ID       string
	Amount   *money.Money
	Withdraw bool
	GameId   string
	RoundId  string
	Time     time.Time
}

func ToGeneralTransactionParse(event map[string]interface{}) *GeneralTransaction {
	if event["type"] == "AccountCreated" {
		return nil
	}

	time, _ := time.Parse(time.RFC3339, event["time"].(string))
	transaction := &GeneralTransaction{
		ID:      event["id"].(string),
		Amount:  money.New(event["amount"].(int64), "EUR"),
		GameId:  event["gameId"].(string),
		RoundId: event["roundId"].(string),
		Time:    time,
	}

	if event["type"] == "Withdraw" {
		transaction.Withdraw = true
	} else if event["type"] == "Deposit" {
		transaction.Withdraw = false
	} else if event["type"] == "Backroll" {
		transaction.Withdraw = event["withdraw"].(bool)
	}

	return transaction
}

func ToGeneralTransaction(event any) *GeneralTransaction {
	switch ev := event.(type) {
	case *events.Withdrawn:
		return &GeneralTransaction{
			ID:       ev.ID,
			Amount:   ev.Amount,
			Withdraw: true,
			GameId:   ev.GameId,
			RoundId:  ev.RoundId,
			Time:     ev.Time,
		}
	case *events.Deposited:
		return &GeneralTransaction{
			ID:       ev.ID,
			Amount:   ev.Amount,
			Withdraw: false,
			GameId:   ev.GameId,
			RoundId:  ev.RoundId,
			Time:     ev.Time,
		}
	case *events.Backroll:
		return &GeneralTransaction{
			ID:       ev.ID,
			Amount:   ev.Amount,
			Withdraw: ev.Withdraw,
			GameId:   ev.GameId,
			RoundId:  ev.RoundId,
			Time:     ev.Time,
		}
	default:
		return nil
	}

}
