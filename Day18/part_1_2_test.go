package Day18

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	input := `set a 1
add a 2
mul a a
mod a 5
snd a
set a 0
rcv a
jgz a -1
set a 1
jgz a -2`

	assert.Equal(t, Solve(strings.NewReader(input), Part1), 4)
}

func TestSolvePart2(t *testing.T) {
	input := `snd 1
snd 2
snd p
rcv a
rcv b
rcv c
rcv d`

	assert.Equal(t, Solve(strings.NewReader(input), Part2), 3)
}

func TestSolvePart1Input(t *testing.T) {
	file, err := os.Open("./input.txt")
	assert.Nil(t, err)
	assert.Equal(t, Solve(file, Part1), 4601)
}

func TestSolvePart2Input(t *testing.T) {
	file, err := os.Open("./input.txt")
	assert.Nil(t, err)

	assert.Equal(t, Solve(file, Part2), 6858)
}
