package repositories

import (
	ycq "github.com/jetbasrawi/go.cqrs"

	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/domain/aggregates"
)

type AccountRepository interface {
	Load(string, string) (*aggregates.Account, error)
	Save(ycq.AggregateRoot, *int) error
}
