package Day02

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart2(t *testing.T) {
	s := `5 9 2 8
9 4 7 3
3 8 6 5`

	assert.Equal(t, SolvePart2(strings.NewReader(s)), 9)
}

func TestSolvePart2Input(t *testing.T) {
	file, err := os.Open("./input.txt")
	assert.Nil(t, err)
	assert.Equal(t, SolvePart2(file), 221)
}
