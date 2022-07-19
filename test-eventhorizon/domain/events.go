package domain

import eh "github.com/looplab/eventhorizon"

const (
	CreateEvent = eh.EventType("creatCustom")
	UpdateEvent = eh.EventType("updateCustom")
)
