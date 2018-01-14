package Day23

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1Input(t *testing.T) {
	file, err := os.Open("./input.txt")
	assert.Nil(t, err)
	assert.Equal(t, Solve(file, Part1), 3969)
}

func TestSolvePart2Input(t *testing.T) {
	file, err := os.Open("./input.txt")
	assert.Nil(t, err)

	out, err := os.Create("./generated-raw/generated-raw.go")
	defer out.Close()
	assert.Nil(t, err)

	GoProgram(file, out)

	// see Generated/generated.go
	assert.Equal(t, 917, 917)
}
