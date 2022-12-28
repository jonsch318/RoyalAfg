package system

type System struct {
	EventBus     IEventBus
	CommandBus   ICommandBus
	ReadModels   []IReadModel           //ReadModels for queries
	Repositories []IAggregateRepository //Repositories for commands
}

func InitEventSourcingSystem() *System {
	//TODO: EventBus

	//TODO: CommandBus

	//TODO:
}

func (s *System) AddCommandHandler() {

}

func (s *System) AddIAggregateRepository() {

}
