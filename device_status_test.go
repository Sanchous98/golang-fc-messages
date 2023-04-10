package messages

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

func FuzzDeviceStatusRequestMarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, transactionId int) {
		value := &DeviceStatusRequest{TransactionId: transactionId}
		result, err := json.Marshal(value)

		require.NoError(t, err)
		assert.Equal(t, []byte(fmt.Sprintf(`{"event":{"eventType":"deviceStatusReq","payload":{},"transactionId":%d}}`, transactionId)), result)
	})
}

func FuzzDeviceStatusRequestUnmarshal(f *testing.F) {
	f.Add(string(DeviceStatusRequestEvent), 0)

	f.Fuzz(func(t *testing.T, eT string, transactionId int) {
		var e DeviceStatusRequest
		value := []byte(fmt.Sprintf(`{"event":{"eventType":%q,"payload":{},"transactionId":%d}}`, eT, transactionId))

		err := json.Unmarshal(value, &e)

		if eventType(eT) != DeviceStatusRequestEvent {
			target := eventType(eT).Error()
			require.ErrorAs(t, err, &target)
			return
		}

		require.NoError(t, err)
		assert.Equal(t, transactionId, e.TransactionId)
	})
}

func FuzzDeviceStatusResponseMarshal(f *testing.F) {
	f.Add(rand.Int(), rand.Int(), rand.Int(), rand.Int(), string(CloudRequestedReason), time.Now().Unix(), 0, 0, 0, 0)
	f.Add(rand.Int(), rand.Int(), rand.Int(), rand.Int(), string(ScheduledUpdateReason), time.Now().Unix(), 0, 0, 0, 0)

	f.Fuzz(func(t *testing.T, shortAddr, extAddr, rssi, transactionId int, reason string, time int64, batteryLevel, batteryLevelLoad, networkState, autoRequest int) {
		value := &DeviceStatusResponse{
			ShortAddr:        fmt.Sprintf("%#x", shortAddr),
			ExtAddr:          fmt.Sprintf("%#x", extAddr),
			Rssi:             rssi,
			TransactionId:    transactionId,
			Reason:           deviceStatusReason(reason),
			Time:             time,
			BatteryLevel:     batteryLevel,
			BatteryLevelLoad: batteryLevelLoad,
			NetworkState:     networkState,
			AutoRequest:      autoRequest,
		}
		result, err := json.Marshal(value)

		switch value.Reason {
		case CloudRequestedReason, ScheduledUpdateReason:
		default:
			target := invalidDeviceStatusReason(value.Reason)
			require.ErrorAs(t, err, &target)
			return
		}

		require.NoError(t, err)
		assert.Equal(t, []byte(fmt.Sprintf(`{"short_addr":"%#x","ext_addr":"%#x","rssi":%d,"eventType":"deviceStatusRsp","payload":{"reason":%q,"time":%d,"batteryLevel":%d,"batteryLevelLoad":%d,"networkState":%d,"autoRequest":%d},"transactionId":%d}`, shortAddr, extAddr, rssi, reason, time, batteryLevel, batteryLevelLoad, networkState, autoRequest, transactionId)), result)
	})
}

func FuzzDeviceStatusResponseUnmarshal(f *testing.F) {
	f.Add(rand.Int(), rand.Int(), rand.Int(), string(DeviceStatusResponseEvent), 0, string(CloudRequestedReason), time.Now().Unix(), 0, 0, 0, 0)
	f.Add(rand.Int(), rand.Int(), rand.Int(), string(DeviceStatusResponseEvent), 0, string(ScheduledUpdateReason), time.Now().Unix(), 0, 0, 0, 0)

	f.Fuzz(func(t *testing.T, shortAddr, extAddr, rssi int, eT string, transactionId int, reason string, time int64, batteryLevel, batteryLevelLoad, networkState, autoRequest int) {
		expected := []byte(fmt.Sprintf(`{"short_addr":"%#x","ext_addr":"%#x","rssi":%d,"eventType":%q,"payload":{"reason":%q,"time":%d,"batteryLevel":%d,"batteryLevelLoad":%d,"networkState":%d,"autoRequest":%d},"transactionId":%d}`, shortAddr, extAddr, rssi, eT, reason, time, batteryLevel, batteryLevelLoad, networkState, autoRequest, transactionId))

		var value DeviceStatusResponse

		err := json.Unmarshal(expected, &value)

		if eventType(eT) != DeviceStatusResponseEvent {
			target := eventType(eT).Error()
			require.ErrorAs(t, err, &target)
		}

		switch value.Reason {
		case CloudRequestedReason, ScheduledUpdateReason:
		default:
			target := invalidDeviceStatusReason(value.Reason)
			require.ErrorAs(t, err, &target)
			return
		}

		require.NoError(t, err)
		assert.Equal(t, DeviceStatusResponse{
			ShortAddr:        fmt.Sprintf("%#x", shortAddr),
			ExtAddr:          fmt.Sprintf("%#x", extAddr),
			Rssi:             rssi,
			TransactionId:    transactionId,
			Reason:           deviceStatusReason(reason),
			Time:             time,
			BatteryLevel:     batteryLevel,
			BatteryLevelLoad: batteryLevelLoad,
			NetworkState:     networkState,
			AutoRequest:      autoRequest,
		}, value)
	})
}
