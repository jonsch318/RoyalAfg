package models

import (
	"fmt"
	"math/rand"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Colors of the card
const (
	Clubs    = 0
	Diamonds = 1
	Hearts   = 2
	Spades   = 3
)

//Card represents a play card which has a color and a value
//Values: [0,2];[1,3];[2,4];[3,5];[4,6];[5,7];[6,8];[7,9];[8,10];[9,J];[10,Q];[11,K];[12,A]
type Card struct {
	Color int `json:"color" mapstructure:"color"`
	Value int `json:"value" mapstructure:"value"`
}

//NewCard creates a card from the color and value given
func NewCard(color, value int) (*Card, error) {
	card := &Card{
		Color: color,
		Value: value,
	}

	err := card.Validate()
	if err != nil {
		return nil, err
	}

	return card, nil
}

//GetRandom picks a random card with an equal chance
func GetRandom() (*Card, error) {
	c := rand.Intn(4)
	v := rand.Intn(13)

	return NewCard(c, v)
}

//Validate validates the color and the value of the card.
func (card Card) Validate() error {
	return validation.ValidateStruct(&card,
		validation.Field(&card.Color, validation.Min(0), validation.Max(3)),
		validation.Field(&card.Value, validation.Min(0), validation.Max(12)))
}

func (card *Card) String() string {
	switch card.Color {
	case 0:
		return fmt.Sprintf("Card Clubs [%v]", card.Value)
	case 1:
		return fmt.Sprintf("Card Diamond [%v]", card.Value)
	case 2:
		return fmt.Sprintf("Card Hearts [%v]", card.Value)
	case 3:
		return fmt.Sprintf("Card Spades [%v]", card.Value)
	default:
		return "Not a card"
	}
}
