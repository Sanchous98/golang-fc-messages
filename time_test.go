package messages

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func FuzzTimeSyncEventMarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, transactionId int) {
		value := &TimeSyncEvent{TransactionId: transactionId}
		result, err := json.Marshal(value)

		require.NoError(t, err)
		assert.Equal(t, []byte(fmt.Sprintf(`{"event":{"eventType":"timeSync","payload":{},"transactionId":%d}}`, transactionId)), result)
	})
}

func FuzzTimeSyncEventUnmarshal(f *testing.F) {
	f.Add(string(TimeSyncEventType), 0)

	f.Fuzz(func(t *testing.T, eT string, transactionId int) {
		var e TimeSyncEvent
		value := []byte(fmt.Sprintf(`{"event":{"eventType":"%s","payload":{},"transactionId":%d}}`, eT, transactionId))

		err := json.Unmarshal(value, &e)

		if eventType(eT) != TimeSyncEventType {
			target := eventType(eT).Error()
			require.ErrorAs(t, err, &target)
			return
		}

		require.NoError(t, err)
		assert.Equal(t, transactionId, e.TransactionId)
	})
}
