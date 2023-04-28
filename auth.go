package messages

import (
	"github.com/goccy/go-json"
	"strings"
)

type authStatus string

type authType string

const AuthEventType eventType = "authEvent"

const (
	NoneStatus            authStatus = "none"
	SuccessOfflineStatus  authStatus = "succesOffline"
	FailedOfflineStatus   authStatus = "failedOffline"
	FailedPrivacyStatus   authStatus = "failedPrivacy"
	VerifyOnlineStatus    authStatus = "verifyOnline"
	FailedOnlineStatus    authStatus = "failedOnline"
	SuccessOnlineStatus   authStatus = "successOnline"
	ErrorTimeNotSetStatus authStatus = "errorTimeNotSet"
	NotFoundOfflineStatus authStatus = "NotFoundOffline"
	ErrorEncryptionStatus authStatus = "errorEncryption"
)

const (
	NoneType   authType = "none"
	NFCType    authType = "NFC"
	QRType     authType = "QR"
	MobileType authType = "Mobile"
	NumPadType authType = "numPad"
)

type AuthRequest struct {
	TransactionId int        `json:"-"`
	HashKey       string     `json:"hashKey"`
	Timestamp     int64      `json:"timestamp"`
	AuthType      authType   `json:"authType"`
	AuthStatus    authStatus `json:"authStatus"`
	ChannelIds    []int      `json:"channelIds"`
}

func (a *AuthRequest) MarshalJSON() ([]byte, error) {
	type auth AuthRequest

	switch a.AuthType {
	case NoneType, NFCType, QRType, MobileType, NumPadType:
	default:
		return nil, InvalidAuthType{a.AuthType}
	}

	switch a.AuthStatus {
	case NoneStatus, SuccessOfflineStatus, FailedOfflineStatus, FailedPrivacyStatus, VerifyOnlineStatus, FailedOnlineStatus, SuccessOnlineStatus, ErrorTimeNotSetStatus, NotFoundOfflineStatus, ErrorEncryptionStatus:
	default:
		return nil, InvalidAuthStatus{a.AuthStatus}
	}

	var e event
	var err error

	if !strings.HasPrefix(a.HashKey, "0x") {
		return nil, InvalidHashKey{a.HashKey}
	}

	if strings.TrimLeft(a.HashKey, "0x") == "" {
		return nil, InvalidHashKey{a.HashKey}
	}

	e.TransactionId = a.TransactionId
	e.EventType = AuthEventType

	if e.Payload, err = json.Marshal((*auth)(a)); err != nil {
		return nil, err
	}

	return json.Marshal(&e)
}

func (a *AuthRequest) UnmarshalJSON(bytes []byte) error {
	type auth AuthRequest

	var e event

	if err := json.Unmarshal(bytes, &e); err != nil {
		return err
	}

	if e.EventType != AuthEventType {
		return e.EventType.Error()
	}

	if err := json.Unmarshal(e.Payload, (*auth)(a)); err != nil {
		return err
	}

	if !strings.HasPrefix(a.HashKey, "0x") {
		return InvalidHashKey{a.HashKey}
	}

	if strings.TrimLeft(a.HashKey, "0x") == "" {
		return InvalidHashKey{a.HashKey}
	}

	switch a.AuthType {
	case NoneType, NFCType, QRType, MobileType, NumPadType:
	default:
		return InvalidAuthType{a.AuthType}
	}

	switch a.AuthStatus {
	case NoneStatus, SuccessOfflineStatus, FailedOfflineStatus, FailedPrivacyStatus, VerifyOnlineStatus, FailedOnlineStatus, SuccessOnlineStatus, ErrorTimeNotSetStatus, NotFoundOfflineStatus, ErrorEncryptionStatus:
	default:
		return InvalidAuthStatus{a.AuthStatus}
	}

	a.TransactionId = e.TransactionId

	return nil
}

type AuthResponse struct {
	ShortAddr     string     `json:"-"`
	ExtAddr       string     `json:"-"`
	Rssi          int        `json:"-"`
	TransactionId int        `json:"-"`
	HashKey       string     `json:"hashKey"`
	Timestamp     int64      `json:"timestamp"`
	AuthType      authType   `json:"authType"`
	AuthStatus    authStatus `json:"authStatus"`
	ChannelIds    []int      `json:"channelIds"`
}

func (a *AuthResponse) UnmarshalJSON(bytes []byte) error {
	type auth AuthResponse

	var r response

	if err := json.Unmarshal(bytes, &r); err != nil {
		return err
	}

	if r.EventType != AuthEventType {
		return r.EventType.Error()
	}

	if err := json.Unmarshal(r.Payload, (*auth)(a)); err != nil {
		return err
	}

	if !strings.HasPrefix(a.HashKey, "0x") {
		return InvalidHashKey{a.HashKey}
	}

	if strings.TrimLeft(a.HashKey, "0x") == "" {
		return InvalidHashKey{a.HashKey}
	}

	switch a.AuthType {
	case NoneType, NFCType, QRType, MobileType, NumPadType:
	default:
		return InvalidAuthType{a.AuthType}
	}

	switch a.AuthStatus {
	case NoneStatus, SuccessOfflineStatus, FailedOfflineStatus, FailedPrivacyStatus, VerifyOnlineStatus, FailedOnlineStatus, SuccessOnlineStatus, ErrorTimeNotSetStatus, NotFoundOfflineStatus, ErrorEncryptionStatus:
	default:
		return InvalidAuthStatus{a.AuthStatus}
	}

	a.TransactionId = r.TransactionId
	a.ShortAddr = r.ShortAddr
	a.ExtAddr = r.ExtAddr
	a.Rssi = r.Rssi

	return nil
}
func (a *AuthResponse) MarshalJSON() ([]byte, error) {
	type auth AuthResponse

	switch a.AuthType {
	case NoneType, NFCType, QRType, MobileType, NumPadType:
	default:
		return nil, InvalidAuthType{a.AuthType}
	}

	switch a.AuthStatus {
	case NoneStatus, SuccessOfflineStatus, FailedOfflineStatus, FailedPrivacyStatus, VerifyOnlineStatus, FailedOnlineStatus, SuccessOnlineStatus, ErrorTimeNotSetStatus, NotFoundOfflineStatus, ErrorEncryptionStatus:
	default:
		return nil, InvalidAuthStatus{a.AuthStatus}
	}

	var r response
	var err error

	if !strings.HasPrefix(a.HashKey, "0x") {
		return nil, InvalidHashKey{a.HashKey}
	}

	if strings.TrimLeft(a.HashKey, "0x") == "" {
		return nil, InvalidHashKey{a.HashKey}
	}

	r.TransactionId = a.TransactionId
	r.ShortAddr = a.ShortAddr
	r.ExtAddr = a.ExtAddr
	r.Rssi = a.Rssi
	r.EventType = AuthEventType

	if r.Payload, err = json.Marshal((*auth)(a)); err != nil {
		return nil, err
	}

	return json.Marshal(&r)
}
