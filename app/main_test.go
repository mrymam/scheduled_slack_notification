package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MainSample(t *testing.T) {
	t.Run("init test", func(t *testing.T) {
		assert.Equal(t, true, true)
	})
}
