package messages

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/rand"
	"strings"
	"testing"
)

func FuzzAuthRequestMarshal(f *testing.F) {
	for _, aT := range [...]authType{NoneType, NFCType, QRType, MobileType, NumPadType} {
		for _, aS := range [...]authStatus{NoneStatus, SuccessOfflineStatus, FailedOfflineStatus, FailedPrivacyStatus, VerifyOnlineStatus, FailedOnlineStatus, SuccessOnlineStatus, ErrorTimeNotSetStatus, NotFoundOfflineStatus, ErrorEncryptionStatus} {
			f.Add(rand.Int(), rand.Int(), rand.Int63(), string(aT), string(aS))
		}
	}

	f.Fuzz(func(t *testing.T, transactionId, hashKey int, timestamp int64, aT string, aS string) {
		value := AuthRequest{
			TransactionId: transactionId,
			HashKey:       fmt.Sprintf("%#x", hashKey),
			Timestamp:     timestamp,
			AuthStatus:    authStatus(aS),
			AuthType:      authType(aT),
		}

		res, err := json.Marshal(&value)

		if !strings.HasPrefix(value.HashKey, "0x") {
			require.ErrorAs(t, err, &InvalidHashKey{value.HashKey})
			return
		}

		if len(strings.TrimLeft(value.HashKey, "0x")) <= 0 {
			require.ErrorAs(t, err, &InvalidHashKey{value.HashKey})
			return
		}

		switch value.AuthType {
		case NoneType, NFCType, QRType, MobileType, NumPadType:
		default:
			require.ErrorAs(t, err, &InvalidAuthType{value.AuthType})
			return
		}

		switch value.AuthStatus {
		case NoneStatus, SuccessOfflineStatus, FailedOfflineStatus, FailedPrivacyStatus, VerifyOnlineStatus, FailedOnlineStatus, SuccessOnlineStatus, ErrorTimeNotSetStatus, NotFoundOfflineStatus, ErrorEncryptionStatus:
		default:
			require.ErrorAs(t, err, &InvalidAuthStatus{value.AuthStatus})
			return
		}

		require.NoError(t, err)
		j, _ := json.Marshal(map[string]any{
			"event": map[string]any{
				"eventType": "authEvent",
				"payload": map[string]any{
					"hashKey":    fmt.Sprintf("%#x", hashKey),
					"timestamp":  timestamp,
					"authType":   aT,
					"authStatus": aS,
					"channelIds": nil,
				},
				"transactionId": transactionId,
			},
		})
		assert.JSONEq(t, string(j), string(res))
	})
}

func FuzzAuthRequestUnmarshal(f *testing.F) {
	for _, aT := range [...]authType{NoneType, NFCType, QRType, MobileType, NumPadType} {
		for _, aS := range [...]authStatus{NoneStatus, SuccessOfflineStatus, FailedOfflineStatus, FailedPrivacyStatus, VerifyOnlineStatus, FailedOnlineStatus, SuccessOnlineStatus, ErrorTimeNotSetStatus, NotFoundOfflineStatus, ErrorEncryptionStatus} {
			f.Add(string(AuthEventType), rand.Int(), rand.Int(), rand.Int63(), string(aT), string(aS))
		}
	}

	f.Fuzz(func(t *testing.T, eT string, transactionId, hashKey int, timestamp int64, aT string, aS string) {
		j, _ := json.Marshal(map[string]any{
			"event": map[string]any{
				"eventType": eT,
				"payload": map[string]any{
					"hashKey":    fmt.Sprintf("%#x", hashKey),
					"timestamp":  timestamp,
					"authType":   aT,
					"authStatus": aS,
				},
				"transactionId": transactionId,
			},
		})
		var value AuthRequest

		err := json.Unmarshal(j, &value)

		if eventType(eT) != AuthEventType {
			target := eventType(eT).Error()
			require.ErrorAs(t, err, &target)
			return
		}

		switch authType(aT) {
		case NoneType, NFCType, QRType, MobileType, NumPadType:
		default:
			require.ErrorAs(t, err, &InvalidAuthType{value.AuthType})
			return
		}

		switch authStatus(aS) {
		case NoneStatus, SuccessOfflineStatus, FailedOfflineStatus, FailedPrivacyStatus, VerifyOnlineStatus, FailedOnlineStatus, SuccessOnlineStatus, ErrorTimeNotSetStatus, NotFoundOfflineStatus, ErrorEncryptionStatus:
		default:
			require.ErrorAs(t, err, &InvalidAuthStatus{value.AuthStatus})
			return
		}

		require.NoError(t, err)
		assert.Equal(t, AuthRequest{
			TransactionId: transactionId,
			HashKey:       fmt.Sprintf("%#x", hashKey),
			Timestamp:     timestamp,
			AuthType:      authType(aT),
			AuthStatus:    authStatus(aS),
			ChannelIds:    nil,
		}, value)
	})
}

func FuzzAuthResponseMarshal(f *testing.F) {
	for _, aT := range [...]authType{NoneType, NFCType, QRType, MobileType, NumPadType} {
		for _, aS := range [...]authStatus{NoneStatus, SuccessOfflineStatus, FailedOfflineStatus, FailedPrivacyStatus, VerifyOnlineStatus, FailedOnlineStatus, SuccessOnlineStatus, ErrorTimeNotSetStatus, NotFoundOfflineStatus, ErrorEncryptionStatus} {
			f.Add(rand.Int(), rand.Int(), rand.Int(), rand.Int(), rand.Int(), rand.Int63(), string(aT), string(aS))
		}
	}

	f.Fuzz(func(t *testing.T, shortAddr, extAddr, rssi, transactionId, hashKey int, timestamp int64, aT string, aS string) {
		value := AuthResponse{
			ExtAddr:       fmt.Sprintf("%#x", extAddr),
			ShortAddr:     fmt.Sprintf("%#x", shortAddr),
			Rssi:          rssi,
			TransactionId: transactionId,
			HashKey:       fmt.Sprintf("%#x", hashKey),
			Timestamp:     timestamp,
			AuthStatus:    authStatus(aS),
			AuthType:      authType(aT),
		}

		res, err := json.Marshal(&value)

		if !strings.HasPrefix(value.HashKey, "0x") {
			require.ErrorAs(t, err, &InvalidHashKey{value.HashKey})
			return
		}

		if len(strings.TrimLeft(value.HashKey, "0x")) <= 0 {
			require.ErrorAs(t, err, &InvalidHashKey{value.HashKey})
			return
		}

		switch value.AuthType {
		case NoneType, NFCType, QRType, MobileType, NumPadType:
		default:
			require.ErrorAs(t, err, &InvalidAuthType{value.AuthType})
			return
		}

		switch value.AuthStatus {
		case NoneStatus, SuccessOfflineStatus, FailedOfflineStatus, FailedPrivacyStatus, VerifyOnlineStatus, FailedOnlineStatus, SuccessOnlineStatus, ErrorTimeNotSetStatus, NotFoundOfflineStatus, ErrorEncryptionStatus:
		default:
			require.ErrorAs(t, err, &InvalidAuthStatus{value.AuthStatus})
			return
		}

		require.NoError(t, err)
		j, _ := json.Marshal(map[string]any{
			"ext_addr":   fmt.Sprintf("%#x", extAddr),
			"short_addr": fmt.Sprintf("%#x", shortAddr),
			"rssi":       rssi,
			"eventType":  "authEvent",
			"payload": map[string]any{
				"hashKey":    fmt.Sprintf("%#x", hashKey),
				"timestamp":  timestamp,
				"authType":   aT,
				"authStatus": aS,
				"channelIds": nil,
			},
			"transactionId": transactionId,
		})
		assert.JSONEq(t, string(j), string(res))
	})
}

func FuzzAuthResponseUnmarshal(f *testing.F) {
	for _, aT := range [...]authType{NoneType, NFCType, QRType, MobileType, NumPadType} {
		for _, aS := range [...]authStatus{NoneStatus, SuccessOfflineStatus, FailedOfflineStatus, FailedPrivacyStatus, VerifyOnlineStatus, FailedOnlineStatus, SuccessOnlineStatus, ErrorTimeNotSetStatus, NotFoundOfflineStatus, ErrorEncryptionStatus} {
			f.Add(string(AuthEventType), rand.Int(), rand.Int(), rand.Int(), rand.Int(), rand.Int(), rand.Int63(), string(aT), string(aS))
		}
	}

	f.Fuzz(func(t *testing.T, eT string, shortAddr, extAddr, rssi, transactionId, hashKey int, timestamp int64, aT string, aS string) {
		j, _ := json.Marshal(map[string]any{
			"ext_addr":   fmt.Sprintf("%#x", extAddr),
			"short_addr": fmt.Sprintf("%#x", shortAddr),
			"rssi":       rssi,
			"eventType":  eT,
			"payload": map[string]any{
				"hashKey":    fmt.Sprintf("%#x", hashKey),
				"timestamp":  timestamp,
				"authType":   aT,
				"authStatus": aS,
			},
			"transactionId": transactionId,
		})
		var value AuthResponse

		err := json.Unmarshal(j, &value)

		if eventType(eT) != AuthEventType {
			target := eventType(eT).Error()
			require.ErrorAs(t, err, &target)
			return
		}

		switch authType(aT) {
		case NoneType, NFCType, QRType, MobileType, NumPadType:
		default:
			require.ErrorAs(t, err, &InvalidAuthType{value.AuthType})
			return
		}

		switch authStatus(aS) {
		case NoneStatus, SuccessOfflineStatus, FailedOfflineStatus, FailedPrivacyStatus, VerifyOnlineStatus, FailedOnlineStatus, SuccessOnlineStatus, ErrorTimeNotSetStatus, NotFoundOfflineStatus, ErrorEncryptionStatus:
		default:
			require.ErrorAs(t, err, &InvalidAuthStatus{value.AuthStatus})
			return
		}

		require.NoError(t, err)
		assert.Equal(t, AuthResponse{
			ExtAddr:       fmt.Sprintf("%#x", extAddr),
			ShortAddr:     fmt.Sprintf("%#x", shortAddr),
			Rssi:          rssi,
			TransactionId: transactionId,
			HashKey:       fmt.Sprintf("%#x", hashKey),
			Timestamp:     timestamp,
			AuthType:      authType(aT),
			AuthStatus:    authStatus(aS),
			ChannelIds:    nil,
		}, value)
	})
}
