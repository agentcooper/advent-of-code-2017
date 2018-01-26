package Day24

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	input := `0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10`

	assert.Equal(t, Solve(strings.NewReader(input), Part1), 31)
}

func TestSolvePart2(t *testing.T) {
	input := `0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10`

	assert.Equal(t, Solve(strings.NewReader(input), Part2), 19)
}

func TestSolvePart1Input(t *testing.T) {
	file, err := os.Open("./input.txt")
	assert.Nil(t, err)
	assert.Equal(t, Solve(file, Part1), 2006)
}

func TestSolvePart2Input(t *testing.T) {
	file, err := os.Open("./input.txt")
	assert.Nil(t, err)

	assert.Equal(t, Solve(file, Part2), 1994)
}
