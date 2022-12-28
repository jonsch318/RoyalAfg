package system

import "github.com/JohnnyS318/RoyalAfgInGo/pkg/logging"

type IHandler interface {
	Handle(any)
}

type ISubcribable interface {
	SubscriptionName() comparable
}

type Bus[I comparable] struct {
	handlers map[I][]IHandler
}

func NewBus[I comparable]() *Bus[I] {
	return &Bus[I]{
		handlers: make(map[I][]IHandler),
	}
}

func (b *Bus[I]) Sub(to I, handler IHandler) {
	handlers, ok := b.handlers[to]
	if !ok {
		handlers = []IHandler{
			handler,
		}
		b.handlers[to] = handlers
	}
	handlers = append(handlers, handler)
	b.handlers[to] = handlers
}

func (b *Bus[I]) Pub(to I, event any) {
	handlers, ok := b.handlers[to]
	if !ok {
		logging.Logger.Debugf("No handler for event %s found", to)
	}
	for _, h := range handlers {
		h.Handle(to)
	}
}

}

type InternalEventBus struct {
	handlers []IEventHandler[any]
}

func NewInternalEventBus() *InternalEventBus {
	return &InternalEventBus{
		handlers: make([]IEventHandler[any], 0),
	}
}

//Publish an event to the event bus
func (evBus *InternalEventBus) Publish(event IEvent[any]) {
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

func (evBus *InternalEventBus) Subscribe(handler IEventHandler[any]) error {
	evBus.handlers = append(evBus.handlers, handler)
	return nil
}
