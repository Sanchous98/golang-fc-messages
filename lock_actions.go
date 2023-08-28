package messages

import "github.com/goccy/go-json"

const (
	LockActionOpenEventType      eventType = "lockActionOpen"
	LockActionCloseEventType     eventType = "lockActionClose"
	LockActionAutoEventType      eventType = "lockActionAuto"
	LockActionResponseEventType  eventType = "lockActionResponse"
	LockOfflineResponseEventType eventType = "lockOfflineResponse"
)

const (
	NoneLockStatus                   lockStatus = "none"
	ExtRelayStateLockStatus          lockStatus = "extRelayState"
	LockOpenedLockStatus             lockStatus = "lockOpened"
	LockClosedLockStatus             lockStatus = "lockClosed"
	DriverOnLockStatus               lockStatus = "driverOn"
	ErrorLockAlreadyOpenLockStatus   lockStatus = "errorLockAlreadyOpen"
	ErrorLockAlreadyClosedLockStatus lockStatus = "errorLockAlreadyClosed"
	ErrorDriverEnabledLockStatus     lockStatus = "errorDriverEnabled"
	DeviceTypeUnknownLockStatus      lockStatus = "deviceTypeUnknown"
	OpenTimeoutLockStatus            lockStatus = "openTimeoutError"
)

var lockActionsPath *json.Path

func init() {
	var err error
	lockActionsPath, err = json.CreatePath("$.payload.lockActionStatus")

	if err != nil {
		panic(err)
	}
}

type lockStatus string

func (s *lockStatus) UnmarshalJSON(bytes []byte) (err error) {
	defer func() {
		if s != nil {
			switch *s {
			case NoneLockStatus, ExtRelayStateLockStatus, LockOpenedLockStatus, LockClosedLockStatus, DriverOnLockStatus, ErrorLockAlreadyOpenLockStatus, ErrorLockAlreadyClosedLockStatus, ErrorDriverEnabledLockStatus, DeviceTypeUnknownLockStatus:
			default:
				err = InvalidLockStatus{*s}
			}
		}
	}()

	err = json.Unmarshal(bytes, (*string)(s))
	return
}

func (s *lockStatus) MarshalJSON() ([]byte, error) {
	if s != nil {
		switch *s {
		case NoneLockStatus, ExtRelayStateLockStatus, LockOpenedLockStatus, LockClosedLockStatus, DriverOnLockStatus, ErrorLockAlreadyOpenLockStatus, ErrorLockAlreadyClosedLockStatus, ErrorDriverEnabledLockStatus, DeviceTypeUnknownLockStatus:
		default:
			return nil, InvalidLockStatus{*s}
		}
	}

	return json.Marshal((*string)(s))
}

type LockAuto struct {
	TransactionId uint32 `json:"-"`
	RecloseDelay  uint   `json:"recloseDelay"`
	ChannelIds    []int  `json:"channelIds,omitempty"`
}

func (l *LockAuto) UnmarshalJSON(bytes []byte) error {
	var e event

	if err := json.Unmarshal(bytes, &e); err != nil {
		return err
	}

	if e.EventType != LockActionAutoEventType {
		return e.EventType.Error()
	}

	type lockAuto LockAuto

	if err := json.Unmarshal(e.Payload, (*lockAuto)(l)); err != nil {
		return err
	}

	l.TransactionId = e.TransactionId

	return nil
}

func (l *LockAuto) MarshalJSON() ([]byte, error) {
	type lockAuto LockAuto

	var e event
	var err error

	e.EventType = LockActionAutoEventType
	e.TransactionId = l.TransactionId

	if e.Payload, err = json.Marshal((*lockAuto)(l)); err != nil {
		return nil, err
	}

	return json.Marshal(&e)
}

type LockResponse struct {
	ShortAddr        string     `json:"-"`
	ExtAddr          string     `json:"-"`
	Rssi             int        `json:"-"`
	TransactionId    uint32     `json:"-"`
	LockActionStatus lockStatus `json:"lockActionStatus"`
	ChannelIds       []int      `json:"channelIds,omitempty"`
}

func (l *LockResponse) UnmarshalJSON(bytes []byte) error {
	var e response

	if err := json.Unmarshal(bytes, &e); err != nil {
		return err
	}

	if e.EventType != LockActionResponseEventType {
		return e.EventType.Error()
	}

	type lockResponse LockResponse

	if err := json.Unmarshal(e.Payload, (*lockResponse)(l)); err != nil {
		return err
	}

	l.TransactionId = e.TransactionId
	l.ShortAddr = e.ShortAddr
	l.ExtAddr = e.ExtAddr
	l.Rssi = e.Rssi

	return nil
}

func (l *LockResponse) MarshalJSON() ([]byte, error) {
	type lockResponse LockResponse

	var e event
	var err error

	e.EventType = LockActionResponseEventType
	e.TransactionId = l.TransactionId

	if e.Payload, err = json.Marshal((*lockResponse)(l)); err != nil {
		return nil, err
	}

	return json.Marshal(&e)
}

type LockClose struct {
	TransactionId uint32
}

func (l *LockClose) UnmarshalJSON(bytes []byte) error {
	var e event

	if err := json.Unmarshal(bytes, &e); err != nil {
		return err
	}

	if e.EventType != LockActionAutoEventType {
		return e.EventType.Error()
	}

	l.TransactionId = e.TransactionId

	return nil
}

func (l *LockClose) MarshalJSON() ([]byte, error) {
	var e event

	e.EventType = LockActionCloseEventType
	e.TransactionId = l.TransactionId

	return json.Marshal(&e)
}

type LockOpen struct {
	TransactionId uint32 `json:"-"`
	ChannelIds    []int  `json:"channelIds,omitempty"`
}

func (l *LockOpen) UnmarshalJSON(bytes []byte) error {
	var e event

	if err := json.Unmarshal(bytes, &e); err != nil {
		return err
	}

	if e.EventType != LockActionOpenEventType {
		return e.EventType.Error()
	}

	type lockOpen LockOpen

	if err := json.Unmarshal(e.Payload, (*lockOpen)(l)); err != nil {
		return err
	}

	l.TransactionId = e.TransactionId

	return nil
}

func (l *LockOpen) MarshalJSON() ([]byte, error) {
	type lockOpen LockOpen

	var e event
	var err error

	e.EventType = LockActionOpenEventType
	e.TransactionId = l.TransactionId

	if e.Payload, err = json.Marshal((*lockOpen)(l)); err != nil {
		return nil, err
	}

	return json.Marshal(&e)
}

type LockOffline struct {
	TransactionId uint32
}

func (l *LockOffline) UnmarshalJSON(bytes []byte) error {
	var e event

	if err := json.Unmarshal(bytes, &e); err != nil {
		return err
	}

	if e.EventType != LockOfflineResponseEventType {
		return e.EventType.Error()
	}

	var status lockStatus
	if err := lockActionsPath.Unmarshal(e.Payload, &status); err != nil {
		return err
	}

	if status != OpenTimeoutLockStatus {
		return ExpectedOfflineTimeoutError{status}
	}

	l.TransactionId = e.TransactionId

	return nil
}

func (l *LockOffline) MarshalJSON() ([]byte, error) {
	var e event

	e.EventType = LockOfflineResponseEventType
	e.TransactionId = l.TransactionId
	e.Payload, _ = json.Marshal(map[string]lockStatus{"lockActionStatus": OpenTimeoutLockStatus})

	return json.Marshal(&e)
}
