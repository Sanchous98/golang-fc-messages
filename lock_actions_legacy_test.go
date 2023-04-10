package messages

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func FuzzLockResponseLegacyMarshal(f *testing.F) {
	f.Add(fmt.Sprintf(`{"state":{"%s":{}}}`, LockActionOpenEventType), string(LockActionOpenEventType), "no error")
	f.Add(fmt.Sprintf(`{"state":{"%s":{"recloseDelay":5}}}`, LockActionAutoEventType), string(LockActionAutoEventType), "no error")
	f.Add(fmt.Sprintf(`{"state":{"%s":{}}}`, LockActionCloseEventType), string(LockActionCloseEventType), "no error")
	f.Add(fmt.Sprintf(`{"error":{"%s":{}}}`, LockActionOpenEventType), string(LockActionOpenEventType), "no state error")
	f.Add(fmt.Sprintf(`{"state":{"%s":{}}`, LockActionOpenEventType), string(LockActionOpenEventType), "broken json")

	f.Fuzz(func(t *testing.T, jsonEvent string, state string, testResult string) {
		test := []byte(jsonEvent)
		var err error
		var e *LegacyLockEvent

		switch eventType(state) {
		case LockActionOpenEventType, LockActionCloseEventType, LockActionAutoEventType:
			e = &LegacyLockEvent{State: eventType(state)}
		default:
			target := eventType(state).Error()
			require.ErrorAs(t, err, &target)
			return
		}

		err = json.Unmarshal(test, e)
		if err != nil {
			target := eventType(state).Error()
			require.ErrorAs(t, err, &target)
			return
		}

		encodedEvent, _ := json.Marshal(e)
		assert.Equal(t, string(test), string(encodedEvent))
		require.NoError(t, err)
	})
}
