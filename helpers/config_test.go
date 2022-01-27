package helpers

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"linkconverter-api/libs/logging"
	"testing"
)

func TestGetAppName(t *testing.T) {
	t.Run("GetAppName", func(t *testing.T) {
		appName := GetAppName()

		assert.Equal(t, "RohanApi-test", appName)
	})
}

func TestMain(m *testing.M) {
	logConfig := logging.LoggerConfig{
		Config:           zap.NewProductionConfig(),
		ContextFieldFunc: nil}

	logging.GetLogger(logConfig)

	m.Run()
}
