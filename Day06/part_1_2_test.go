package Day06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	input := `0 2 7 0`

	assert.Equal(t, Solve(input, Part1), 5)
}

func TestSolvePart2(t *testing.T) {
	input := `0 2 7 0`

	assert.Equal(t, Solve(input, Part2), 4)
}

func TestSolvePart1Input(t *testing.T) {
	input := `5	1	10	0	1	7	13	14	3	12	8	10	7	12	0	6`

	assert.Equal(t, Solve(input, Part1), 5042)
}

func TestSolvePart2Input(t *testing.T) {
	input := `5	1	10	0	1	7	13	14	3	12	8	10	7	12	0	6`

	assert.Equal(t, Solve(input, Part2), 1086)
}
