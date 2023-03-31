package messages

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func FuzzGetNetworkInfoMarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, transactionId int) {
		value := GetNetworkInfo{TransactionId: transactionId}
		res, err := json.Marshal(&value)

		require.NoError(t, err)
		assert.Equal(t, []byte(fmt.Sprintf(`{"event":{"eventType":"getNwkInfoReq","payload":{},"transactionId":%d}}`, transactionId)), res)
	})
}

func FuzzGetNetworkInfoUnmarshal(f *testing.F) {
	f.Add(string(GetNetworkInfoRequestEventType), 0)

	f.Fuzz(func(t *testing.T, eT string, transactionId int) {
		var e GetNetworkInfo
		value := []byte(fmt.Sprintf(`{"event":{"eventType":%q,"payload":{},"transactionId":%d}}`, eT, transactionId))

		err := json.Unmarshal(value, &e)

		if eventType(eT) != GetNetworkInfoRequestEventType {
			target := eventType(eT).Error()
			require.ErrorAs(t, err, &target)
			return
		}

		require.NoError(t, err)
		assert.Equal(t, transactionId, e.TransactionId)
	})
}
