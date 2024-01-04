package events

type EventDispatcherInterface interface {
	Register(eventName string, eventHandler EventHandlerInterface) error
	Dispatch(event EventInterface)
	Remove(eventName string, eventHandler EventHandlerInterface) error
	Has(eventName string, eventHandler EventHandlerInterface) bool
	Clear() error
}
