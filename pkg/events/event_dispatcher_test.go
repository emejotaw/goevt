package events

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name      string
	Payload   interface{}
	CreatedAt time.Time
}

func NewTestEvent(name string, payload interface{}) *TestEvent {
	return &TestEvent{
		Name:      name,
		Payload:   payload,
		CreatedAt: time.Now(),
	}
}

func (t *TestEvent) GetName() string {
	return t.Name
}
func (t *TestEvent) GetPayload() interface{} {
	return t.Payload
}
func (t *TestEvent) GetDateTime() time.Time {
	return t.CreatedAt
}

type TestHandler struct {
}

func (th *TestHandler) Handle(event EventInterface) error {
	return nil
}

type TestCase struct {
	event         EventInterface
	handler       EventHandlerInterface
	expectedError error
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventInterface) error {
	m.Called(event)
	return nil
}

type EventDispatcherTestSuite struct {
	suite.Suite
	eventA          EventInterface
	eventB          EventInterface
	eventC          EventInterface
	handlerA        EventHandlerInterface
	handlerB        EventHandlerInterface
	handlerC        EventHandlerInterface
	eventDispacther EventDispatcherInterface
	tests           []TestCase
}

func TestRun(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}

func (e *EventDispatcherTestSuite) SetupTest() {

	handlerA := new(TestHandler)
	handlerB := new(TestHandler)
	handlerC := new(TestHandler)
	tests := []TestCase{
		{event: NewTestEvent("event-a", "message 1"), handler: new(TestHandler), expectedError: nil},
		{event: NewTestEvent("event-b", "message 2"), handler: new(TestHandler), expectedError: nil},
		{event: NewTestEvent("event-c", "message 3"), handler: new(TestHandler), expectedError: nil},
	}
	e.eventA = &TestEvent{Name: "event-a"}
	e.eventB = &TestEvent{Name: "event-b"}
	e.eventC = &TestEvent{Name: "event-c"}
	e.handlerA = handlerA
	e.handlerB = handlerB
	e.handlerC = handlerC
	e.eventDispacther = NewEventDispatcher()
	e.tests = tests
}

func (e *EventDispatcherTestSuite) TestRegister() {

	eventDispatcher := e.eventDispacther
	tests := append(e.tests, TestCase{event: e.eventA, handler: e.handlerA, expectedError: errors.New("handler is already present on the event event-a")})

	for _, test := range tests {

		err := eventDispatcher.Register(test.event.GetName(), test.handler)
		e.Equal(test.expectedError, err)
	}
}

func (e *EventDispatcherTestSuite) TestDispatch() {

	eventDispatcher := e.eventDispacther

	for _, test := range e.tests {

		mockHandler := &MockHandler{}
		mockHandler.On("Handle", test.event)
		err := eventDispatcher.Register(test.event.GetName(), mockHandler)
		e.Nil(err)

		eventDispatcher.Dispatch(test.event)

		mockHandler.AssertExpectations(e.T())
		mockHandler.AssertNumberOfCalls(e.T(), "Handle", 1)
	}
}

func (e *EventDispatcherTestSuite) TestRemove() {

	eventDispatcher := e.eventDispacther

	for _, test := range e.tests {
		err := eventDispatcher.Register(test.event.GetName(), test.handler)
		e.Nil(err)
	}

	tests := append(e.tests, TestCase{event: &TestEvent{Name: "a"}, handler: new(TestHandler), expectedError: errors.New("event does not exists")})
	for _, test := range tests {
		err := eventDispatcher.Remove(test.event.GetName(), test.handler)
		e.Equal(test.expectedError, err)
	}
}

func (e *EventDispatcherTestSuite) TestClear() {

	eventDispatcher := e.eventDispacther
	for _, test := range e.tests {

		err := eventDispatcher.Register(test.event.GetName(), test.handler)
		e.Equal(test.expectedError, err)
	}

	err := eventDispatcher.Clear()
	e.Nil(err)
}
