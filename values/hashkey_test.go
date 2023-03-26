package values

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashKey_Validate(t *testing.T) {
	tests := [...]struct {
		name    string
		h       HashKey
		valid   bool
		wantErr error
	}{
		{},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.Validate(); tt.valid {
				assert.ErrorIsf(t, err, tt.wantErr, "Validate() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
