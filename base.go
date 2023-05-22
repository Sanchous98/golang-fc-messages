package messages

import (
	"github.com/goccy/go-json"
)

type eventType string

func (e eventType) Error() error { return InvalidEventType{e} }

type event struct {
	EventType     eventType       `json:"eventType"`
	Payload       json.RawMessage `json:"payload"`
	TransactionId uint32          `json:"transactionId"`
}

func (e *event) MarshalJSON() ([]byte, error) {
	type ev event

	if e.Payload == nil {
		e.Payload = []byte{'{', '}'}
	}

	return json.Marshal(map[string]*ev{"event": (*ev)(e)})
}

func (e *event) UnmarshalJSON(bytes []byte) error {
	type ev event
	path, err := json.CreatePath("$.event")

	if err != nil {
		return err
	}

	values, err := path.Extract(bytes)

	if err != nil {
		return err
	}

	if len(values) != 1 {
		return nil
	}

	if err = json.Unmarshal(values[0], (*ev)(e)); err != nil {
		return err
	}

	return nil
}

type response struct {
	ShortAddr     string          `json:"short_addr"`
	ExtAddr       string          `json:"ext_addr"`
	Rssi          int             `json:"rssi"`
	EventType     eventType       `json:"eventType"`
	Payload       json.RawMessage `json:"payload"`
	TransactionId uint32          `json:"transactionId"`
}

func (r *response) MarshalJSON() ([]byte, error) {
	type rsp response

	if r.Payload == nil {
		r.Payload = []byte{'{', '}'}
	}

	return json.Marshal((*rsp)(r))
}

type eventResponse response

func (e *eventResponse) UnmarshalJSON(bytes []byte) error {
	path, err := json.CreatePath("$.event")

	if err != nil {
		return err
	}

	values, err := path.Extract(bytes)

	if err != nil {
		return err
	}

	if len(values) != 1 {
		return nil
	}

	if err = json.Unmarshal(values[0], (*response)(e)); err != nil {
		return err
	}

	return nil
}

func (e *eventResponse) MarshalJSON() ([]byte, error) {
	if e.Payload == nil {
		e.Payload = []byte{'{', '}'}
	}

	return json.Marshal(map[string]*response{
		"event": (*response)(e),
	})
}
