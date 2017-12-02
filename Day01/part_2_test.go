package Day01

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, SolvePart2([]byte("1212")), 6)
	assert.Equal(t, SolvePart2([]byte("1221")), 0)
	assert.Equal(t, SolvePart2([]byte("123425")), 4)
	assert.Equal(t, SolvePart2([]byte("123123")), 12)
	assert.Equal(t, SolvePart2([]byte("12131415")), 4)
}

func TestSolvePart2Input(t *testing.T) {
	b, err := ioutil.ReadFile("./input.txt")
	assert.Nil(t, err)
	assert.Equal(t, SolvePart2(b), 1358)
}
