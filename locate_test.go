package messages

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func FuzzLocateEventMarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, transactionId int) {
		value := &LocateRequest{TransactionId: transactionId}
		result, err := json.Marshal(value)

		require.NoError(t, err)
		assert.Equal(t, []byte(fmt.Sprintf(`{"event":{"eventType":"locateReq","payload":{},"transactionId":%d}}`, transactionId)), result)
	})
}

func FuzzLocateUnmarshal(f *testing.F) {
	f.Add(string(LocateRequestEventType), 0)

	f.Fuzz(func(t *testing.T, eT string, transactionId int) {
		var e LocateRequest
		value := []byte(fmt.Sprintf(`{"event":{"eventType":%q,"payload":{},"transactionId":%d}}`, eT, transactionId))

		err := json.Unmarshal(value, &e)

		if eventType(eT) != LocateRequestEventType {
			target := eventType(eT).Error()
			require.ErrorAs(t, err, &target)
			return
		}

		require.NoError(t, err)
		assert.Equal(t, transactionId, e.TransactionId)
	})
}