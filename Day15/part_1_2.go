package Day15

// Part is either Part1 or Part2
type Part int

const (
	// Part1 is part 1
	Part1 Part = 1
	// Part2 is part 2
	Part2 Part = 2
)

type PuzzleInput struct {
	genAStart  int
	genBStart  int
	iterations int
}

type Generator struct {
	state  int
	factor int

	multiples int
}

func (g *Generator) next() {
	const f = 2147483647

	for {
		g.state = (g.state * g.factor) % f

		if g.state%g.multiples == 0 {
			return
		}
	}
}

func matchBits(a int, b int) bool {
	mask := (1 << 16) - 1
	return (a & mask) == (b & mask)
}

// Solve solves the puzzle
func Solve(p PuzzleInput, part Part) int {

	if part == Part1 {
		gA := Generator{state: p.genAStart, factor: 16807, multiples: 1}
		gB := Generator{state: p.genBStart, factor: 48271, multiples: 1}

		count := 0

		for i := 0; i < p.iterations; i++ {
			gA.next()
			gB.next()

			if matchBits(gA.state, gB.state) {
				count++
			}
		}

		return count
	}

	if part == Part2 {
		gA := Generator{state: p.genAStart, factor: 16807, multiples: 4}
		gB := Generator{state: p.genBStart, factor: 48271, multiples: 8}

		count := 0

		for i := 0; i < p.iterations; i++ {
			gA.next()
			gB.next()

			if matchBits(gA.state, gB.state) {
				count++
			}
		}

		return count
	}

	panic("Unknown part")
}
