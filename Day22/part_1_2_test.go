package Day22

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTurn(t *testing.T) {
	assert.Equal(t, Turn(Up, Left), Left)
	assert.Equal(t, Turn(Up, Right), Right)
	assert.Equal(t, Turn(Left, Left), Down)
	assert.Equal(t, Turn(Down, Right), Left)
}

func TestSolvePart1(t *testing.T) {
	input := `..#
#..
...`

	assert.Equal(t, Solve(strings.NewReader(input), 10000, Part1), 5587)
}

func TestSolvePart2(t *testing.T) {
	input := `..#
#..
...`

	assert.Equal(t, Solve(strings.NewReader(input), 100, Part2), 26)
	assert.Equal(t, Solve(strings.NewReader(input), 10000000, Part2), 2511944)
}

func TestSolvePart1Input(t *testing.T) {
	file, err := os.Open("./input.txt")
	assert.Nil(t, err)
	assert.Equal(t, Solve(file, 10000, Part1), 5223)
}

func TestSolvePart2Input(t *testing.T) {
	file, err := os.Open("./input.txt")
	assert.Nil(t, err)
	assert.Equal(t, Solve(file, 10000000, Part2), 2511456)
}
