package messages

import (
	"bitbucket.org/4suites/golang-fc-messages/values"
	"errors"
	"github.com/goccy/go-json"
)

type eventType string

func (e eventType) Error() error { return errors.New("invalid event type " + string(e)) }

type event struct {
	EventType     eventType       `json:"eventType"`
	Payload       json.RawMessage `json:"payload"`
	TransactionId int             `json:"transactionId"`
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

	extract, err := path.Extract(bytes)

	if err != nil {
		return err
	}

	if err = json.Unmarshal(extract[0], (*ev)(e)); err != nil {
		return err
	}

	return nil
}

type response struct {
	ShortAddr     string          `json:"short_addr"`
	ExtAddr       string          `json:"ext_addr"`
	Rssi          values.RSSI     `json:"rssi"`
	EventType     eventType       `json:"eventType"`
	Payload       json.RawMessage `json:"payload"`
	TransactionId int             `json:"transactionId"`
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

	extract, err := path.Extract(bytes)

	if err != nil {
		return err
	}

	if err = json.Unmarshal(extract[0], (*response)(e)); err != nil {
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
