package Day17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, Solve(3, 2017, Part1), 638)
}

func TestSolvePart1Input(t *testing.T) {
	assert.Equal(t, Solve(329, 2017, Part1), 725)
}

func TestSolvePart2Input(t *testing.T) {
	assert.Equal(t, Solve(329, 50000000, Part2), 27361412)
}
