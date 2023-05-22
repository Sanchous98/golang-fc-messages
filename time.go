package messages

import "github.com/goccy/go-json"

const TimeSyncEventType eventType = "timeSync"

type TimeSyncEvent struct {
	TransactionId uint32
}

func (t *TimeSyncEvent) MarshalJSON() ([]byte, error) {
	var e event

	e.EventType = TimeSyncEventType
	e.TransactionId = t.TransactionId

	return json.Marshal(&e)
}

func (t *TimeSyncEvent) UnmarshalJSON(bytes []byte) error {
	var e event

	if err := json.Unmarshal(bytes, &e); err != nil {
		return err
	}

	if e.EventType != TimeSyncEventType {
		return e.EventType.Error()
	}

	t.TransactionId = e.TransactionId

	return nil
}
