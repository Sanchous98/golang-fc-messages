package messages

import "github.com/goccy/go-json"

var eventPath *json.Path

func init() {
	var err error
	eventPath, err = json.CreatePath("$.event")

	if err != nil {
		panic(err)
	}
}

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

	values, err := eventPath.Extract(bytes)

	if err != nil {
		return err
	}

	if len(values) != 1 {
		return nil
	}

	return json.Unmarshal(values[0], (*ev)(e))
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
	values, err := eventPath.Extract(bytes)

	if err != nil {
		return err
	}

	if len(values) != 1 {
		return nil
	}

	return json.Unmarshal(values[0], (*response)(e))
}

func (e *eventResponse) MarshalJSON() ([]byte, error) {
	if e.Payload == nil {
		e.Payload = []byte{'{', '}'}
	}

	return json.Marshal(map[string]*response{
		"event": (*response)(e),
	})
}
