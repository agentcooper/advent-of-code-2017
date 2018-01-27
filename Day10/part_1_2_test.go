package Day10

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	out := [][]int{}

	Partition([]int{1, 2, 3, 4, 5}, 3, func(i int, part []int) {
		out = append(out, part)
	})

	assert.EqualValues(t, out[0], []int{1, 2, 3})
	assert.EqualValues(t, out[1], []int{4, 5})
}

func TestGrab(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}

	assert.EqualValues(t, Grab(input, 1, 3), []int{2, 3, 4})
	assert.EqualValues(t, Grab(input, 3, 3), []int{4, 5, 1})
}

func TestSolvePart1Input(t *testing.T) {
	file, err := os.Open("./input.txt")
	assert.Nil(t, err)
	assert.Equal(t, SolvePart1(file), 2928)
}

func TestSolvePart2Input(t *testing.T) {
	file, err := os.Open("./input.txt")
	assert.Nil(t, err)

	assert.Equal(t, SolvePart2(file), "0c2f794b2eb555f7830766bf8fb65a16")
}
