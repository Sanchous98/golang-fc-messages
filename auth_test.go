package messages

import (
	crypto "crypto/rand"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/rand"
	"strings"
	"testing"
)

func FuzzAuthMarshal(f *testing.F) {
	hash := make([]byte, 8)

	for _, aT := range [...]authType{NoneType, NFCType, QRType, MobileType, NumPadType} {
		for _, aS := range [...]authStatus{NoneStatus, SuccessOfflineStatus, FailedOfflineStatus, FailedPrivacyStatus, VerifyOnlineStatus, FailedOnlineStatus, SuccessOnlineStatus, ErrorTimeNotSetStatus, NotFoundOfflineStatus, ErrorEncryptionStatus} {
			_, _ = crypto.Read(hash)
			f.Add(rand.Int(), string(hash), rand.Int63(), string(aT), string(aS))
		}
	}

	f.Fuzz(func(t *testing.T, transactionId int, hashKey string, timestamp int64, aT string, aS string) {
		value := Auth{
			TransactionId: transactionId,
			HashKey:       hashKey,
			Timestamp:     timestamp,
			AuthStatus:    authStatus(aS),
			AuthType:      authType(aT),
		}

		res, err := json.Marshal(&value)

		if !strings.HasPrefix(value.HashKey, "0x") {
			require.ErrorAs(t, err, &InvalidHashKey{value.HashKey})
			return
		}

		if len(strings.TrimLeft(value.HashKey, "0x")) <= 0 || len(strings.TrimLeft(value.HashKey, "0x"))%2 != 0 {
			require.ErrorAs(t, err, &InvalidHashKey{value.HashKey})
			return
		}

		for _, letter := range strings.TrimLeft(value.HashKey, "0x") {
			if !(letter >= '0' && letter <= '9' || letter >= 'a' && letter <= 'f' || letter >= 'A' && letter <= 'F') {
				require.ErrorAs(t, err, &InvalidHashKey{value.HashKey})
				return
			}
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
			require.ErrorAs(t, err, &InvalidAuthType{value.AuthType})
			return
		}

		require.NoError(t, err)
		assert.Equal(t, []byte(fmt.Sprintf(`{"event":{"eventType":"authEvent","payload":{"hashKey":%q,"timestamp":%d,"authType":%q,"authStatus":%q,"channelIds":null},"transactionId":%d}}`, hashKey, timestamp, aT, aS, transactionId)), res)
	})
}

func FuzzAuthUnmarshal(f *testing.F) {
	for _, aT := range [...]authType{NoneType, NFCType, QRType, MobileType, NumPadType} {
		for _, aS := range [...]authStatus{NoneStatus, SuccessOfflineStatus, FailedOfflineStatus, FailedPrivacyStatus, VerifyOnlineStatus, FailedOnlineStatus, SuccessOnlineStatus, ErrorTimeNotSetStatus, NotFoundOfflineStatus, ErrorEncryptionStatus} {
			f.Add(string(AuthEventType), rand.Int(), rand.Int(), rand.Int63(), string(aT), string(aS))
		}
	}

	f.Fuzz(func(t *testing.T, eT string, transactionId, hashKey int, timestamp int64, aT string, aS string) {
		j := []byte(fmt.Sprintf(`{"event":{"eventType":%q,"payload":{"hashKey":"%#x","timestamp":%d,"authType":%q,"authStatus":%q,"channelIds":null},"transactionId":%d}}`, eT, hashKey, timestamp, aT, aS, transactionId))
		var value Auth

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
			require.ErrorAs(t, err, &InvalidAuthType{value.AuthType})
			return
		}

		require.NoError(t, err)
		assert.Equal(t, Auth{
			TransactionId: transactionId,
			HashKey:       fmt.Sprintf("%#x", hashKey),
			Timestamp:     timestamp,
			AuthType:      authType(aT),
			AuthStatus:    authStatus(aS),
			ChannelIds:    nil,
		}, value)
	})
}
