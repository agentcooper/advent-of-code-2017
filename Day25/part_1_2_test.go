package Day25

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	input := `Begin in state A.
Perform a diagnostic checksum after 6 steps.

In state A:
	If the current value is 0:
		- Write the value 1.
		- Move one slot to the right.
		- Continue with state B.
	If the current value is 1:
		- Write the value 0.
		- Move one slot to the left.
		- Continue with state B.

In state B:
	If the current value is 0:
		- Write the value 1.
		- Move one slot to the left.
		- Continue with state A.
	If the current value is 1:
		- Write the value 1.
		- Move one slot to the right.sdfgs
		- Continue with state A.`

	assert.Equal(t, Solve(strings.NewReader(input), Part1), 3)
}

func TestSolvePart1Input(t *testing.T) {
	file, err := os.Open("./input.txt")
	assert.Nil(t, err)
	assert.Equal(t, Solve(file, Part1), 1)
}

func TestSolvePart2Input(t *testing.T) {
	file, err := os.Open("./input.txt")
	assert.Nil(t, err)

	assert.Equal(t, Solve(file, Part2), 0)
}
