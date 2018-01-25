package Day21

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	input := `../.# => ##./#../...
.#./..#/### => #..#/..../..../#..#`

	assert.Equal(t, Solve(strings.NewReader(input), 2, Part1), 12)
}

func TestSolvePart1Input(t *testing.T) {
	file, err := os.Open("./input.txt")
	assert.Nil(t, err)
	assert.Equal(t, Solve(file, 5, Part1), 179)
}

func TestSolvePart2Input(t *testing.T) {
	file, err := os.Open("./input.txt")
	assert.Nil(t, err)
	assert.Equal(t, Solve(file, 18, Part1), 2766750)
}
