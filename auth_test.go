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
			f.Add(rand.Int(), string(hash), rand.Int(), string(aT), string(aS))
		}
	}

	f.Fuzz(func(t *testing.T, transactionId int, hashKey string, timestamp int, aT string, aS string) {
		value := Auth{
			TransactionId: transactionId,
			HashKey:       hashKey,
			Timestamp:     timestamp,
			AuthStatus:    authStatus(aS),
			AuthType:      authType(aT),
		}

		res, err := json.MarshalWithOption(&value)

		if !strings.HasPrefix(value.HashKey, "0x") {
			target := invalidHashKey(value.HashKey)
			require.ErrorAs(t, err, &target)
			return
		}

		if len(strings.TrimLeft(value.HashKey, "0x")) <= 0 || len(strings.TrimLeft(value.HashKey, "0x"))%2 != 0 {
			target := invalidHashKey(value.HashKey)
			require.ErrorAs(t, err, &target)
			return
		}

		for _, letter := range strings.TrimLeft(value.HashKey, "0x") {
			if !(letter >= '0' && letter <= '9' || letter >= 'a' && letter <= 'f' || letter >= 'A' && letter <= 'F') {
				target := invalidHashKey(value.HashKey)
				require.ErrorAs(t, err, &target)
				return
			}
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

//func FuzzAuthUnmarshal(f *testing.F) {
//	f.Fuzz(func(t *testing.T) {})
//}
