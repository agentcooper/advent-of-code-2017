package Day01

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, SolvePart1([]byte("1122")), 3)
	assert.Equal(t, SolvePart1([]byte("1111")), 4)
	assert.Equal(t, SolvePart1([]byte("1234")), 0)
	assert.Equal(t, SolvePart1([]byte("91212129")), 9)
}

func TestSolvePart1Input(t *testing.T) {
	b, err := ioutil.ReadFile("./input.txt")
	assert.Nil(t, err)
	assert.Equal(t, SolvePart1(b), 997)
}
