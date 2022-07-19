package application

import (
	"awesomeProject/test-eventhorizon/domain"
	"context"
	"encoding/json"
	"fmt"
	eh "github.com/looplab/eventhorizon"
	"log"
	"sync"
)

var once sync.Once

func init() {
	once.Do(func() {
		CommonHandler = eh.UseCommandHandlerMiddleware(domain.CommandBus, CommandLogger)
	})
}

type Result struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`  // message
	Data    interface{} `json:"data"` // data object
}

func HandCustomCommand(data []byte, command string) (result Result) {
	cmd, err := eh.CreateCommand(eh.CommandType(command))
	if err != nil {
		result.Msg = "could not create command: " + err.Error()
		return
	}
	fmt.Printf("data:%s\ncmd:%+v\n", data, cmd)
	if err := json.Unmarshal(data, cmd); err != nil {
		result.Msg = "could not decode Json" + err.Error()
		return
	}
	err = CommonHandler.HandleCommand(context.Background(), cmd)
	if err != nil {
		result.Msg = "could not execute command" + err.Error()
		return
	}
	return
}

var CommonHandler eh.CommandHandler

func CommandLogger(h eh.CommandHandler) eh.CommandHandler {
	return eh.CommandHandlerFunc(func(ctx context.Context, cmd eh.Command) error {
		log.Printf("CMD: %#v", cmd)

		return h.HandleCommand(ctx, cmd)
	})
}
