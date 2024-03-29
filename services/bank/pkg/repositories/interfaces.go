package repositories

import (
	ycq "github.com/jetbasrawi/go.cqrs"

	"github.com/jonsch318/royalafg/services/bank/pkg/aggregates"
)

// AccountRepository is the interface for a repository for bank accounts
type AccountRepository interface {
	Load(string, string) (*aggregates.Account, error)
	Save(ycq.AggregateRoot, *int) error
}
