package Day13

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTicks(t *testing.T) {
	l1 := NewLayer(0, 3)
	l2 := NewLayer(0, 3)

	ticks := 888

	for j := 0; j < ticks; j++ {
		l1.move()
	}

	l2.initWithTicks(ticks)

	assert.Equal(t, l1.current, l2.current)
}
func TestSolvePart1(t *testing.T) {
	input := `0: 3
1: 2
4: 4
6: 4`

	assert.Equal(t, Solve(strings.NewReader(input), Part1), 24)
}

func TestSolvePart2(t *testing.T) {
	input := `0: 3
1: 2
4: 4
6: 4`

	assert.Equal(t, Solve(strings.NewReader(input), Part2), 10)
}

func TestSolvePart1Input(t *testing.T) {
	file, err := os.Open("./input.txt")
	assert.Nil(t, err)
	assert.Equal(t, Solve(file, Part1), 1580)
}

func TestSolvePart2Input(t *testing.T) {
	file, err := os.Open("./input.txt")
	assert.Nil(t, err)

	assert.Equal(t, Solve(file, Part2), 3943252)
}
