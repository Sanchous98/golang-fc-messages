package messages

import (
	"github.com/goccy/go-json"
)

const (
	TransactionIdReq eventType = "transactionIdReq"
	TransactionIdRsp eventType = "transactionIdRsp"
)

const (
	TransactionActionRead  transactionIdAction = "read"
	TransactionActionReset transactionIdAction = "reset"
)

type transactionIdAction string

type TransactionIdAction struct {
	Action transactionIdAction `json:"transactionIdAction"`
}

func (t *TransactionIdAction) UnmarshalJSON(bytes []byte) error {
	var e event
	var err error

	if err = json.Unmarshal(bytes, &e); err != nil {
		return err
	}

	if e.EventType != TransactionIdReq {
		return e.EventType.Error()
	}

	type tIdAction TransactionIdAction

	if err = json.Unmarshal(e.Payload, (*tIdAction)(t)); err != nil {
		return err
	}

	switch t.Action {
	case TransactionActionRead, TransactionActionReset:
	default:
		return InvalidTransactionIdAction{t.Action}
	}

	return nil
}

func (t *TransactionIdAction) MarshalJSON() ([]byte, error) {
	switch t.Action {
	case TransactionActionRead, TransactionActionReset:
	default:
		return nil, &InvalidTransactionIdAction{t.Action}
	}

	type tIdAction TransactionIdAction

	var e event
	var err error

	e.EventType = TransactionIdReq

	if e.Payload, err = json.Marshal((*tIdAction)(t)); err != nil {
		return nil, err
	}

	return json.Marshal(&e)
}

type TransactionIdResponse struct {
	ShortAddr           string `json:"-"`
	ExtAddr             string `json:"-"`
	Rssi                int    `json:"-"`
	DeviceTransactionId int    `json:"deviceTransactionId"`
}

func (t *TransactionIdResponse) UnmarshalJSON(bytes []byte) error {
	type transactionIdResponse TransactionIdResponse

	var r response
	var err error

	if err = json.Unmarshal(bytes, &r); err != nil {
		return err
	}

	if r.EventType != TransactionIdRsp {
		return r.EventType.Error()
	}

	if err = json.Unmarshal(r.Payload, (*transactionIdResponse)(t)); err != nil {
		return err
	}

	t.ExtAddr = r.ExtAddr
	t.ShortAddr = r.ShortAddr
	t.Rssi = r.Rssi

	return nil
}

func (t *TransactionIdResponse) MarshalJSON() ([]byte, error) {
	type transactionIdResponse TransactionIdResponse

	var r response
	var err error

	r.EventType = TransactionIdRsp

	if r.Payload, err = json.Marshal((*transactionIdResponse)(t)); err != nil {
		return nil, err
	}

	return json.Marshal(&r)
}
