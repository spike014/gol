package gol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetExcutePath(t *testing.T) {
	t.Run("pos", func(t *testing.T) {
		path, err := getExcutePath()
		assert.Nil(t, err)
		assert.NotEqual(t, "", path)
	})
}

func TestGetLogPath(t *testing.T) {
	t.Run("pos", func(t *testing.T) {
		path := getLogPath()
		assert.NotEqual(t, "", path)
	})
}

func TestPathExists(t *testing.T) {
	t.Run("neg", func(t *testing.T) {
		exist := pathExists("/home")
		assert.False(t, exist)
	})
}
