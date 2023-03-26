package values

import (
	"github.com/goccy/go-json"
	"time"
)

type Timestamp time.Time

func (t *Timestamp) Validate() error { return nil }

func (t *Timestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal((*time.Time)(t).Unix())
}

func (t *Timestamp) UnmarshalJSON(bytes []byte) error {
	var i int64

	if err := json.Unmarshal(bytes, &i); err != nil {
		return err
	}

	*t = (Timestamp)(time.Unix(i, 0))

	return nil
}
