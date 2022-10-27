package dtos

// The DTO of a spin request
type SpinDTO struct {
	Number    uint32 `json:"number"`
	Prove     []byte `json:"prove"`
	WinAmount int    `json:"winAmount"`
}

type CryptoInfoDto struct {
	PublicKey string `json:"publicKey"`
}
