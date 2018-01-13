package Day19

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	file, err := os.Open("./test-input.txt")
	assert.Nil(t, err)
	output := Solve(file)
	assert.Equal(t, output.letters, "ABCDEF")
}

func TestSolvePart2(t *testing.T) {
	file, err := os.Open("./test-input.txt")
	assert.Nil(t, err)
	output := Solve(file)
	assert.Equal(t, output.steps, 38)
}

func TestSolvePart1Input(t *testing.T) {
	file, err := os.Open("./input.txt")
	assert.Nil(t, err)
	output := Solve(file)
	assert.Equal(t, output.letters, "BPDKCZWHGT")
}

func TestSolvePart2Input(t *testing.T) {
	file, err := os.Open("./input.txt")
	assert.Nil(t, err)
	output := Solve(file)
	assert.Equal(t, output.steps, 17728)
}
