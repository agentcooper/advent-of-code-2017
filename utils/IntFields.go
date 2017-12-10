package utils

import (
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
