package management

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnomaly(t *testing.T) {
	t.Run("CheckIP", func(t *testing.T) {
		isBlocked, err := m.Anomaly.CheckIP("1.1.1.1")
		assert.NoError(t, err)
		assert.False(t, isBlocked, "IP should not be blocked")
	})

	t.Run("UnblockIP", func(t *testing.T) {
		err := m.Anomaly.UnblockIP("1.1.1.1")
		assert.NoError(t, err)
	})
}
