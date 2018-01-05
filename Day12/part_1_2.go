package Day12

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
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

// ParseLine parses the line
func ParseLine(s string) (int, []int) {
	parts := strings.Split(s, " <-> ")

	in, err := strconv.Atoi(parts[0])

	if err != nil {
		panic("Bad input (in)")
	}

	out, err := utils.StringsToInts(strings.Split(parts[1], ", "))

	if err != nil {
		panic("Bad input (out)")
	}

	return in, out
}

func bfs(start int, table map[int][]int, f func(int)) {
	seen := map[int]bool{}
	next := []int{start}

	for len(next) > 0 {
		var first int
		first, next = next[0], next[1:]

		if !seen[first] {
			for _, j := range table[first] {
				if !seen[j] {
					next = append(next, j)
				}
			}
		}

		f(first)
		seen[first] = true
	}
}

func findKeyByValue(m map[int]int, value int) (int, bool) {
	for k, v := range m {
		if v == value {
			return k, true
		}
	}

	return -1, false
}

// Solve solves the puzzle
func Solve(r io.Reader, part Part) int {
	scanner := bufio.NewScanner(r)

	table := map[int][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		in, out := ParseLine(line)
		table[in] = out
	}

	if part == Part1 {
		c := 0
		bfs(0, table, func(first int) {
			c++
		})
		return c - 1
	}

	if part == Part2 {
		group := map[int]int{}
		for k := range table {
			group[k] = 0
		}

		groupCount := 1

		for {
			start, found := findKeyByValue(group, 0)
			if !found {
				break
			}

			bfs(start, table, func(first int) {
				group[first] = groupCount
			})

			groupCount++
		}

		return groupCount - 1
	}

	panic("Invalid part")
}
