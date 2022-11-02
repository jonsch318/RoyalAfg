package system

import "github.com/JohnnyS318/RoyalAfgInGo/pkg/logging"

//CommandBus is the central command bus
type ICommandBus interface {
	Dispatch(command ICommand) error
	Subscribe(commandType string, handler ICommandHandler)
}

type ICommandHandler interface {
	Handle(command ICommand) error
}

type IEventBus interface {
	Publish(event IEvent[any])
	Subscribe(handler IEventHandler[any]) error
}

type IEventHandler[T any] interface {
	Handle(event IEvent[T])
	HandleEventType(string) bool
}

type CommandBus struct {
	handlers map[string][]ICommandHandler
}

func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[string][]ICommandHandler),
	}
}

//Dispatch a command to the command bus
func (cmdBus *CommandBus) Dispatch(command ICommand) {
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
func (cmdBus *CommandBus) Subscribe(commandType string, handler ICommandHandler) {
	handlers, ok := cmdBus.handlers[commandType]
	if !ok {
		handlers = []ICommandHandler{}
	}
	handlers = append(handlers, handler)
	cmdBus.handlers[commandType] = handlers
}

type EventBus struct {
	handlers []IEventHandler[any]
}

func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make([]IEventHandler[any], 0),
	}
}

//Publish an event to the event bus
func (evBus *EventBus) Publish(event IEvent[any]) {
	if evBus.handlers == nil {
		logging.Logger.Debugf("No handler for event %s found", event.GetType())
	}
	for _, h := range evBus.handlers {
		go func() {

			if h.HandleEventType(event.GetType()) {
				h.Handle(event)
			}
		}()
	}
}

func (evBus *EventBus) Subscribe(handler IEventHandler[any]) error {
	evBus.handlers = append(evBus.handlers, handler)
	return nil
}
