package errorz

import "errors"

var (
	EventNotFound  = errors.New("event not found")
	EventsNotFound = errors.New("events not found")
)
