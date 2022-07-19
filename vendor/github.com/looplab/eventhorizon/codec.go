// Copyright (c) 2021 - The Event Horizon authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package eventhorizon

import "context"

// EventCodec is a codec for marshaling and unmarshaling events to and from bytes.
type EventCodec interface {
	// MarshalEvent marshals an event and the supported parts of context into bytes.
	MarshalEvent(context.Context, Event) ([]byte, error)
	// UnmarshalEvent unmarshals an event and supported parts of context from bytes.
	UnmarshalEvent(context.Context, []byte) (Event, context.Context, error)
}

// CommandCodec is a codec for marshaling and unmarshaling commands to and from bytes.
type CommandCodec interface {
	// MarshalCommand marshals a command and the supported parts of context into bytes.
	MarshalCommand(context.Context, Command) ([]byte, error)
	// UnmarshalCommand unmarshals a command and supported parts of context from bytes.
	UnmarshalCommand(context.Context, []byte) (Command, context.Context, error)
}
