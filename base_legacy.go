package messages

import (
	"errors"
	"github.com/goccy/go-json"
)

func (e legacyEvent) Error(EventType string) error {
	return errors.New("invalid event type for legacy event" + EventType)
}

type clientToken string

type legacyEvent struct {
	State       json.RawMessage `json:"state,omitempty"`
	ClientToken clientToken     `json:"clientToken,omitempty"`
}

func (e *legacyEvent) MarshalJSON() ([]byte, error) {
	type le legacyEvent

	if e.State == nil {
		e.State = []byte{'{', '}'}
	}

	return json.Marshal((*le)(e))
}
