package messages

import (
	"github.com/goccy/go-json"
)

const (
	DeviceStatusRequestEvent  eventType = "deviceStatusReq"
	DeviceStatusResponseEvent eventType = "deviceStatusRsp"
)

type deviceStatusReason string

const (
	CloudRequestedReason  deviceStatusReason = "cloudRequested"
	ScheduledUpdateReason deviceStatusReason = "scheduledUpdate"
)

type DeviceStatusRequest struct {
	TransactionId int
}

func (d *DeviceStatusRequest) UnmarshalJSON(bytes []byte) error {
	var e event

	if err := json.Unmarshal(bytes, &e); err != nil {
		return err
	}

	if e.EventType != DeviceStatusRequestEvent {
		return e.EventType.Error()
	}

	d.TransactionId = e.TransactionId

	return nil
}

func (d *DeviceStatusRequest) MarshalJSON() ([]byte, error) {
	var e event

	e.EventType = DeviceStatusRequestEvent
	e.TransactionId = d.TransactionId

	return json.Marshal(&e)
}

type DeviceStatusResponse struct {
	TransactionId    int                `json:"-"`
	Reason           deviceStatusReason `json:"reason"`
	Time             int                `json:"time"`
	BatteryLevel     int                `json:"batteryLevel"`
	BatteryLevelLoad int                `json:"batteryLevelLoad"`
	NetworkState     int                `json:"networkState"`
	AutoRequest      int                `json:"autoRequest"`
}

func (d *DeviceStatusResponse) UnmarshalJSON(bytes []byte) error {
	var e event

	if err := json.Unmarshal(bytes, &e); err != nil {
		return err
	}

	if e.EventType != DeviceStatusResponseEvent {
		return e.EventType.Error()
	}

	type deviceStatusResponse DeviceStatusResponse

	if err := json.Unmarshal(e.Payload, (*deviceStatusResponse)(d)); err != nil {
		return err
	}

	switch d.Reason {
	case CloudRequestedReason, ScheduledUpdateReason:
	default:
		return invalidDeviceStatusReason(d.Reason)
	}

	d.TransactionId = e.TransactionId

	return nil
}

func (d *DeviceStatusResponse) MarshalJSON() ([]byte, error) {
	switch d.Reason {
	case CloudRequestedReason, ScheduledUpdateReason:
	default:
		return nil, invalidDeviceStatusReason(d.Reason)
	}

	type deviceStatusResponse DeviceStatusResponse

	var e event
	var err error

	e.EventType = DeviceStatusResponseEvent
	e.TransactionId = d.TransactionId

	if e.Payload, err = json.Marshal((*deviceStatusResponse)(d)); err != nil {
		return nil, err
	}

	return json.Marshal(&e)
}
