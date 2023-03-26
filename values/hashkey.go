package values

import (
	"encoding/hex"
	"fmt"
	"github.com/goccy/go-json"
	"strings"
)

type InvalidHashKey struct {
	hashKey HashKey
}

func (e *InvalidHashKey) Error() string {
	return fmt.Sprintf(`invalid hashKey "%s"`, e.hashKey)
}

type HashKey string

func (h *HashKey) UnmarshalJSON(bytes []byte) error {
	type hashKey HashKey

	if err := json.Unmarshal(bytes, (*hashKey)(h)); err != nil {
		return err
	}

	return h.Validate()
}

func (h *HashKey) MarshalJSON() ([]byte, error) {
	if err := h.Validate(); err != nil {
		return nil, err
	}

	type hashKey HashKey

	return json.Marshal((*hashKey)(h))
}

func (h *HashKey) Validate() error {
	if !strings.HasPrefix(string(*h), "0x") || strings.TrimLeft(string(*h), "0x") == "" {
		return &InvalidHashKey{*h}
	}

	if _, err := hex.DecodeString(strings.TrimLeft(string(*h), "0x")); err != nil {
		return &InvalidHashKey{*h}
	}

	return nil
}
