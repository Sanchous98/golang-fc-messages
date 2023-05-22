package messages

import "github.com/goccy/go-json"

const (
	FwVersionRequestEventType        eventType = "fwVersionReq"
	FwVersionResponseEventType       eventType = "fwVersionRsp"
	FwVersionUpdateRequestEventType  eventType = "fwUpdateReq"
	FwVersionUpdateResponseEventType eventType = "fwUpdateRsp"
	FwBlockResponseEventType         eventType = "fwBlockRsp"
	FwUpdateAbortType                eventType = "fwUpdateAbortReq"
)

const (
	UpgradeSuccessStatus        firmwareUpgradeStatus = "success"
	UpgradeDeviceNotFoundStatus firmwareUpgradeStatus = "deviceNotFound"
	UpgradeInvalidStateStatus   firmwareUpgradeStatus = "invalid_state"
	UpgradeInvalidFileStatus    firmwareUpgradeStatus = "invalid_file"
	UpgradeInvalidFileIdStatus  firmwareUpgradeStatus = "invalid_file_id"
	UpgradeUnknownErrorStatus   firmwareUpgradeStatus = "unknownError"
)

type firmwareUpgradeStatus string

type FirmwareVersionRequest struct {
	TransactionId uint32 `json:"-"`
}

func (f *FirmwareVersionRequest) UnmarshalJSON(bytes []byte) error {
	var e event

	if err := json.Unmarshal(bytes, &e); err != nil {
		return err
	}

	if e.EventType != FwVersionRequestEventType {
		return e.EventType.Error()
	}

	f.TransactionId = e.TransactionId

	return nil
}

func (f *FirmwareVersionRequest) MarshalJSON() ([]byte, error) {
	var e event

	e.EventType = FwVersionRequestEventType
	e.TransactionId = f.TransactionId

	return json.Marshal(&e)
}

type FirmwareVersionResponse struct {
	ShortAddr string `json:"-"`
	ExtAddr   string `json:"-"`
	Rssi      int    `json:"-"`
	FwVersion string `json:"fwVersion"`
}

func (f *FirmwareVersionResponse) UnmarshalJSON(bytes []byte) error {
	var e eventResponse

	if err := json.Unmarshal(bytes, &e); err != nil {
		return err
	}

	if e.EventType != FwVersionResponseEventType {
		return e.EventType.Error()
	}

	type firmwareVersionResponse FirmwareVersionResponse

	if err := json.Unmarshal(e.Payload, (*firmwareVersionResponse)(f)); err != nil {
		return err
	}

	f.ShortAddr = e.ShortAddr
	f.ExtAddr = e.ExtAddr
	f.Rssi = e.Rssi

	return nil
}

func (f *FirmwareVersionResponse) MarshalJSON() ([]byte, error) {
	var e eventResponse
	var err error

	type firmwareVersionResponse FirmwareVersionResponse

	e.EventType = FwVersionResponseEventType
	e.ShortAddr = f.ShortAddr
	e.ExtAddr = f.ExtAddr
	e.Rssi = f.Rssi

	if e.Payload, err = json.Marshal((*firmwareVersionResponse)(f)); err != nil {
		return nil, err
	}

	return json.Marshal(&e)
}

type FirmwareVersionUpgradeRequest struct {
	TransactionId uint32 `json:"-"`
	FileName      string `json:"fileName"`
}

func (f *FirmwareVersionUpgradeRequest) UnmarshalJSON(bytes []byte) error {
	var e event

	if err := json.Unmarshal(bytes, &e); err != nil {
		return err
	}

	if e.EventType != FwVersionUpdateRequestEventType {
		return e.EventType.Error()
	}

	type firmwareVersionUpgradeRequest FirmwareVersionUpgradeRequest

	if err := json.Unmarshal(e.Payload, (*firmwareVersionUpgradeRequest)(f)); err != nil {
		return err
	}

	f.TransactionId = e.TransactionId

	return nil
}

func (f *FirmwareVersionUpgradeRequest) MarshalJSON() ([]byte, error) {
	var e event
	var err error

	type firmwareVersionUpgradeRequest FirmwareVersionUpgradeRequest

	e.EventType = FwVersionUpdateRequestEventType
	e.TransactionId = f.TransactionId

	if e.Payload, err = json.Marshal((*firmwareVersionUpgradeRequest)(f)); err != nil {
		return nil, err
	}

	return json.Marshal(&e)
}

type FirmwareVersionUpgradeResponse struct {
	ShortAddr     string                `json:"-"`
	ExtAddr       string                `json:"-"`
	Rssi          int                   `json:"-"`
	TransactionId uint32                `json:"-"`
	ErrorCode     int                   `json:"errorCode"`
	Status        firmwareUpgradeStatus `json:"status"`
}

func (f *FirmwareVersionUpgradeResponse) UnmarshalJSON(bytes []byte) error {
	var r response

	if err := json.Unmarshal(bytes, &r); err != nil {
		return err
	}

	if r.EventType != FwVersionUpdateResponseEventType {
		return r.EventType.Error()
	}

	type firmwareVersionUpgradeResponse FirmwareVersionUpgradeResponse

	if err := json.Unmarshal(r.Payload, (*firmwareVersionUpgradeResponse)(f)); err != nil {
		return err
	}

	switch f.Status {
	case UpgradeSuccessStatus, UpgradeDeviceNotFoundStatus, UpgradeInvalidStateStatus,
		UpgradeInvalidFileStatus, UpgradeInvalidFileIdStatus, UpgradeUnknownErrorStatus:
	default:
		return InvalidFirmwareUpgradeStatus{f.Status}
	}

	f.Rssi = r.Rssi
	f.ExtAddr = r.ExtAddr
	f.ShortAddr = r.ShortAddr
	f.TransactionId = r.TransactionId

	return nil
}

func (f *FirmwareVersionUpgradeResponse) MarshalJSON() ([]byte, error) {
	switch f.Status {
	case UpgradeSuccessStatus, UpgradeDeviceNotFoundStatus, UpgradeInvalidStateStatus,
		UpgradeInvalidFileStatus, UpgradeInvalidFileIdStatus, UpgradeUnknownErrorStatus:
	default:
		return nil, &InvalidFirmwareUpgradeStatus{f.Status}
	}

	var r response
	var err error

	type firmwareVersionUpgradeResponse FirmwareVersionUpgradeResponse

	r.EventType = FwVersionUpdateResponseEventType
	r.Rssi = f.Rssi
	r.ExtAddr = f.ExtAddr
	r.ShortAddr = f.ShortAddr
	r.TransactionId = f.TransactionId

	if r.Payload, err = json.Marshal((*firmwareVersionUpgradeResponse)(f)); err != nil {
		return nil, err
	}

	return json.Marshal(&r)
}

type FirmwareBlockResponse struct {
	ShortAddr     string `json:"-"`
	ExtAddr       string `json:"-"`
	Rssi          int    `json:"-"`
	TransactionId uint32 `json:"-"`
	BlockNr       int    `json:"blockNr"`
	TotalBlocksNr int    `json:"totalBlocksNr"`
}

func (f *FirmwareBlockResponse) UnmarshalJSON(bytes []byte) error {
	var r response

	if err := json.Unmarshal(bytes, &r); err != nil {
		return err
	}

	if r.EventType != FwBlockResponseEventType {
		return r.EventType.Error()
	}

	type firmwareBlockResponse FirmwareBlockResponse

	if err := json.Unmarshal(r.Payload, (*firmwareBlockResponse)(f)); err != nil {
		return err
	}

	f.Rssi = r.Rssi
	f.ExtAddr = r.ExtAddr
	f.ShortAddr = r.ShortAddr
	f.TransactionId = r.TransactionId

	return nil
}

func (f *FirmwareBlockResponse) MarshalJSON() ([]byte, error) {
	var r response
	var err error

	type firmwareBlockResponse FirmwareBlockResponse

	r.EventType = FwBlockResponseEventType
	r.Rssi = f.Rssi
	r.ExtAddr = f.ExtAddr
	r.ShortAddr = f.ShortAddr
	r.TransactionId = f.TransactionId

	if r.Payload, err = json.Marshal((*firmwareBlockResponse)(f)); err != nil {
		return nil, err
	}

	return json.Marshal(&r)
}

type FirmwareUpdateAbort struct {
	TransactionId uint32 `json:"-"`
}

func (f *FirmwareUpdateAbort) UnmarshalJSON(bytes []byte) error {
	var e event
	if err := json.Unmarshal(bytes, &e); err != nil {
		return err
	}

	if e.EventType != FwUpdateAbortType {
		return e.EventType.Error()
	}

	f.TransactionId = e.TransactionId
	return nil
}

func (f *FirmwareUpdateAbort) MarshalJSON() ([]byte, error) {
	var e event

	e.TransactionId = f.TransactionId
	e.EventType = FwUpdateAbortType

	return json.Marshal(&e)
}
