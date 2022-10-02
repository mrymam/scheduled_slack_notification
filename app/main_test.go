package main_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MainSample(t *testing.T) {
	t.Run("init test", func(t *testing.T) {
		assert.Equal(t, true, true)
	})
}
