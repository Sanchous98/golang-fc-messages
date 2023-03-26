package messages

import (
	"bitbucket.org/4suites/golang-fc-messages/values"
	crypto "crypto/rand"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

func FuzzAuthMarshal(f *testing.F) {
	hash := make([]byte, 8)

	for _, aT := range [...]authType{NoneType, NFCType, QRType, MobileType, NumPadType} {
		for _, aS := range [...]authStatus{NoneStatus, SuccessOfflineStatus, FailedOfflineStatus, FailedPrivacyStatus, VerifyOnlineStatus, FailedOnlineStatus, SuccessOnlineStatus, ErrorTimeNotSetStatus, NotFoundOfflineStatus, ErrorEncryptionStatus} {
			_, _ = crypto.Read(hash)
			f.Add(rand.Int(), string(hash), rand.Int(), string(aT), string(aS))
		}
	}

	f.Fuzz(func(t *testing.T, transactionId int, hashKey string, timestamp int64, aT string, aS string) {
		value := Auth{
			TransactionId: transactionId,
			HashKey:       values.HashKey(hashKey),
			Timestamp:     (values.Timestamp)(time.Unix(timestamp, 0)),
			AuthStatus:    authStatus(aS),
			AuthType:      authType(aT),
		}

		res, err := json.Marshal(&value)

		if target := (*values.HashKey)(&hashKey).Validate(); target != nil {
			require.ErrorAs(t, err, &target)
			return
		}

		switch value.AuthType {
		case NoneType, NFCType, QRType, MobileType, NumPadType:
		default:
			target := invalidAuthType(value.AuthType)
			require.ErrorAs(t, err, &target)
			return
		}

		switch value.AuthStatus {
		case NoneStatus, SuccessOfflineStatus, FailedOfflineStatus, FailedPrivacyStatus, VerifyOnlineStatus, FailedOnlineStatus, SuccessOnlineStatus, ErrorTimeNotSetStatus, NotFoundOfflineStatus, ErrorEncryptionStatus:
		default:
			target := invalidAuthStatus(value.AuthStatus)
			require.ErrorAs(t, err, &target)
			return
		}

		require.NoError(t, err)
		assert.Equal(t, []byte(fmt.Sprintf(`{"event":{"eventType":"authEvent","payload":{"hashKey":"%s","timestamp":%d,"authType":"%s","authStatus":"%s","channelIds":null},"transactionId":%d}}}`, hashKey, timestamp, aT, aS, transactionId)), res)
	})
}

func FuzzAuthUnmarshal(f *testing.F) {
	hash := make([]byte, 8)

	for _, aT := range [...]authType{NoneType, NFCType, QRType, MobileType, NumPadType} {
		for _, aS := range [...]authStatus{NoneStatus, SuccessOfflineStatus, FailedOfflineStatus, FailedPrivacyStatus, VerifyOnlineStatus, FailedOnlineStatus, SuccessOnlineStatus, ErrorTimeNotSetStatus, NotFoundOfflineStatus, ErrorEncryptionStatus} {
			_, _ = crypto.Read(hash)
			f.Add(string(AuthEventType), rand.Int(), string(hash), rand.Int(), string(aT), string(aS))
		}
	}

	f.Fuzz(func(t *testing.T, eT string, transactionId int, hashKey string, timestamp int64, aT string, aS string) {
		j := []byte(fmt.Sprintf(`{"event":{"%s":"authEvent","payload":{"hashKey":"%s","timestamp":%d,"authType":"%s","authStatus":"%s","channelIds":null},"transactionId":%d}}}`, eT, hashKey, timestamp, aT, aS, transactionId))
		var value Auth

		err := json.Unmarshal(j, &value)

		if target := (*values.HashKey)(&hashKey).Validate(); target != nil {
			require.ErrorAs(t, err, &target)
			return
		}

		switch authType(aT) {
		case NoneType, NFCType, QRType, MobileType, NumPadType:
		default:
			target := invalidAuthType(authType(aT))
			require.ErrorAs(t, err, &target)
			return
		}

		switch authStatus(aS) {
		case NoneStatus, SuccessOfflineStatus, FailedOfflineStatus, FailedPrivacyStatus, VerifyOnlineStatus, FailedOnlineStatus, SuccessOnlineStatus, ErrorTimeNotSetStatus, NotFoundOfflineStatus, ErrorEncryptionStatus:
		default:
			target := invalidAuthStatus(authStatus(aS))
			require.ErrorAs(t, err, &target)
			return
		}

		require.NoError(t, err)
		assert.Equal(t, Auth{
			TransactionId: transactionId,
			HashKey:       values.HashKey(hashKey),
			Timestamp:     (values.Timestamp)(time.Unix(timestamp, 0)),
			AuthType:      authType(aT),
			AuthStatus:    authStatus(aS),
			ChannelIds:    nil,
		}, value)
	})
}
