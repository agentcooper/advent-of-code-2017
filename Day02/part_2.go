package Day02

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// IntFields is the same as strings.Fields, but returns ints
func IntFields(s *string) ([]int, error) {
	ss := strings.Fields(strings.TrimSpace(*s))

	var ns []int

	for _, r := range ss {
		n, err := strconv.Atoi(r)
		if err != nil {
			return ns, err
		}
		ns = append(ns, n)
	}

	return ns, nil
}

// FindTwoDivisible finds a, b in ns where a / b == 0
func FindTwoDivisible(ns []int) (int, int, error) {
	for i, a := range ns {
		for j, b := range ns {
			if i != j && a >= b && b != 0 && a%b == 0 {
				return a, b, nil
			}
		}
	}
	return 0, 0, fmt.Errorf("Not found")
}

// SolvePart2 solves day 02, part 1
func SolvePart2(r io.Reader) int {
	fscanner := bufio.NewScanner(r)
	sum := 0

	for fscanner.Scan() {
		s := fscanner.Text()

		ns, err := IntFields(&s)
		if err != nil {
			panic(err)
		}
		// fmt.Println(ns)

		a, b, err := FindTwoDivisible(ns)
		// fmt.Println("Found", a, b)

		if err != nil {
			panic(err)
		}

		sum += a / b
	}

	return sum
}
