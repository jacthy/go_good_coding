package main

import (
	"awesomeProject/test-eventhorizon/domain"
	"awesomeProject/test-eventhorizon/presentation"
	"context"
	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/eventbus/local"
	"time"

	//"github.com/looplab/eventhorizon/tracing"
	"log"
)

// Main test-eventhorizon的程序入口
func main() {
	var eventBus eh.EventBus
	eventBus = local.NewEventBus()
	aggEvent := AggEvent{}
	//eventBus = tracing.NewEventBus(eventBus)
	//if err := eventBus.AddHandler(context.Background(), eh.MatchAll{},
	//	eh.UseEventHandlerMiddleware(aggEvent,
	//		observer.NewMiddleware(observer.NamedGroup("global")),
	//	),
	//); err != nil {
	//	log.Fatal("could not add event logger: ", err)
	//}
	err := eventBus.AddHandler(context.Background(), eh.MatchAll{}, aggEvent)
	if err != nil {
		log.Fatal("could not add event logger: ", err)
	}
	//err = aggEvent.HandleEvent(context.Background(), eh.NewEvent(eh.EventType("test"), eh.EventData("{}"), time.Now()))
	//if err != nil {
	//	log.Fatal("could not add event logger: ", err)
	//}
	go func() {
		for err := range eventBus.Errors() {
			log.Print("eventbus:", err)
		}
	}()
	eh.RegisterEventData("test", func() eh.EventData {
		return &AggEventData{}
	})
	var data eh.EventData
	data = &AggEventData{Data: "data value"}

	err = eventBus.HandleEvent(context.Background(),
		eh.NewEvent(eh.EventType("test"), data, time.Now()))
	if err != nil {
		log.Fatal("could not add event logger: ", err)
	}
	domain.SetupDomain()
	presentation.InitRouter()
}

type AggEventData struct {
	Data string `json:"data"`
}

type AggEvent struct{}

// HandlerType implements the HandlerType method of the eventhorizon.EventHandler interface.
func (l AggEvent) HandlerType() eh.EventHandlerType {
	return "agg"
}

// HandleEvent implements the HandleEvent method of the EventHandler interface.
func (l AggEvent) HandleEvent(ctx context.Context, event eh.Event) error {
	log.Print("handle agg event")
	log.Printf("EVENT: %s", event)

	return nil
}
