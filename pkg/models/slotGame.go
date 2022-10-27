package models

import (
	"github.com/Kamva/mgm"
	validation "github.com/go-ozzo/ozzo-validation"
)

type SlotGame struct {
	mgm.DefaultModel `bson:",inline"`
	Numbers          []int  `json:"numbers" bson:"numbers"`
	Win              int64  `json:"win" bson:"win"`
	Proof            string `json:"proof" bson:"proof"`
	Alpha            string `json:"alpha" bson:"alpha"`
	Beta             string `json:"beta" bson:"beta"`
}

func (game SlotGame) Validate() error {
	return validation.ValidateStruct(&game,
		validation.Field(&game.Numbers, validation.Required, validation.Length(3, 3)),
		validation.Field(&game.Win, validation.Required),
		validation.Field(&game.Proof, validation.Required),
		validation.Field(&game.Alpha, validation.Required),
		validation.Field(&game.Beta, validation.Required),
	)
}

func NewSlotGame(gameNumber int64, numbers []int, win int64, proof, alpha, beta string) *SlotGame {
	return &SlotGame{
		Numbers: numbers,
		Win:     win,
		Proof:   proof,
		Alpha:   alpha,
		Beta:    beta,
	}
}
