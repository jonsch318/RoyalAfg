package payment

import "github.com/google/uuid"

type PaymentProvider string

const (
	GamePayment   PaymentProvider = "GamePayment"
	CryptoDeposit PaymentProvider = "CryptoDeposit"
	TestDeposit   PaymentProvider = "TestDeposit"
	RollbackOrder PaymentProvider = "RollbackOrder"
	Transaction   PaymentProvider = "Transaction"
)

// PaymentInfo is the payment information for a deposit. If any Information is not needed for an selected PaymentProvider leave it blank.
type PaymentInfo struct {
	TransactionId   string
	Provider        PaymentProvider
	TransactionHash string
	Details         map[string]interface{}
}

func NewBackrollOrder() *PaymentInfo {
	return &PaymentInfo{
		TransactionId: uuid.New().String(),
		Provider:      RollbackOrder,
	}
}
