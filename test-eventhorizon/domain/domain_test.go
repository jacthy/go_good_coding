package domain

import (
	"encoding/json"
	"github.com/google/uuid"
	"testing"
)

func TestName(t *testing.T) {
	data := "{\"name\":\"first name\",\"age\":15}"
	cmd := new(Create)
	println(uuid.New().String())
	if err := json.Unmarshal([]byte(data), cmd); err != nil {
		println("could not decode Json" + err.Error())
		return
	}
	println("done")
}
