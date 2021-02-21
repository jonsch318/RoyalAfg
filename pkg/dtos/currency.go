package dtos

import (
	"github.com/Rhymond/go-money"
)

//CurrencyDto is the data transfer object for a monetary value
type CurrencyDto struct {
	Value int64 `json:"value"` //Value of the monetary value
	Currency string `json:"currency"` //Currency Code of the monetary value
}

//FromMoney converts between the used currency and precision library used to the currency dto
func FromMoney(money *money.Money) *CurrencyDto{
	return &CurrencyDto{
		Value:    money.Amount(),
		Currency: money.Currency().Code,
	}
}

//FromMoney converts between the dto to the used currency and precision library used.
func FromDTO(dto *CurrencyDto) *money.Money{
	return money.New(dto.Value, dto.Currency)
}