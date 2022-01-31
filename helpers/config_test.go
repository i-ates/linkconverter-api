package helpers

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"runtime"
	"testing"
)

func TestNewConfig(t *testing.T) {
	t.Run("GetPortConfig", func(t *testing.T) {
		config := NewConfig()

		assert.NotZero(t, config)
		assert.Equal(t, "2626", config.Port)
	})
}

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}
