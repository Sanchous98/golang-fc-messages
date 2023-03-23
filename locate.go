package messages

import "github.com/goccy/go-json"

const LocateRequestEventType eventType = "locateReq"

type LocateRequest struct {
	TransactionId int
}

func (r *LocateRequest) MarshalJSON() ([]byte, error) {
	var e event
	e.EventType = LocateRequestEventType
	e.TransactionId = r.TransactionId

	return json.Marshal(&e)
}

func (r *LocateRequest) UnmarshalJSON(raw []byte) error {
	var e event

	if err := json.Unmarshal(raw, &e); err != nil {
		return err
	}

	if e.EventType != LocateRequestEventType {
		return e.EventType.Error()
	}

	r.TransactionId = e.TransactionId

	return nil
}
