package Day14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1Input(t *testing.T) {
	assert.Equal(t, Solve("vbqugkhl", Part1), 8148)
}

func TestSolvePart2Input(t *testing.T) {
	assert.Equal(t, Solve("vbqugkhl", Part2), 1180)
}
