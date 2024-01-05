package events

import "fmt"

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: map[string][]EventHandlerInterface{},
	}
}

func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {

	if ed.Has(eventName, handler) {
		return fmt.Errorf("handler is already present on the event %s", eventName)
	}

	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

func (ed *EventDispatcher) Has(eventName string, eventHandler EventHandlerInterface) bool {

	if _, ok := ed.handlers[eventName]; ok {

		for _, handler := range ed.handlers[eventName] {
			return handler == eventHandler
		}
	}

	return false
}

func (ed *EventDispatcher) Dispatch(event EventInterface) {

	if _, ok := ed.handlers[event.GetName()]; ok {

		for _, handler := range ed.handlers[event.GetName()] {
			handler.Handle(event)
		}
	}
}

func (ed *EventDispatcher) Remove(eventName string, eventHandler EventHandlerInterface) error {
	return nil
}

func (ed *EventDispatcher) Clear() error {

	ed.handlers = make(map[string][]EventHandlerInterface)
	return nil
}
