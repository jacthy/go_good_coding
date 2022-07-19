package domain

import (
	"context"
	"fmt"
	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/aggregatestore/events"
	"github.com/looplab/eventhorizon/uuid"
	"time"
)

func init() {
	eh.RegisterAggregate(func(id uuid.UUID) eh.Aggregate {
		return &CustomAggregate{
			AggregateBase: events.NewAggregateBase(CustomAgg, id),
		}
	})
}

const CustomAgg eh.AggregateType = "customAgg"

var _ eh.Aggregate = new(CustomAggregate)

type CustomAggregate struct {
	*events.AggregateBase
	customEntity
}

//func NewCustomAggregate() *CustomAggregate {
//	return &CustomAggregate{
//		AggregateBase: events.NewAggregateBase(CustomAgg, uuid.New()),
//	}
//}

func (c *CustomAggregate) HandleCommand(ctx context.Context, command eh.Command) error {
	println("CustomAggregate HandleCommand")
	switch command.(type) {
	case *Create:
		c.AppendEvent(CreateEvent, ctx, time.Now())
	case *Update:
		c.AppendEvent(UpdateEvent, ctx, time.Now())
	}
	return nil
}

func (c *CustomAggregate) ApplyEvent(ctx context.Context, event eh.Event) (err error) {
	fmt.Printf("CustomAggregate event, eventType:%v\n,event: %+v\n", event.EventType(), event)
	switch event.EventType() {
	case CreateEvent:
		data := event.Data().(CustomModel)
		c.Create(data)
	case UpdateEvent:
		data := event.Data().(CustomModel)
		c.Update(data)
	}
	return
}
