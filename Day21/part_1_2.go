package Day21

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Part is either Part1 or Part2
type Part int

const (
	// Part1 is part 1
	Part1 Part = 1
	// Part2 is part 2
	Part2 Part = 2
)

type Rule struct {
	from Grid
	to   Grid
}

func parseGrid(s string) Grid {
	grid := Grid{}
	parts := strings.Split(s, "/")
	for _, s := range parts {
		grid = append(grid, []rune(s))
	}
	return grid
}

func parseRule(s string) Rule {
	parts := strings.Split(s, " => ")

	if len(parts) != 2 {
		panic(fmt.Errorf("Bad input: %s", s))
	}

	return Rule{from: parseGrid(parts[0]), to: parseGrid(parts[1])}
}

func RotateClockwise(s Grid) Grid {
	size := len(s)
	output := NewGrid(size)

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			output[y][x] = s[size-x-1][y]
		}
	}

	return output
}

func FlipHorizontal(s Grid) Grid {
	size := len(s)
	output := NewGrid(size)

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			output[len(s)-1-y][x] = s[y][x]
		}
	}

	return output
}

func ToString(s [][]rune) string {
	out := ""
	size := len(s)

	for y := 0; y < size; y++ {
		out += string(s[y])
		if y != size-1 {
			out += "\n"
		}
	}

	return out
}

func ToSlashes(s [][]rune) string {
	out := ""
	size := len(s)

	for y := 0; y < size; y++ {
		out += string(s[y])
		if y != size-1 {
			out += "/"
		}
	}

	return out
}

type Grid [][]rune

func NewGrid(size int) Grid {
	output := [][]rune{}
	for y := 0; y < size; y++ {
		output = append(output, make([]rune, size))
	}
	return output
}

func GrabGrid(source Grid, xStart int, yStart int, size int) Grid {
	output := NewGrid(size)

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			output[y][x] = source[yStart+y][xStart+x]
		}
	}

	return output
}

func CopyGrid(to Grid, from Grid, xStart int, yStart int) {
	size := len(from)
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			to[yStart+y][xStart+x] = from[y][x]
		}
	}
}

func CountPixels(grid Grid) int {
	count := 0
	size := len(grid)
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {

			if grid[y][x] == '#' {
				count++
			}
		}
	}
	return count
}

func perform(picture Grid, rules map[string]Grid, from int, to int) Grid {
	size := len(picture)

	out := NewGrid(size / from * to)

	for y := 0; y < size/from; y++ {
		for x := 0; x < size/from; x++ {
			part := GrabGrid(picture, x*from, y*from, from)

			replacement, ok := rules[ToString(part)]
			if !ok {
				panic(fmt.Errorf("No match, picture: %s", ToSlashes(part)))
			}

			CopyGrid(out, replacement, x*to, y*to)
		}
	}

	return out
}

// Solve solves the puzzle
func Solve(r io.Reader, iterations int, part Part) int {
	rules := map[string]Grid{}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		rule := parseRule(line)

		up := rule.from
		right := RotateClockwise(rule.from)
		down := RotateClockwise(RotateClockwise(rule.from))
		left := RotateClockwise(RotateClockwise(RotateClockwise(rule.from)))

		rules[ToString(up)] = rule.to
		rules[ToString(right)] = rule.to
		rules[ToString(down)] = rule.to
		rules[ToString(left)] = rule.to

		rules[ToString(FlipHorizontal(up))] = rule.to
		rules[ToString(FlipHorizontal(right))] = rule.to
		rules[ToString(FlipHorizontal(down))] = rule.to
		rules[ToString(FlipHorizontal(left))] = rule.to
	}

	picture := parseGrid(`.#./..#/###`)

	var out Grid

	for i := 0; i < iterations; i++ {
		size := len(picture)

		if size%3 == 0 {
			out = perform(picture, rules, 3, 4)
		}
		if size%2 == 0 {
			out = perform(picture, rules, 2, 3)
		}
		picture = out
	}

	return CountPixels(picture)
}
