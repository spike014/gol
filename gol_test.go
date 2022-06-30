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