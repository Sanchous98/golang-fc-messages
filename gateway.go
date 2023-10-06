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

const (
	NetworkOpenAction  networkAction = "open"
	NetworkCloseAction networkAction = "close"
)

type networkAction string

func (a *networkAction) UnmarshalJSON(bytes []byte) (err error) {
	defer func() {
		if a != nil {
			switch *a {
			case NetworkOpenAction, NetworkCloseAction:
			default:
				err = InvalidNetworkAction{*a}
			}
		}
	}()

	err = json.Unmarshal(bytes, (*string)(a))
	return
}

func (a *networkAction) MarshalJSON() ([]byte, error) {
	if a != nil {
		switch *a {
		case NetworkOpenAction, NetworkCloseAction:
		default:
			return nil, InvalidNetworkAction{*a}
		}
	}

	return json.Marshal((*string)(a))
}

type GetNetworkInfo struct {
	TransactionId uint32
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
	FwVersion       string   `json:"fw_version"`
	Devices         []Device `json:"devices"`
}

type UpdateNetworkState struct {
	TransactionId uint32        `json:"-"`
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

	u.TransactionId = e.TransactionId

	return nil
}

func (u *UpdateNetworkState) MarshalJSON() ([]byte, error) {
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
	TransactionId uint32 `json:"-"`
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
