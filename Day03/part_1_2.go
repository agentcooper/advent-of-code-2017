package Day03

import "github.com/agentcooper/advent-of-code-2017/utils"

// Part is either Part1 or Part2
type Part int

const (
	// Part1 is part 1
	Part1 Part = 1
	// Part2 is part 2
	Part2 Part = 2
)

// Coord represents coordinate
type Coord struct {
	x int
	y int
}

// Direction represents coordinate offset
type Direction struct {
	x int
	y int
}

// Bounds represents coordinate rectangle
type Bounds struct {
	xMin int
	xMax int

	yMin int
	yMax int
}

func (c *Coord) step(d Direction) {
	c.x += d.x
	c.y += d.y
}

func (b *Bounds) add(c Coord) {
	b.xMin = utils.Min(b.xMin, c.x)
	b.xMax = utils.Max(b.xMax, c.x)

	b.yMin = utils.Min(b.yMin, c.y)
	b.yMax = utils.Max(b.yMax, c.y)
}

func (b Bounds) isOut(c Coord) bool {
	return c.y < b.yMin || c.y > b.yMax || c.x < b.xMin || c.x > b.xMax
}

func computeValue(grid map[Coord]int, c Coord) int {
	sum := 0
	for y := c.y - 1; y <= c.y+1; y++ {
		for x := c.x - 1; x <= c.x+1; x++ {
			sum += grid[Coord{x, y}]
		}
	}
	return sum
}

// Solve solves the puzzle based on part
func Solve(n int, part Part) int {
	clockwise := []Direction{
		Direction{1, 0},
		Direction{0, -1},
		Direction{-1, 0},
		Direction{0, 1},
	}

	currentDirection := 0
	bounds := Bounds{0, 0, 0, 0}

	grid := map[Coord]int{
		Coord{0, 0}: 1,
	}

	coord := Coord{0, 0}

	for i := 1; i < n; i++ {
		value := computeValue(grid, coord)

		// if part == Part1 {
		// 	fmt.Printf("'%d' at {%d, %d}, currentDirection = %d\n", i, x, y, currentDirection)
		// }
		// if part == Part2 {
		// 	fmt.Printf("'%d' at {%d, %d}, currentDirection = %d\n", value, x, y, currentDirection)
		// }

		if part == Part2 && value > n {
			return value
		}

		grid[coord] = value

		coord.step(clockwise[currentDirection])

		if bounds.isOut(coord) {
			bounds.add(coord)
			currentDirection = (currentDirection + 1) % len(clockwise)
		}
	}

	return utils.Abs(coord.x) + utils.Abs(coord.y)
}
