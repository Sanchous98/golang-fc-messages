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
	NoneReason            deviceStatusReason = "none"
	CloudRequestedReason  deviceStatusReason = "cloudRequested"
	ScheduledUpdateReason deviceStatusReason = "scheduledUpdate"
	StatusChangeReason    deviceStatusReason = "statusChange"
	ErrorDetectedReason   deviceStatusReason = "errorDetected"
)

type DeviceStatusRequest struct {
	TransactionId uint32
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
	ShortAddr        string             `json:"-"`
	ExtAddr          string             `json:"-"`
	Rssi             int                `json:"-"`
	TransactionId    uint32             `json:"-"`
	Reason           deviceStatusReason `json:"reason"`
	Time             int64              `json:"time"`
	Timezone         int                `json:"timezone"`
	BatteryLevel     int                `json:"batteryLevel"`
	BatteryLevelLoad int                `json:"batteryLevelLoad"`
	NetworkState     int                `json:"networkState"`
	AutoRequest      int                `json:"autoRequest"`
	LockSensor       *struct {
		Raw     byte `json:"raw"`
		Privacy byte `json:"privacy"`
		Handle  byte `json:"handle"`
		Key     byte `json:"key"`
	} `json:"lockSensor,omitempty"`
}

func (d *DeviceStatusResponse) UnmarshalJSON(bytes []byte) error {
	var e response

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
	case NoneReason, CloudRequestedReason, ScheduledUpdateReason, StatusChangeReason, ErrorDetectedReason:
	default:
		return InvalidDeviceStatusReason{d.Reason}
	}

	d.TransactionId = e.TransactionId
	d.ShortAddr = e.ShortAddr
	d.ExtAddr = e.ExtAddr
	d.Rssi = e.Rssi

	return nil
}

func (d *DeviceStatusResponse) MarshalJSON() ([]byte, error) {
	switch d.Reason {
	case NoneReason, CloudRequestedReason, ScheduledUpdateReason, StatusChangeReason, ErrorDetectedReason:
	default:
		return nil, &InvalidDeviceStatusReason{d.Reason}
	}

	type deviceStatusResponse DeviceStatusResponse

	var e response
	var err error

	e.EventType = DeviceStatusResponseEvent
	e.TransactionId = d.TransactionId
	e.ShortAddr = d.ShortAddr
	e.ExtAddr = d.ExtAddr
	e.Rssi = d.Rssi

	if e.Payload, err = json.Marshal((*deviceStatusResponse)(d)); err != nil {
		return nil, err
	}

	return json.Marshal(&e)
}
