package Day19

import (
	"bufio"
	"io"
	"unicode"
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

type Direction struct {
	x int
	y int
}

type Output struct {
	letters string
	steps   int
}

func (c *Coord) move(d Direction) {
	c.x += d.x
	c.y += d.y
}

func (c Coord) add(d Direction) Coord {
	return Coord{x: c.x + d.x, y: c.y + d.y}
}

var top Direction = Direction{x: 0, y: -1}
var right Direction = Direction{x: +1, y: 0}
var bottom Direction = Direction{x: 0, y: +1}
var left Direction = Direction{x: -1, y: 0}

func isHorizontal(d Direction) bool {
	return d == left || d == right
}

func isVertical(d Direction) bool {
	return d == top || d == bottom
}

// Solve solves the puzzle
func Solve(r io.Reader) Output {
	scanner := bufio.NewScanner(r)

	m := map[Coord]rune{}
	xMax := 0
	yMax := 0

	s := []rune{}
	steps := 0

	// fill input
	j := 0
	for scanner.Scan() {
		line := scanner.Text()

		for i, r := range line {
			if !unicode.IsSpace(r) {
				c := Coord{x: i, y: j}
				m[c] = r
			}

			if i > xMax {
				xMax = i
			}
		}

		if j > yMax {
			yMax = j
		}

		j++
	}

	// find start
	pos := Coord{x: 0, y: 0}
	d := Direction{x: 0, y: 1}

	for i := 0; i < xMax; i++ {
		pos.x = i
		if m[pos] == '|' {
			break
		}
	}

Loop:
	for {
		pos.move(d)
		steps++
		switch rr := m[pos]; rr {
		case '|':
			break
		case '-':
			break
		case '+':
			if isVertical(d) {
				if _, ok := m[pos.add(left)]; ok {
					d = left
					break
				}
				if _, ok := m[pos.add(right)]; ok {
					d = right
					break
				}
			}
			if isHorizontal(d) {
				if _, ok := m[pos.add(top)]; ok {
					d = top
					break
				}
				if _, ok := m[pos.add(bottom)]; ok {
					d = bottom
					break
				}
			}
			break
		default:
			if unicode.IsLetter(rr) {
				s = append(s, rr)
			} else {
				break Loop
			}
		}
	}

	return Output{letters: string(s), steps: steps}
}
