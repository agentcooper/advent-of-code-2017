package Day11

import (
	"math"
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

// Coord is X,Y coord
type Coord struct {
	X int
	Y int
}

func (c Coord) distance(other Coord) int {
	return utils.Abs(c.X-other.X) + utils.Abs(c.Y-other.Y)
}

func (c Coord) equal(other Coord) bool {
	return c.distance(other) == 0
}

func (c Coord) add(other Coord) Coord {
	return Coord{c.X + other.X, c.Y + other.Y}
}

var clockwise = map[string]Coord{
	"ne": Coord{0, 1},
	"n":  Coord{-1, 1},
	"nw": Coord{-1, 0},
	"sw": Coord{0, -1},
	"s":  Coord{1, -1},
	"se": Coord{1, 0},
}

func steps(dest Coord) int {
	coord := Coord{0, 0}
	i := 0

	for !dest.equal(coord) {
		d := math.MaxInt64

		var m Coord

		for _, move := range clockwise {
			candidate := coord.add(move)

			if dest.distance(candidate) < d {
				d = dest.distance(candidate)
				m = move
			}
		}

		coord = coord.add(m)
		i++
	}

	return i
}

// Solve solves the puzzle
func Solve(s string, part Part) int {
	finalPos := Coord{0, 0}

	dirs := strings.Split(s, ",")

	max := 0

	for _, dir := range dirs {
		finalPos = finalPos.add(clockwise[dir])

		if part == Part2 {
			s := steps(finalPos)
			if s > max {
				max = s
			}
		}
	}

	if part == Part1 {
		return steps(finalPos)
	}

	if part == Part2 {
		return max
	}

	return steps(finalPos)
}
