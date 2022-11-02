package system

type IEventParser[E any] interface {
	Parse(*E) (IEvent[any], error)
}

type IEvent[T any] interface {
	GetData() T
	GetAggregatId() string
}

type Event[T any] struct {
	AggregatId string
	Data       T
}

func (e *Event[T]) GetData() T {
	return e.Data
}

func (e *Event[T]) GetAggregatId() string {
	return e.AggregatId
}
