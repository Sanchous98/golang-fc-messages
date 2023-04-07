package messages

import (
	"errors"
	"github.com/goccy/go-json"
)

const (
	StateOpen   = "lockActionOpen"
	StateClosed = "lockActionClose"
	StateAuto   = "lockActionAuto"
)

type LockState string

func (l LegacyLockEvent) Error() error { return errors.New("invalid lock event type ") }

type LegacyLockEvent struct {
	State   string
	Reclose uint8
}

func (l *LegacyLockEvent) UnmarshalJSON(bytes []byte) error {
	var ev legacyEvent

	if err := json.Unmarshal(bytes, &ev); err != nil {
		return err
	}

	var e map[string]map[string]float64

	if err := json.Unmarshal(ev.State, &e); err != nil {
		return err
	}

	if _, ok := e[StateOpen]; ok {
		l.State = StateOpen
		return nil
	}

	if _, ok := e[StateClosed]; ok {
		l.State = StateClosed
		return nil
	}

	if _, ok := e[StateAuto]; ok {
		l.State = StateAuto
		l.Reclose = uint8(e[StateAuto]["recloseDelay"])

		if l.Reclose == 0 {
			l.Reclose = 5
		}

		return nil
	}

	return errors.New("unknown lock state")
}

func (l *LegacyLockEvent) MarshalJSON() ([]byte, error) {
	state := map[string]uint8{}

	if l.State == StateAuto {
		if l.Reclose == 0 {
			l.Reclose = 5
		}

		state["recloseDelay"] = l.Reclose
	}

	return json.Marshal(map[string]any{
		"state": map[string]any{
			l.State: state,
		},
	})
}
