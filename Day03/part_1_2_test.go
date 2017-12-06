package Day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, Solve(1, Part1), 0)
	assert.Equal(t, Solve(12, Part1), 3)
	assert.Equal(t, Solve(23, Part1), 2)
	assert.Equal(t, Solve(1024, Part1), 31)
}

func TestSolvePart1Input(t *testing.T) {
	assert.Equal(t, Solve(277678, Part1), 475)
}

func TestSolvePart2Input(t *testing.T) {
	assert.Equal(t, Solve(277678, Part2), 279138)
}
