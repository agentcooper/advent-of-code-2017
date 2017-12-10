package Day07

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	input := `pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`

	assert.Equal(t, SolvePart1(strings.NewReader(input)), "tknk")
}

func TestSolvePart1Input(t *testing.T) {
	file, err := os.Open("./input.txt")
	assert.Nil(t, err)
	assert.Equal(t, SolvePart1(file), "vgzejbd")
}

func TestSolvePart2Input(t *testing.T) {
	file, err := os.Open("./input.txt")
	assert.Nil(t, err)

	assert.Equal(t, SolvePart2(file), 1226)
}
