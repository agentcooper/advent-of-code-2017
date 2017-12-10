package Day06

import (
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

func findLargestIndex(ns []int) int {
	max := 0
	index := -1

	for i, n := range ns {
		if n > max {
			max = n
			index = i
		}
	}

	return index
}

func makeKey(ns []int) string {
	s := ""
	for _, n := range ns {
		s += string(n)
	}
	return s
}

// MaxIter is set to protect against infinite loop
const MaxIter = 10000

// Loop stores loop information
type Loop struct {
	key   string
	cycle int
}

// Solve solves the puzzle based on part
func Solve(s string, part Part) int {

	ns, err := utils.IntFields(&s)
	if err != nil {
		panic("Bad input")
	}

	// fmt.Println("Got input")
	// fmt.Println(ns)

	seen := map[string]bool{}
	var foundLoop *Loop

	cycle := 0
	for {
		lIndex := findLargestIndex(ns)
		// fmt.Printf("Memory bank with most blocks at index: %d (%d blocks)\n, now reallocating...", lIndex, ns[lIndex])

		offset := 1
		blocks := ns[lIndex]
		ns[lIndex] = 0

		for j := blocks; j > 0; j-- {
			nextIndex := (lIndex + offset) % len(ns)
			ns[nextIndex]++
			offset++
		}

		// fmt.Println("Reallocating done")
		// fmt.Println(ns)

		cycle++
		// fmt.Println("End cycle", cycle)

		key := makeKey(ns)

		if seen[key] {
			if part == Part1 {
				return cycle
			}

			if part == Part2 {
				if foundLoop != nil && key == foundLoop.key {
					return cycle - foundLoop.cycle
				}
				if foundLoop == nil {
					foundLoop = &Loop{key, cycle}
				}
			}
		}
		seen[key] = true

		if cycle > MaxIter {
			panic("Max iterations reached, exiting")
		}
	}
}
