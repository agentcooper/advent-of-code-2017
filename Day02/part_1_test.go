package Day02

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	s := `5 1 9 5
7 5 3
2 4 6 8`

	assert.Equal(t, SolvePart1(strings.NewReader(s)), 18)
}

func TestSolvePart1Input(t *testing.T) {
	file, err := os.Open("./input.txt")
	assert.Nil(t, err)
	assert.Equal(t, SolvePart1(file), 34925)
}
