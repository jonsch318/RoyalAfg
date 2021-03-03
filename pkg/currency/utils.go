package currency

import (
	"github.com/Rhymond/go-money"
)

const Code = "EUR"

//Contains helper functions and handy objects for the money library used

//Zero returns a zero monetary value in the default currency (EUR)
func Zero() *money.Money{
	return money.New(0, Code)
}

//ZeroWC returns a zero monetary value in the specified currency code
func ZeroWC(code string) *money.Money{
	return money.New(0, code)
}