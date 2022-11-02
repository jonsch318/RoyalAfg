package system

type IEvent[T any] interface {
	GetDat() T
	GetAggregatId() string
}

type Event[T any] struct {
	AggregatId string
	Data       T
}

func (e *Event[T]) GetDat() T {
	return e.Data
}

func (e *Event[T]) GetAggregatId() string {
	return e.AggregatId
}
