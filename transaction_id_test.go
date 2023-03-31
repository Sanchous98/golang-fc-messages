package messages

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func FuzzTransactionIdActionMarshal(f *testing.F) {
	f.Add(string(TransactionActionRead))
	f.Add(string(TransactionActionReset))

	f.Fuzz(func(t *testing.T, a string) {
		value := &TransactionIdAction{Action: transactionIdAction(a)}
		result, err := json.Marshal(value)

		switch value.Action {
		case TransactionActionRead, TransactionActionReset:
			require.NoError(t, err)
			assert.Equal(t, []byte(fmt.Sprintf(`{"event":{"eventType":"transactionIdReq","payload":{"transactionIdAction":%q},"transactionId":0}}`, string(value.Action))), result)
		default:
			target := invalidTransactionIdAction(value.Action)
			require.ErrorAs(t, err, &target)
		}
	})
}

func FuzzTransactionIdActionUnmarshal(f *testing.F) {
	f.Add(string(TransactionActionRead), string(TransactionIdReq))
	f.Add(string(TransactionActionReset), string(TransactionIdReq))

	f.Fuzz(func(t *testing.T, a string, eT string) {
		value := []byte(fmt.Sprintf(`{"event":{"eventType":%q,"payload":{"transactionIdAction":%q},"transactionId":0}}`, eT, a))
		var j TransactionIdAction
		err := json.Unmarshal(value, &j)

		if eventType(eT) != TransactionIdReq {
			target := eventType(eT).Error()
			require.ErrorAs(t, err, &target)
			return
		}

		switch transactionIdAction(a) {
		case TransactionActionRead, TransactionActionReset:
			require.NoError(t, err)
			assert.Equal(t, j.Action, transactionIdAction(a))
		default:
			target := invalidTransactionIdAction(j.Action)
			require.ErrorAs(t, err, &target)
		}
	})
}

func FuzzTransactionIdResponseMarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, deviceTransactionId int) {
		value := &TransactionIdResponse{DeviceTransactionId: deviceTransactionId}
		result, err := json.Marshal(value)

		require.NoError(t, err)
		assert.Equal(t, []byte(fmt.Sprintf(`{"short_addr":"","ext_addr":"","rssi":0,"eventType":"transactionIdRsp","payload":{"deviceTransactionId":%d},"transactionId":0}`, deviceTransactionId)), result)
	})
}

func FuzzTransactionIdResponseUnmarshal(f *testing.F) {
	f.Add(string(TransactionIdRsp), 0)

	f.Fuzz(func(t *testing.T, eT string, deviceTransactionId int) {
		value := []byte(fmt.Sprintf(`{"short_addr":"0x2","ext_addr":"0x124b001cbd4efa","rssi":-45,"eventType":%q,"payload":{"deviceTransactionId":%d},"transactionId":0}`, eT, deviceTransactionId))
		var j TransactionIdResponse
		err := json.Unmarshal(value, &j)

		if eT != string(TransactionIdRsp) {
			target := TransactionIdRsp.Error()
			require.ErrorAs(t, err, &target)
			return
		}

		require.NoError(t, err)
		assert.Equal(t, deviceTransactionId, j.DeviceTransactionId)
	})
}
