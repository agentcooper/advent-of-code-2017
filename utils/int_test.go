package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	assert.Equal(t, Max(3, 4), 4)
	assert.Equal(t, Max(4, 3), 4)
	assert.Equal(t, Max(-3, 4), 4)
	assert.Equal(t, Max(3, -4), 3)
}
