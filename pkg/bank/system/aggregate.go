package system

type IAggregate interface {
	GetId() string
	GetType() string
	GetVersion() int
	Apply(event IEvent[any], isNew bool)
	GetChanges() []IEvent[any]
	ClearChanges()
}

type ICreatableAggregate interface {
	GetId() string
	GetType() string
	GetVersion() int
	Apply(event IEvent[any], isNew bool)
	GetChanges() []IEvent[any]
	ClearChanges()
	CreateFromEvents(events []IEvent[any])
}
