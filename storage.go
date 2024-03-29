package messages

import (
	"github.com/goccy/go-json"
	"time"
)

const (
	LocalStorageAddKeyEventType    eventType = "localStorageAddKey"
	LocalStorageUpdateKeyEventType eventType = "localStorageUpdateKey"
	LocalStorageGetKeyEventType    eventType = "localStorageGetKey"
	LocalStorageDeleteKeyEventType eventType = "localStorageDeleteKey"
	LocalStorageResponseEventType  eventType = "localStorageResponse"
)

const (
	StorageResponseStatusOk storageResponseStatus = iota
	StorageResponseStatusReadOk
	StorageResponseStatusErrorKeyNotFound
	StorageResponseStatusErrorKeyAlreadyExists
	StorageResponseStatusErrorFlashStorageFull
	StorageResponseStatusErrorCritical
	_maxStorageResponseStatus = iota - 1
)

type storageResponseStatus uint8

func (s *storageResponseStatus) UnmarshalJSON(bytes []byte) (err error) {
	defer func() {
		if s != nil && *s < 0 || *s > _maxStorageResponseStatus {
			err = InvalidStorageResponseStatus{*s}
		}
	}()

	err = json.Unmarshal(bytes, (*uint8)(s))
	return
}

func (s *storageResponseStatus) MarshalJSON() ([]byte, error) {
	if s != nil && *s < 0 || *s > _maxStorageResponseStatus {
		return nil, InvalidStorageResponseStatus{*s}
	}

	return json.Marshal((*uint8)(s))
}

type MasterKey struct {
	ChannelIds []int `json:"channelIds,omitempty"`
}

type TimeKey struct {
	StartTime  int   `json:"startTime"`
	EndTime    int   `json:"endTime"`
	ChannelIds []int `json:"channelIds,omitempty"`
}

type AclKey struct {
	DaysOfWeek []time.Weekday `json:"daysOfWeek"`
	StartTime  string         `json:"startTime"`
	EndTime    string         `json:"endTime"`
	ChannelIds []int          `json:"channelIds,omitempty"`
}

type Flags struct {
	MasterKey            bool `json:"masterKey"`
	PrivacyOverride      bool `json:"privacyOverride"`
	IsMultiChannel       bool `json:"isMultiChannel"`
	IsMeetingModeAllowed bool `json:"isMeetingModeAllowed"`
}

type StorageData struct {
	Status    storageResponseStatus `json:"status"`
	HashKey   string                `json:"hashKey"`
	Flags     Flags                 `json:"flags"`
	MasterKey MasterKey             `json:"masterKey,omitempty"`
	TimeKeys  []TimeKey             `json:"timeKeys,omitempty"`
	AclKeys   []AclKey              `json:"aclKeys,omitempty"`
}

func (s *StorageData) UnmarshalJSON(bytes []byte) error {
	type storageData StorageData

	if err := json.Unmarshal(bytes, (*storageData)(s)); err != nil {
		return err
	}

	return nil
}

func (s *StorageData) MarshalJSON() ([]byte, error) {
	type storageData StorageData

	return json.Marshal((*storageData)(s))
}

type StorageAddKey struct {
	TransactionId uint32 `json:"-"`
	StorageData
}

func (s *StorageAddKey) UnmarshalJSON(bytes []byte) error {
	type storageAddKey StorageAddKey

	var e event

	if err := json.Unmarshal(bytes, &e); err != nil {
		return err
	}

	if e.EventType != LocalStorageAddKeyEventType {
		return e.EventType.Error()
	}

	if err := json.Unmarshal(e.Payload, (*storageAddKey)(s)); err != nil {
		return err
	}

	s.TransactionId = e.TransactionId

	return nil
}

func (s *StorageAddKey) MarshalJSON() ([]byte, error) {
	type storageAddKey StorageAddKey

	var e event
	var err error

	e.EventType = LocalStorageAddKeyEventType
	e.TransactionId = s.TransactionId

	if e.Payload, err = json.Marshal((*storageAddKey)(s)); err != nil {
		return nil, err
	}

	return json.Marshal(&e)
}

type StorageUpdateKey struct {
	TransactionId uint32 `json:"-"`
	StorageData
}

func (s *StorageUpdateKey) UnmarshalJSON(bytes []byte) error {
	type storageUpdateKey StorageUpdateKey

	var e event

	if err := json.Unmarshal(bytes, &e); err != nil {
		return err
	}

	if e.EventType != LocalStorageUpdateKeyEventType {
		return e.EventType.Error()
	}

	if err := json.Unmarshal(e.Payload, (*storageUpdateKey)(s)); err != nil {
		return err
	}

	s.TransactionId = e.TransactionId

	return nil
}

func (s *StorageUpdateKey) MarshalJSON() ([]byte, error) {
	type storageUpdateKey StorageUpdateKey

	var e event
	var err error

	e.EventType = LocalStorageUpdateKeyEventType
	e.TransactionId = s.TransactionId

	if e.Payload, err = json.Marshal((*storageUpdateKey)(s)); err != nil {
		return nil, err
	}

	return json.Marshal(&e)
}

type StorageGetKey struct {
	TransactionId uint32 `json:"-"`
	HashKey       string `json:"hashKey"`
}

func (s *StorageGetKey) UnmarshalJSON(bytes []byte) error {
	var e event

	if err := json.Unmarshal(bytes, &e); err != nil {
		return err
	}

	if e.EventType != LocalStorageGetKeyEventType {
		return e.EventType.Error()
	}

	type storageGetKey StorageGetKey

	if err := json.Unmarshal(e.Payload, (*storageGetKey)(s)); err != nil {
		return err
	}

	s.TransactionId = e.TransactionId

	return nil
}

func (s *StorageGetKey) MarshalJSON() ([]byte, error) {
	var e event
	var err error

	type storageGetKey StorageGetKey

	e.EventType = LocalStorageGetKeyEventType
	e.TransactionId = s.TransactionId

	if e.Payload, err = json.Marshal((*storageGetKey)(s)); err != nil {
		return nil, err
	}

	return json.Marshal(&e)
}

type StorageDeleteKey struct {
	TransactionId uint32 `json:"-"`
	HashKey       string `json:"hashKey"`
}

func (s *StorageDeleteKey) UnmarshalJSON(bytes []byte) error {
	var e event

	if err := json.Unmarshal(bytes, &e); err != nil {
		return err
	}

	if e.EventType != LocalStorageDeleteKeyEventType {
		return e.EventType.Error()
	}

	type storageDeleteKey StorageDeleteKey

	if err := json.Unmarshal(e.Payload, (*storageDeleteKey)(s)); err != nil {
		return err
	}

	s.TransactionId = e.TransactionId

	return nil
}

func (s *StorageDeleteKey) MarshalJSON() ([]byte, error) {
	var e event
	var err error

	type storageDeleteKey StorageDeleteKey

	e.EventType = LocalStorageDeleteKeyEventType
	e.TransactionId = s.TransactionId

	if e.Payload, err = json.Marshal((*storageDeleteKey)(s)); err != nil {
		return nil, err
	}

	return json.Marshal(&e)
}

type StorageResponse struct {
	ShortAddr     string `json:"-"`
	ExtAddr       string `json:"-"`
	Rssi          int    `json:"-"`
	TransactionId uint32 `json:"-"`
	StorageData
}

func (s *StorageResponse) UnmarshalJSON(bytes []byte) error {
	var r response

	if err := json.Unmarshal(bytes, &r); err != nil {
		return err
	}

	if r.EventType != LocalStorageResponseEventType {
		return r.EventType.Error()
	}

	type storageResponse StorageResponse

	if err := json.Unmarshal(r.Payload, (*storageResponse)(s)); err != nil {
		return err
	}

	s.Rssi = r.Rssi
	s.ExtAddr = r.ExtAddr
	s.ShortAddr = r.ShortAddr
	s.TransactionId = r.TransactionId

	return nil
}

func (s *StorageResponse) MarshalJSON() ([]byte, error) {
	var r response
	var err error

	type storageResponse StorageResponse

	r.EventType = LocalStorageResponseEventType
	r.Rssi = s.Rssi
	r.ExtAddr = s.ExtAddr
	r.ShortAddr = s.ShortAddr
	r.TransactionId = s.TransactionId

	if r.Payload, err = json.Marshal((*storageResponse)(s)); err != nil {
		return nil, err
	}

	return json.Marshal(&r)
}
