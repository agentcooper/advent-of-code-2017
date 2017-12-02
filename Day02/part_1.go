package Day02

import (
	"bufio"
	"io"
	"math"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// SolvePart1 solves day 02, part 1
func SolvePart1(r io.Reader) int {
	fscanner := bufio.NewScanner(r)
	sum := 0

	for fscanner.Scan() {
		rowMax := 0
		rowMin := math.MaxInt64

		ss := strings.Fields(strings.TrimSpace(fscanner.Text()))

		for _, r := range ss {
			n, err := strconv.Atoi(r)
			if err != nil {
				panic(err)
			}

			rowMin = min(n, rowMin)
			rowMax = max(n, rowMax)
		}

		sum += rowMax - rowMin
	}

	return sum
}
