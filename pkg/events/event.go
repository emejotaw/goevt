package events

import "time"

type EventInterface interface {
	GetName() string
	GetPayload() interface{}
	GetDateTime() time.Time
}
