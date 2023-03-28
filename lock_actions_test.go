package messages

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/require"
	"testing"
)

func FuzzLockResponseMarshal(f *testing.F) {
	f.Add(string(NoneLockStatus))
	f.Add(string(ExtRelayStateLockStatus))
	f.Add(string(LockOpenedLockStatus))
	f.Add(string(LockClosedLockStatus))
	f.Add(string(DriverOnLockStatus))
	f.Add(string(ErrorLockAlreadyOpenLockStatus))
	f.Add(string(ErrorLockAlreadyClosedLockStatus))
	f.Add(string(ErrorDriverEnabledLockStatus))
	f.Add(string(DeviceTypeUnknownLockStatus))
	f.Add(string(OpenTimeoutLockStatus))

	f.Fuzz(func(t *testing.T, lockActionStatus string) {
		test := []byte(fmt.Sprintf(`{"short_addr":"0x87","ext_addr":"0x124b001e19b66b","rssi":-60,"eventType":"lockActionResponse","payload":{"lockActionStatus":"%s"},"transactionId":0}`, lockActionStatus))

		var rsp LockResponse
		err := json.Unmarshal(test, &rsp)

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
	})
}
