package models

import "encoding/json"

type Event struct {
	Name string      `json:"event"`
	Data interface{} `json:"data"`
}

func NewEvent(name string, data interface{}) *Event {
	return &Event{
		Name: name,
		Data: data,
	}
}

func NewEventFromRaw(raw []byte) (*Event, error) {
	event := new(Event)

	err := json.Unmarshal(raw, event)

	return event, err
}

func (e *Event) ToRaw() []byte {
	// Structs can nearly always be decoded back to json, so a error is unnecessery
	d, _ := json.Marshal(e)
	return d
}
