package messages

import (
	"github.com/goccy/go-json"
	"time"
)

const (
	GetNetworkInfoRequestEventType eventType = "getNwkInfoReq"
	UpdateNetworkStateEventType    eventType = "updateNetworkState"
	RemoveDeviceRequestEventType   eventType = "removeDeviceReq"
	RemoveDeviceResponseEventType  eventType = "removeDeviceRsp"
)

type networkAction string

const (
	NetworkOpenAction  networkAction = "open"
	NetworkCloseAction networkAction = "close"
)

type GetNetworkInfo struct {
	TransactionId int
}

func (g *GetNetworkInfo) UnmarshalJSON(bytes []byte) error {
	var e event

	if err := json.Unmarshal(bytes, &e); err != nil {
		return err
	}

	if e.EventType != GetNetworkInfoRequestEventType {
		return e.EventType.Error()
	}

	g.TransactionId = e.TransactionId

	return nil
}

func (g *GetNetworkInfo) MarshalJSON() ([]byte, error) {
	var e event

	e.EventType = GetNetworkInfoRequestEventType
	e.TransactionId = g.TransactionId

	return json.Marshal(&e)
}

type Device struct {
	Name         string   `json:"name"`
	Active       string   `json:"active"`
	ShortAddr    string   `json:"short_addr"`
	ExtAddr      string   `json:"ext_addr"`
	Topic        string   `json:"topic"`
	SmartObjects struct{} `json:"smart_objects"`
}

type GetNetworkInfoResponse struct {
	Name            string   `json:"name"`
	Channels        int      `json:"channels"`
	PanId           string   `json:"pan_id"`
	ShortAddr       string   `json:"short_addr"`
	ExtAddr         string   `json:"ext_addr"`
	SecurityEnabled int      `json:"security_enabled"`
	Mode            string   `json:"mode"`
	State           string   `json:"state"`
	Devices         []Device `json:"devices"`
}

type UpdateNetworkState struct {
	TransactionId int           `json:"-"`
	Action        networkAction `json:"action"`
	Duration      time.Duration `json:"duration"`
}

func (u *UpdateNetworkState) UnmarshalJSON(bytes []byte) error {
	var e event

	if err := json.Unmarshal(bytes, &e); err != nil {
		return err
	}

	if e.EventType != UpdateNetworkStateEventType {
		return e.EventType.Error()
	}

	type updateNetworkState UpdateNetworkState

	if err := json.Unmarshal(e.Payload, (*updateNetworkState)(u)); err != nil {
		return err
	}

	switch u.Action {
	case NetworkOpenAction, NetworkCloseAction:
	default:
		return InvalidNetworkAction{u.Action}
	}

	u.TransactionId = e.TransactionId

	return nil
}

func (u *UpdateNetworkState) MarshalJSON() ([]byte, error) {
	switch u.Action {
	case NetworkOpenAction, NetworkCloseAction:
	default:
		return nil, &InvalidNetworkAction{u.Action}
	}

	var e event
	var err error

	type updateNetworkState UpdateNetworkState

	e.EventType = UpdateNetworkStateEventType
	e.TransactionId = u.TransactionId

	if e.Payload, err = json.Marshal((*updateNetworkState)(u)); err != nil {
		return nil, err
	}

	return json.Marshal(&e)
}

type RemoveDeviceRequest struct {
	TransactionId int    `json:"-"`
	ExtAddress    string `json:"extAddress"`
}

func (r *RemoveDeviceRequest) UnmarshalJSON(bytes []byte) error {
	var e event

	if err := json.Unmarshal(bytes, &e); err != nil {
		return err
	}

	if e.EventType != RemoveDeviceRequestEventType {
		return e.EventType.Error()
	}

	type removeDeviceRequest RemoveDeviceRequest

	if err := json.Unmarshal(e.Payload, (*removeDeviceRequest)(r)); err != nil {
		return err
	}

	r.TransactionId = e.TransactionId

	return nil
}

func (r *RemoveDeviceRequest) MarshalJSON() ([]byte, error) {
	var e event
	var err error

	type removeDeviceRequest RemoveDeviceRequest

	e.EventType = RemoveDeviceRequestEventType
	e.TransactionId = r.TransactionId

	if e.Payload, err = json.Marshal((*removeDeviceRequest)(r)); err != nil {
		return nil, err
	}

	return json.Marshal(&e)
}

type RemoveDeviceResponse struct {
	ExtAddr          string `json:"ext_addr"`
	RemoveDeviceAddr string `json:"removeDeviceAddr,omitempty"`
	Error            string `json:"error,omitempty"`
}
