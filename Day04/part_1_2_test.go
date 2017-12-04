package Day04

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, IsValid("aa bb cc dd ee", IdentityKey), true)
	assert.Equal(t, IsValid("aa bb cc dd aa", IdentityKey), false)
	assert.Equal(t, IsValid("aa bb cc dd aaa", IdentityKey), true)
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, IsValid("abcde fghij", AnagramKey), true)
	assert.Equal(t, IsValid("abcde xyz ecdab", AnagramKey), false)
	assert.Equal(t, IsValid("a ab abc abd abf abj", AnagramKey), true)
	assert.Equal(t, IsValid("iiii oiii ooii oooi oooo", AnagramKey), true)
	assert.Equal(t, IsValid("oiii ioii iioi iiio", AnagramKey), false)
}

func TestSolvePart1Input(t *testing.T) {
	file, err := os.Open("./input.txt")
	assert.Nil(t, err)
	assert.Equal(t, Solve(file, Part1), 466)
}

func TestSolvePart2Input(t *testing.T) {
	file, err := os.Open("./input.txt")
	assert.Nil(t, err)
	assert.Equal(t, Solve(file, Part2), 251)
}
