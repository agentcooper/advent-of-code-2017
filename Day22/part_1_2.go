package Day22

import (
	"bufio"
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

type Coord struct {
	x int
	y int
}

type Offset struct {
	x int
	y int
}

var Up Offset = Offset{x: 0, y: -1}
var Down Offset = Offset{x: 0, y: +1}
var Left Offset = Offset{x: -1, y: 0}
var Right Offset = Offset{x: +1, y: 0}

var clockwise []Offset = []Offset{Up, Right, Down, Left}

type Virus struct {
	position  Coord
	direction Offset
}

func (v *Virus) move() {
	v.position.x += v.direction.x
	v.position.y += v.direction.y
}

func Turn(facing Offset, to Offset) Offset {
	current := 0

	for i, offset := range clockwise {
		if offset == facing {
			current = i
			break
		}
	}

	newIndex := 0
	switch to {
	case Left:
		newIndex = current - 1
	case Right:
		newIndex = current + 1
	}

	if newIndex < 0 {
		return clockwise[len(clockwise)-1]
	}

	return clockwise[(newIndex)%len(clockwise)]
}

func (v *Virus) turn(to Offset) {
	v.direction = Turn(v.direction, to)
}

type Cell int

const (
	Clean    Cell = 0
	Infected Cell = 1

	Weakened Cell = 2
	Flagged  Cell = 3
)

type Map map[Coord]Cell

// Solve solves the puzzle
func Solve(r io.Reader, bursts int, part Part) int {
	scanner := bufio.NewScanner(r)

	m := Map{}
	xMax := 0
	yMax := 0

	j := 0
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)

		for i, r := range line {
			var cell Cell
			switch r {
			case '.':
				cell = Clean
			case '#':
				cell = Infected
			}

			m[Coord{x: i, y: j}] = cell

			if i > xMax {
				xMax = i
			}
			if j > yMax {
				yMax = j
			}
		}

		j++
	}

	// fmt.Println(m)

	v := Virus{position: Coord{x: xMax / 2, y: yMax / 2}, direction: Up}
	infectedCount := 0

	for i := 0; i < bursts; i++ {
		if part == Part1 {
			if m[v.position] == Infected {
				v.turn(Right)
			} else {
				v.turn(Left)
			}

			if m[v.position] == Infected {
				m[v.position] = Clean
			} else {
				m[v.position] = Infected
				infectedCount++
			}

			v.move()
		}
		if part == Part2 {
			switch m[v.position] {
			case Infected:
				v.turn(Right)
			case Weakened:
				break
			case Flagged:
				v.turn(Left)
				v.turn(Left)
			default:
				v.turn(Left)
			}

			switch m[v.position] {
			case Infected:
				m[v.position] = Flagged
			case Weakened:
				m[v.position] = Infected
				infectedCount++
			case Flagged:
				m[v.position] = Clean
			default:
				m[v.position] = Weakened
			}

			v.move()
		}
	}

	return infectedCount
}
