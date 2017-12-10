package Day07

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
)

// Part is either Part1 or Part2
type Part int

const (
	// Part1 is part 1
	Part1 Part = 1
	// Part2 is part 2
	Part2 Part = 2
)

// Input holds input from each program
type Input struct {
	name   string
	weight int
	above  []string
}

func getInput(s string) Input {
	ProgramInput := regexp.MustCompile(`(\w+)\s+\((\d+)\)`)
	parts := regexp.MustCompile(`\s+->\s+`).Split(s, -1)

	submatch := ProgramInput.FindStringSubmatch(parts[0])
	if len(submatch) != 3 {
		panic("Bad input")
	}

	n, err := strconv.Atoi(submatch[2])
	if err != nil {
		panic("Bad input")
	}

	if len(parts) == 1 {
		return Input{
			name:   submatch[1],
			weight: n,
			above:  nil,
		}
	}

	if len(parts) == 2 {
		return Input{
			name:   submatch[1],
			weight: n,
			above:  regexp.MustCompile(`,\s+`).Split(parts[1], -1),
		}
	}

	panic("Bad input")
}

func incrLevel(level map[string]int, parent map[string]string, start string) {
	p := parent[start]

	if p != "" {
		level[p]++
		incrLevel(level, parent, p)
	}
}

func findMaxKey(m map[string]int) string {
	k := ""
	max := 0

	for key, n := range m {
		if n > max {
			max = n
			k = key
		}
	}

	return k
}

func getWeight(m map[string]Input, s string) int {
	w := m[s].weight
	for _, aboveName := range m[s].above {
		w += getWeight(m, aboveName)
	}
	return w
}

// Subtree stores name and weight
type Subtree struct {
	name   string
	weight int
}

func getWeights(m map[string]Input, s string) []Subtree {
	ws := []Subtree{}

	for _, aboveName := range m[s].above {
		w := getWeight(m, aboveName)
		ws = append(ws, Subtree{aboveName, w})
	}

	return ws
}

func findDifferentSubtree(s []Subtree) (*Subtree, int) {
	for _, item := range s {
		diff := item.weight - s[0].weight
		if diff > 0 {
			return &item, diff
		}
	}
	return nil, 0
}

func find(m map[string]Input, s string, backtrack int) int {
	ws := getWeights(m, s)
	st, diff := findDifferentSubtree(ws)

	if st != nil {
		return find(m, st.name, diff)
	}

	return m[s].weight - backtrack
}

func buildLookup(r io.Reader) (map[string]Input, map[string]string, map[string]int) {
	scanner := bufio.NewScanner(r)

	lookup := map[string]Input{}
	parent := map[string]string{}

	for scanner.Scan() {
		line := scanner.Text()
		input := getInput(line)

		lookup[input.name] = input

		for _, aboveName := range input.above {
			parent[aboveName] = input.name
		}
	}

	level := map[string]int{}

	for _, input := range lookup {
		if len(input.above) > 0 {
			level[input.name]++
		}

		if parent[input.name] != "" {
			incrLevel(level, parent, input.name)
		}
	}

	return lookup, parent, level
}

// SolvePart2 solves part 2
func SolvePart2(r io.Reader) int {
	lookup, _, level := buildLookup(r)

	rootKey := findMaxKey(level)

	return find(lookup, rootKey, 0)
}

// SolvePart1 solves part 1
func SolvePart1(r io.Reader) string {
	_, _, level := buildLookup(r)

	return findMaxKey(level)
}
