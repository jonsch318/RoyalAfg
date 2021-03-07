package money

import (
	"github.com/Rhymond/go-money"
)

//ConvertToIMoney converts the integer money used in communication with clients (e.g. 200 => 2€) to the internal money library go-money
func ConvertToIMoney(amount int) *money.Money {
	//we only use the euro currency for simplicity.
	//Should be changed in production or be converted on Deposit and Withdrawal of the bank service
	return money.New(int64(amount), Currency)
}

func Zero() *money.Money{
	return ConvertToIMoney(0)
}