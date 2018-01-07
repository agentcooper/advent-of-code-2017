package Day16

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	input := `s1,x3/4,pe/b`

	assert.Equal(t, Solve(input, "abcde", Part1), "baedc")
}

func TestSolvePart1Input(t *testing.T) {
	b, err := ioutil.ReadFile("./input.txt")
	assert.Nil(t, err)

	assert.Equal(t, Solve(string(b), "abcdefghijklmnop", Part1), "padheomkgjfnblic")
}

func TestSolvePart2Input(t *testing.T) {
	b, err := ioutil.ReadFile("./input.txt")
	assert.Nil(t, err)

	assert.Equal(t, Solve(string(b), "abcdefghijklmnop", Part2), "bfcdeakhijmlgopn")
}
