package repositories

import (
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/domain/aggregates"
	ycq "github.com/jetbasrawi/go.cqrs"
)

type AccountRepository interface {
	Load(string, string) (*aggregates.Account, error)
	Save(ycq.AggregateRoot, *int) error
}


