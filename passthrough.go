package messages

type PassThroughEvent struct {
	Event struct {
		EventType string `json:"eventType"`
		Payload   struct {
			CommandId int `json:"commandId"`
			Data      []struct {
				Type string `json:"type,omitempty"`
				Data string `json:"data"`
			} `json:"data"`
		} `json:"payload"`
		Status        int `json:"status"`
		TransactionId int `json:"transactionId"`
	} `json:"event"`
}
