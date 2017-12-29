package Day11

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, Solve(`ne,ne,ne`, Part1), 3)
	assert.Equal(t, Solve(`ne,ne,sw,sw`, Part1), 0)
	assert.Equal(t, Solve(`ne,ne,s,s`, Part1), 2)
	assert.Equal(t, Solve(`se,sw,se,sw,sw`, Part1), 3)

	assert.Equal(t, Solve(`ne,ne,ne,se,se,sw,sw,sw,nw,ne`, Part1), 2)
}

func TestSolvePart1Input(t *testing.T) {
	b, err := ioutil.ReadFile("./input.txt")

	assert.Nil(t, err)
	assert.Equal(t, Solve(string(b), Part1), 761)
}

func TestSolvePart2Input(t *testing.T) {
	b, err := ioutil.ReadFile("./input.txt")

	assert.Nil(t, err)
	assert.Equal(t, Solve(string(b), Part2), 1542)
}
