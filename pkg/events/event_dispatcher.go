package events

type EventDispatcher struct {
	events map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		events: map[string][]EventHandlerInterface{},
	}
}
