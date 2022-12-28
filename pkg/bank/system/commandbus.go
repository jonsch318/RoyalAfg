package system

//CommandBus is the central command bus
type ICommandBus interface {
	Dispatch(command ICommand) error
	Subscribe(commandType string, handler ICommandHandler)
}

type ICommandHandler interface {
	Handle(command ICommand) error
}



type IEventHandler[T any] interface {
	Handle(event IEvent[T])
	HandleEventType(string) bool
}

type InternalCommandBus struct {
	handlers map[string][]ICommandHandler
}

func NewInternalCommandBus() *InternalCommandBus {
	return &InternalCommandBus{
		handlers: make(map[string][]ICommandHandler),
	}
}

//Dispatch a command to the command bus
func (cmdBus *InternalCommandBus) Dispatch(command ICommand) {
	handlers, ok := cmdBus.handlers[command.GetType()]
	if !ok {
		logging.Logger.Debugf("No handler for command %s found", command.GetType())
	}
	for _, h := range handlers {
		if err := h.Handle(command); err != nil {
			logging.Logger.Errorf("Error handling command %s: %s", command.GetType(), err.Error())
		}
	}
}

//Subscribe a command handler to a command type
func (cmdBus *InternalCommandBus) Subscribe(commandType string, handler ICommandHandler) {
	handlers, ok := cmdBus.handlers[commandType]
	if !ok {
		handlers = []ICommandHandler{}
	}
	handlers = append(handlers, handler)
	cmdBus.handlers[commandType] = handlers
