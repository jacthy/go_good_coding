package domain

import (
	"github.com/google/uuid"
	eh "github.com/looplab/eventhorizon"
)

func SetupDomain() error {
	commands := []eh.CommandType{
		CreatCustom,
	}
	for _, cmdType := range commands {
		agg, err := eh.CreateAggregate(CustomAgg, uuid.Nil)
		if err != nil {
			return err
		}
		err = CommandBus.SetHandler(agg, cmdType)
		if err != nil {
			return err
		}
	}
	return nil
}
