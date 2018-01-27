package Day14

import (
	"fmt"
	"strconv"

	"github.com/agentcooper/advent-of-code-2017/Day10"
)

// Part is either Part1 or Part2
type Part int

const (
	// Part1 is part 1
	Part1 Part = 1
	// Part2 is part 2
	Part2 Part = 2
)

func HashToBinary(hash string) string {
	out := ""
	for _, r := range hash {
		d, _ := strconv.ParseInt(string(r), 16, 64)
		out += fmt.Sprintf("%04b", d)
	}
	return out
}

type Coord struct {
	x int
	y int
}

type Cell struct {
	zone int
}

type Map map[Coord]*Cell

func (m Map) ClaimZone(coord Coord, zone int, currentZone *int) {
	if _, ok := m[coord]; !ok {
		return
	}

	if m[coord].zone > 0 {
		return
	}

	z := zone
	if zone == -1 {
		*currentZone++
		z = *currentZone
	}
	m[coord].zone = z

	m.ClaimZone(Coord{x: coord.x - 1, y: coord.y}, z, currentZone)
	m.ClaimZone(Coord{x: coord.x + 1, y: coord.y}, z, currentZone)
	m.ClaimZone(Coord{x: coord.x, y: coord.y - 1}, z, currentZone)
	m.ClaimZone(Coord{x: coord.x, y: coord.y + 1}, z, currentZone)
}

// Solve solves the puzzle
func Solve(input string, part Part) int {
	m := Map{}
	totalCount := 0

	for i := 0; i < 128; i++ {
		s := fmt.Sprintf("%s-%d", input, i)
		line := HashToBinary(Day10.KnotHash(s))

		for j, r := range line {
			if r == '1' {
				totalCount++
				m[Coord{x: i, y: j}] = &Cell{}
			}
		}
	}

	if part == Part1 {
		return totalCount
	}

	if part == Part2 {
		currentZone := 0

		for x := 0; x < 128; x++ {
			for y := 0; y < 128; y++ {
				m.ClaimZone(Coord{x: x, y: y}, -1, &currentZone)
			}
		}

		return currentZone
	}

	panic("Unknown part")
}
