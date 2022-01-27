package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAppName(t *testing.T) {
	t.Run("GetAppName", func(t *testing.T) {
		appName := GetAppName()

		assert.Equal(t, "LinkConverterApi-local", appName)
	})
}
