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

type Auth struct {
	TransactionId int        `json:"-"`
	HashKey       string     `json:"hashKey"`
	Timestamp     int64      `json:"timestamp"`
	AuthType      authType   `json:"authType"`
	AuthStatus    authStatus `json:"authStatus"`
	ChannelIds    []int      `json:"channelIds"`
}

func (a *Auth) UnmarshalJSON(bytes []byte) error {
	type auth Auth

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
func (a *Auth) MarshalJSON() ([]byte, error) {
	type auth Auth

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