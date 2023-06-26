package messages

import (
	"errors"
	"github.com/goccy/go-json"
)

const (
	LockActionOpenEventType  eventType = "lockActionOpen"
	LockActionCloseEventType eventType = "lockActionClose"
	LockActionAutoEventType  eventType = "lockActionAuto"
)

type eventType string

func (e eventType) Error() error { return errors.New("invalid lock event type " + string(e)) }

type LegacyLockEvent struct {
	State   eventType
	Reclose uint8
}

func (l *LegacyLockEvent) UnmarshalJSON(bytes []byte) error {
	var ev legacyEvent

	if err := json.Unmarshal(bytes, &ev); err != nil {
		return err
	}

	var e map[eventType]map[string]float64

	if err := json.Unmarshal(ev.State, &e); err != nil {
		return err
	}

	if _, ok := e[LockActionOpenEventType]; ok {
		l.State = LockActionOpenEventType
		return nil
	}

	if _, ok := e[LockActionCloseEventType]; ok {
		l.State = LockActionCloseEventType
		return nil
	}

	if _, ok := e[LockActionAutoEventType]; ok {
		l.State = LockActionAutoEventType
		l.Reclose = uint8(e[LockActionAutoEventType]["recloseDelay"])

		if l.Reclose == 0 {
			l.Reclose = 5
		}

		return nil
	}

	for k := range e {
		return k.Error()
	}

	return nil
}

func (l *LegacyLockEvent) MarshalJSON() ([]byte, error) {
	state := map[string]uint8{}

	if l.State == LockActionAutoEventType {
		if l.Reclose == 0 {
			l.Reclose = 5
		}

		state["recloseDelay"] = l.Reclose
	}

	return json.Marshal(map[string]interface{}{
		"state": map[eventType]interface{}{
			l.State: state,
		},
	})
}
