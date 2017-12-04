package Day04

import (
	"bufio"
	"io"
	"sort"
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

// IsValid checks if s containts valid password
func IsValid(s string, keyFunc func(string) string) bool {
	words := strings.Fields(s)

	seen := map[string]bool{}
	for _, word := range words {
		key := keyFunc(word)

		if _, ok := seen[key]; ok {
			return false
		}

		seen[key] = true
	}

	return true
}

// IdentityKey returns the input string
func IdentityKey(s string) string {
	return s
}

// AnagramKey returns the anagram key of the word
func AnagramKey(s string) string {
	chars := map[rune]int{}
	runes := []rune(s)

	for _, rune := range s {
		chars[rune]++
	}

	sort.Sort(utils.SortedRunes(runes))

	out := []byte{}
	for _, rune := range runes {
		out = strconv.AppendInt(out, int64(chars[rune]), 10)
		out = append(out, byte(rune))
	}

	return string(out)
}

// Solve solves the puzzle based on part
func Solve(r io.Reader, part Part) int {
	var keyFunc func(string) string

	if part == Part1 {
		keyFunc = IdentityKey
	}
	if part == Part2 {
		keyFunc = AnagramKey
	}

	fscanner := bufio.NewScanner(r)
	count := 0

	for fscanner.Scan() {
		password := fscanner.Text()
		if len(password) > 0 && IsValid(password, keyFunc) {
			count++
		}
	}

	return count
}
