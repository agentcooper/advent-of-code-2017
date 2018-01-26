package Day24

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/agentcooper/advent-of-code-2017/utils"
)

// Part is either Part1 or Part2
type Part int

const (
	// Part1 is part 1
	Part1 Part = 1
	// Part2 is part 2
	Part2 Part = 2
)

type Component struct {
	a int
	b int
}

func ParseComponent(s string) Component {
	ns, err := utils.StringsToInts(strings.Split(s, "/"))
	if err != nil {
		panic(fmt.Errorf("Bad input: %s", s))
	}

	return Component{a: ns[0], b: ns[1]}
}

func add(target map[int][]Component, k int, c Component) {
	if target[k] == nil {
		target[k] = []Component{}
	}

	target[k] = append(target[k], c)
}

func markAsSeen(m map[Component]bool, seen Component) map[Component]bool {
	out := map[Component]bool{}

	for k := range m {
		if k == seen {
			out[k] = true
		} else {
			out[k] = m[k]
		}
	}

	return out
}

func concat(list []Component, c Component) []Component {
	out := make([]Component, len(list)+1)
	copy(out, list)
	out[len(list)] = c
	return out
}

// too much params!
func run(c Component, port int, lookup map[int][]Component, seen map[Component]bool, strength int, length int, f func(int, int)) {
	for _, lc := range lookup[port] {
		if !seen[lc] {
			newStrength := strength + lc.a + lc.b
			f(length, newStrength)

			if lc.a == port {
				run(lc, lc.b, lookup, markAsSeen(seen, lc), newStrength, length+1, f)
			} else {
				run(lc, lc.a, lookup, markAsSeen(seen, lc), newStrength, length+1, f)
			}
		}
	}
}

// Solve solves the puzzle
func Solve(r io.Reader, part Part) int {
	scanner := bufio.NewScanner(r)

	lookup := map[int][]Component{}
	all := map[Component]bool{}

	for scanner.Scan() {
		line := scanner.Text()
		c := ParseComponent(line)

		if _, ok := all[c]; ok {
			panic("Duplicate component")
		}

		all[c] = false

		add(lookup, c.a, c)
		if c.a != c.b {
			add(lookup, c.b, c)
		}
	}

	maxStrength := 0

	longest := 0
	strengthByLength := map[int]int{}

	for _, c := range lookup[0] {
		run(c, c.b, lookup, markAsSeen(all, c), c.a+c.b, 1, func(length int, strength int) {
			if part == Part1 {
				if strength > maxStrength {
					maxStrength = strength
				}
			}

			if part == Part2 {
				if length >= longest {
					longest = length
					if strength > strengthByLength[longest] {
						strengthByLength[longest] = strength
					}
				}
			}
		})
	}

	if part == Part1 {
		return maxStrength
	}

	if part == Part2 {
		return strengthByLength[longest]
	}

	panic("Unknown part")
}
