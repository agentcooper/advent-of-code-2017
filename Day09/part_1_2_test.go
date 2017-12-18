package Day09

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, Solve(`{}`, Part1), 1)
	assert.Equal(t, Solve(`{{{}}}`, Part1), 6)
	assert.Equal(t, Solve(`{{},{}}`, Part1), 5)
	assert.Equal(t, Solve(`{{{},{},{{}}}}`, Part1), 16)
	assert.Equal(t, Solve(`{<a>,<a>,<a>,<a>}`, Part1), 1)
	assert.Equal(t, Solve(`{{<ab>},{<ab>},{<ab>},{<ab>}}`, Part1), 9)
	assert.Equal(t, Solve(`{{<!!>},{<!!>},{<!!>},{<!!>}}`, Part1), 9)
	assert.Equal(t, Solve(`{{<a!>},{<a!>},{<a!>},{<ab>}}`, Part1), 3)
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, Solve(`<>`, Part2), 0)
	assert.Equal(t, Solve(`<random characters>`, Part2), 17)
	assert.Equal(t, Solve(`<<<<>`, Part2), 3)
	assert.Equal(t, Solve(`<{!>}>`, Part2), 2)
	assert.Equal(t, Solve(`<!!>`, Part2), 0)
	assert.Equal(t, Solve(`<!!!>>`, Part2), 0)
	assert.Equal(t, Solve(`<{o"i!a,<{i<a>`, Part2), 10)
}

func TestSolvePart1Input(t *testing.T) {
	b, err := ioutil.ReadFile("./input.txt")

	assert.Nil(t, err)
	assert.Equal(t, Solve(string(b), Part1), 7640)
}

func TestSolvePart2Input(t *testing.T) {
	b, err := ioutil.ReadFile("./input.txt")

	assert.Nil(t, err)
	assert.Equal(t, Solve(string(b), Part2), 4368)
}
