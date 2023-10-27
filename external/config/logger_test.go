package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	t.Run("should return a new logger instance", func(t *testing.T) {
		logger, err := NewLogger(true)

		assert.NotNil(t, logger)
		assert.Nil(t, err)
	})

	t.Run("should return a new logger instance with zap.NewNop()", func(t *testing.T) {
		logger, err := NewLogger(false)

		assert.NotNil(t, logger)
		assert.Nil(t, err)
	})
}
