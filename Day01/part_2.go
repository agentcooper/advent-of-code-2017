package Day01

import (
	"strconv"
	"strings"
)

// SolvePart2 solves the day 01, part 2
func SolvePart2(b []byte) int {
	input := strings.Split(strings.TrimSpace(string(b)), "")
	sum := 0

	l := len(input)
	hl := l / 2

	for i := 0; i < l; i++ {
		j := (i + hl) % l

		if input[i] == input[j] {
			d, err := strconv.Atoi(input[i])
			if err == nil {
				sum += d
			}
		}
	}

	return sum
}
