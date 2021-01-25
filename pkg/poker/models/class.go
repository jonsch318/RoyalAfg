package models

type Class struct {
	Min   int `json:"min"`
	Max   int `json:"max"`
	Blind int `json:"blind"`
}

func NewClass(min, max, blind int) *Class {
	return &Class{
		Min:   min,
		Max:   max,
		Blind: blind,
	}
}
