package DayTemplate

import (
	"bufio"
	"fmt"
	"io"
)

// Part is either Part1 or Part2
type Part int

const (
	// Part1 is part 1
	Part1 Part = 1
	// Part2 is part 2
	Part2 Part = 2
)

// Solve solves the puzzle
func Solve(r io.Reader, part Part) int {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	return 0
}
