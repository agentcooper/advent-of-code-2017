package Day01

import (
	"strconv"
	"strings"
)

// SolvePart1 solves the day 01, part 1
func SolvePart1(b []byte) int {
	input := strings.Split(strings.TrimSpace(string(b)), "")
	sum := 0

	l := len(input)

	for i := 0; i < l; i++ {
		j := (i + 1) % l

		if input[i] == input[j] {
			d, err := strconv.Atoi(input[i])
			if err == nil {
				sum += d
			}
		}
	}

	return sum
}
