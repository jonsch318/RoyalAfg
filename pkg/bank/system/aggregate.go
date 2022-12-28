package system

type IAggregate interface {
	GetId() string
	GetType() string
	GetVersion() int
	Apply(event IEvent[any], isNew bool)
	GetChanges() []IEvent[any]
	ClearChanges()
}

type IAggregateFactory[T IAggregate] interface {
	Create([]IEvent[any]) (T, error)
}
