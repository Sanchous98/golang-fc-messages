package messages

import (
	"github.com/goccy/go-json"
)

const (
	SerialConnectionRequestEventType  eventType = "serialConnectionReq"
	SerialConnectionResponseEventType eventType = "serialConnectionRsp"
)

const (
	SerialConnectionActionStart serialConnectionAction = "start"
	SerialConnectionActionReset serialConnectionAction = "reset"
)

type serialConnectionAction string

type SerialConnectionRequest struct {
	TransactionId int                    `json:"-"`
	Action        serialConnectionAction `json:"transactionIdAction"`
}

func (s *SerialConnectionRequest) UnmarshalJSON(bytes []byte) error {
	var e event
	var err error

	if err = json.Unmarshal(bytes, &e); err != nil {
		return err
	}

	if e.EventType != SerialConnectionRequestEventType {
		return e.EventType.Error()
	}

	type serialConnectionRequest SerialConnectionRequest

	if err = json.Unmarshal(e.Payload, (*serialConnectionRequest)(s)); err != nil {
		return err
	}

	switch s.Action {
	case SerialConnectionActionStart, SerialConnectionActionReset:
	default:
		return InvalidSerialConnectionAction{s.Action}
	}

	s.TransactionId = e.TransactionId

	return nil
}

func (s *SerialConnectionRequest) MarshalJSON() ([]byte, error) {
	switch s.Action {
	case SerialConnectionActionStart, SerialConnectionActionReset:
	default:
		return nil, &InvalidSerialConnectionAction{s.Action}
	}

	type serialConnectionRequest SerialConnectionRequest

	var e event
	var err error

	e.EventType = SerialConnectionRequestEventType
	e.TransactionId = s.TransactionId

	if e.Payload, err = json.Marshal((*serialConnectionRequest)(s)); err != nil {
		return nil, err
	}

	return json.Marshal(&e)
}

type SerialConnectionResponse struct {
	ShortAddr string `json:"-"`
	ExtAddr   string `json:"-"`
	Rssi      int    `json:"-"`
	Status    int    `json:"status"`
}

func (s *SerialConnectionResponse) UnmarshalJSON(bytes []byte) error {
	var r response

	if err := json.Unmarshal(bytes, &r); err != nil {
		return err
	}

	if r.EventType != SerialConnectionResponseEventType {
		return r.EventType.Error()
	}

	type serialConnectionResponse SerialConnectionResponse

	if err := json.Unmarshal(r.Payload, (*serialConnectionResponse)(s)); err != nil {
		return err
	}

	s.ShortAddr = r.ShortAddr
	s.ExtAddr = r.ExtAddr
	s.Rssi = r.Rssi

	return nil
}

func (s *SerialConnectionResponse) MarshalJSON() ([]byte, error) {
	var r response
	var err error

	type serialConnectionResponse SerialConnectionResponse

	r.EventType = SerialConnectionResponseEventType
	r.Rssi = s.Rssi
	r.ShortAddr = s.ShortAddr
	r.ExtAddr = s.ExtAddr

	if r.Payload, err = json.Marshal((*serialConnectionResponse)(s)); err != nil {
		return nil, err
	}

	return json.Marshal(&r)
}
