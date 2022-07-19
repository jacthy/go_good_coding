package domain

import (
	"github.com/google/uuid"
	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/commandhandler/bus"
)

func init() {
	eh.RegisterCommand(func() eh.Command { return &Create{ID: uuid.New()} })
	eh.RegisterCommand(func() eh.Command { return &Update{ID: uuid.New()} })
}

var CommandBus = bus.NewCommandHandler()

const (
	CreatCustom  = eh.CommandType("createCustom")
	UpdateCustom = eh.CommandType("updCustom")
)

type Create struct {
	ID   uuid.UUID `json:"uuid,omitempty"`
	Id   int       `json:"id,omitempty"`
	Name string    `json:"name"`
	Age  int       `json:"age"`
}

func (c *Create) AggregateID() uuid.UUID          { return c.ID }
func (c *Create) CommandType() eh.CommandType     { return CreatCustom }
func (c *Create) AggregateType() eh.AggregateType { return CustomAgg }

type Update struct {
	ID   uuid.UUID `json:"uuid,omitempty"`
	Id   int       `json:"id,omitempty"`
	Name string    `json:"name"`
	Age  int       `json:"age"`
}

func (c *Update) AggregateID() uuid.UUID          { return c.ID }
func (c *Update) CommandType() eh.CommandType     { return UpdateCustom }
func (c *Update) AggregateType() eh.AggregateType { return CustomAgg }
