package values

import (
	"fmt"
	"github.com/goccy/go-json"
)

type InvalidRSSIValue struct {
	r RSSI
}

func (e *InvalidRSSIValue) Error() string {
	return fmt.Sprintf("invalid rssi value! Expected values: -110 - 0, got: %d", e.r)
}

type RSSI int8

func (r *RSSI) Validate() error {
	if *r > 0 || *r < -100 {
		return &InvalidRSSIValue{*r}
	}

	return nil
}

func (r *RSSI) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, (*int8)(r)); err != nil {
		return err
	}

	return r.Validate()
}

func (r *RSSI) MarshalJSON() ([]byte, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}

	return json.Marshal((*int8)(r))
}
