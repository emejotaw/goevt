package events

type EventHandlerInterface interface {
	Handle(event EventInterface) error
}
