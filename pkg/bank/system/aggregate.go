package system

type IAggregate interface {
	GetId() string
	GetVersion() int
	Apply(event IEvent[any], isNew bool)
	GetChanges() []IEvent[any]
	CleearChanges()
}
