package messages

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func FuzzLockResponseLegacyMarshal(f *testing.F) {
	f.Add(fmt.Sprintf(`{"state":{"%s":{}}}`, StateOpen), StateOpen, "no error")
	f.Add(fmt.Sprintf(`{"state":{"%s":{"recloseDelay":5}}}`, StateAuto), StateAuto, "no error")
	f.Add(fmt.Sprintf(`{"state":{"%s":{}}}`, StateClosed), StateClosed, "no error")
	f.Add(fmt.Sprintf(`{"error":{"%s":{}}}`, StateOpen), StateOpen, "no state error")
	f.Add(fmt.Sprintf(`{"state":{"%s":{}}`, StateOpen), StateOpen, "broken json")

	f.Fuzz(func(t *testing.T, jsonEvent string, state string, testResult string) {
		test := []byte(jsonEvent)
		var err error
		var e *LegacyLockEvent

		switch state {
		case StateOpen:
			e = &LegacyLockEvent{State: StateOpen}
		case StateClosed:
			e = &LegacyLockEvent{State: StateClosed}
		case StateAuto:
			e = &LegacyLockEvent{State: StateAuto}
		default:
			target := invalidLockStatus(lockStatus(state))
			require.ErrorAs(t, err, &target)
			return
		}

		err = json.Unmarshal(test, e)
		if err != nil {
			target := invalidLockStatus(lockStatus(state))
			require.ErrorAs(t, err, &target)
			return
		}

		encodedEvent, _ := json.Marshal(e)
		assert.Equal(t, string(test), string(encodedEvent))
		require.NoError(t, err)
	})
}
