package models

import (
	"github.com/Kamva/mgm"
	validation "github.com/go-ozzo/ozzo-validation"
)

type SlotGame struct {
	mgm.DefaultModel `bson:",inline"`
	ID               string `json:"id" bson:"id"`
	Time             int64  `json:"time" bson:"time"`
	Numbers          []uint `json:"numbers" bson:"numbers"`
	Win              int64  `json:"win" bson:"win"`
	Proof            string `json:"proof" bson:"proof"`
	Alpha            string `json:"alpha" bson:"alpha"`
	Beta             string `json:"beta" bson:"beta"`
}

func (game SlotGame) Validate() error {
	return validation.ValidateStruct(&game,
		validation.Field(&game.Numbers, validation.Required, validation.Length(4, 4)),
		validation.Field(&game.Win, validation.Required),
		validation.Field(&game.Proof, validation.Required),
		validation.Field(&game.Alpha, validation.Required),
		validation.Field(&game.Beta, validation.Required),
		validation.Field(&game.Time, validation.Required),
		validation.Field(&game.ID, validation.Required),
	)
}

func NewSlotGame(gameId string, numbers []uint, win int64, proof, alpha, beta string, time int64) *SlotGame {
	return &SlotGame{
		Numbers: numbers,
		Win:     win,
		Proof:   proof,
		Alpha:   alpha,
		Beta:    beta,
		Time:    time,
	}
}
