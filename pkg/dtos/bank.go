package dtos

type VerifyAmount struct {
	VerificationResult bool `json:"result"`
}

func NewVerifyAmount(res bool) *VerifyAmount {
	return &VerifyAmount{
		VerificationResult: res,
	}
}