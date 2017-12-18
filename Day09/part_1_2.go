package Day09

// Part is either Part1 or Part2
type Part int

const (
	// Part1 is part 1
	Part1 Part = 1
	// Part2 is part 2
	Part2 Part = 2
)

// Solve solves the puzzle
func Solve(s string, part Part) int {
	totalScore := 0
	score := 1

	isInGarbage := false
	ignore := false

	garbageCount := 0

	for _, b := range s {
		if ignore {
			ignore = false
			continue
		}

		if b == '>' {
			isInGarbage = false
			ignore = false
		}

		if isInGarbage && b != '!' {
			garbageCount++
		}

		if !isInGarbage && b == '{' {
			totalScore += score
			score++
		}

		if !isInGarbage && b == '}' {
			score--
		}

		if b == '<' {
			isInGarbage = true
		}

		if isInGarbage && b == '!' {
			ignore = true
		}
	}

	if part == Part2 {
		return garbageCount
	}

	return totalScore
}
