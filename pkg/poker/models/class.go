package models

//Class is the model for a lobby class specified by a blind and a minumum and maximum buyin
type Class struct {
	//Min is the minimum buy in
	Min int `json:"min"`
	//Max is the maximum buy in
	Max int `json:"max"`
	//Blind is the small blind. Big Blind should be Blind*2
	Blind int `json:"blind"`
}

//NewClass creates a new lobby class
func NewClass(min, max, blind int) *Class {
	return &Class{
		Min:   min,
		Max:   max,
		Blind: blind,
	}
}
