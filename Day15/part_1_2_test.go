package Day15

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const Iterations1 = 40 * 1000000
const Iterations2 = 5 * 1000000

var Input1 PuzzleInput = PuzzleInput{genAStart: 289, genBStart: 629, iterations: Iterations1}
var Input2 PuzzleInput = PuzzleInput{genAStart: 289, genBStart: 629, iterations: Iterations2}

func TestSolvePart1a(t *testing.T) {
	assert.Equal(t, Solve(PuzzleInput{genAStart: 65, genBStart: 8921, iterations: 5}, Part1), 1)
}

func TestSolvePart1b(t *testing.T) {
	assert.Equal(t, Solve(PuzzleInput{genAStart: 65, genBStart: 8921, iterations: Iterations1}, Part1), 588)
}

func TestSolvePart1Input(t *testing.T) {
	assert.Equal(t, Solve(Input1, Part1), 638)
}

func TestSolvePart2a(t *testing.T) {
	assert.Equal(t, Solve(PuzzleInput{genAStart: 65, genBStart: 8921, iterations: 1056}, Part2), 1)
}

func TestSolvePart2Input(t *testing.T) {
	assert.Equal(t, Solve(Input2, Part2), 343)
}
