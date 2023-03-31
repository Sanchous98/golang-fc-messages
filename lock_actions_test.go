package messages

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

func FuzzLockResponseUnmarshal(f *testing.F) {
	for _, lT := range [...]lockStatus{NoneLockStatus, ExtRelayStateLockStatus, LockOpenedLockStatus, LockClosedLockStatus, DriverOnLockStatus, ErrorLockAlreadyOpenLockStatus, ErrorLockAlreadyClosedLockStatus, ErrorDriverEnabledLockStatus, DeviceTypeUnknownLockStatus, OpenTimeoutLockStatus} {
		f.Add(string(lT), rand.Int(), rand.Int(), rand.Int(), string(LockActionResponseEventType), rand.Int())
	}

	f.Fuzz(func(t *testing.T, lockActionStatus string, shortAddr, extAddr, rssi int, eT string, transactionId int) {
		test := []byte(fmt.Sprintf(`{"short_addr":"%#x","ext_addr":"%#x","rssi":%d,"eventType":%q,"payload":{"lockActionStatus":%q},"transactionId":%d}`, shortAddr, extAddr, rssi, eT, lockActionStatus, transactionId))

		var rsp LockResponse
		err := json.Unmarshal(test, &rsp)

		if eventType(eT) != LockActionResponseEventType {
			target := eventType(eT).Error()
			require.ErrorAs(t, err, &target)
			return
		}

		switch lockStatus(lockActionStatus) {
		case NoneLockStatus, ExtRelayStateLockStatus, LockOpenedLockStatus, LockClosedLockStatus, DriverOnLockStatus,
			ErrorLockAlreadyOpenLockStatus, ErrorLockAlreadyClosedLockStatus, ErrorDriverEnabledLockStatus,
			DeviceTypeUnknownLockStatus:
		case OpenTimeoutLockStatus:
			fallthrough
		default:
			target := invalidLockStatus(lockStatus(lockActionStatus))
			require.ErrorAs(t, err, &target)
			return
		}

		require.NoError(t, err)
		assert.Equal(t, LockResponse{
			ShortAddr:        fmt.Sprintf("%#x", shortAddr),
			ExtAddr:          fmt.Sprintf("%#x", extAddr),
			Rssi:             rssi,
			TransactionId:    transactionId,
			LockActionStatus: lockStatus(lockActionStatus),
		}, rsp)
	})
}
