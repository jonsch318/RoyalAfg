package system

import "encoding/json"

type IEventParser[T any] interface {
	Parse(event T) (*IEvent[any], error)
}

type IEvent[T any] interface {
	GetData() T
	GetAggregatId() string
	GetType() string
	GetDataRaw() []byte
	FromParsedData(data []byte) error
	GetContentType() string
}

type Event[T any] struct {
	AggregatId string `json:"aggregatId"`
	Type       string `json:"type"`
	Data       T      `json:"data"`
}

func (e *Event[T]) GetContentType() string {
	return "application/json"
}

func (e *Event[T]) GetData() T {
	return e.Data
}

func (e *Event[T]) GetType() string {
	return e.Type
}

func (e *Event[T]) GetAggregatId() string {
	return e.AggregatId
}

func (e *Event[T]) GetDataRaw() []byte {
	data, _ := json.Marshal(e.Data)
	return data
}

func (e *Event[T]) FromParsedData(data []byte) error {
	return json.Unmarshal(data, &e.Data)
}
