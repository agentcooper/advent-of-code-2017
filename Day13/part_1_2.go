package Day13

import (
	"bufio"
	"fmt"
	"io"
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

type Layer struct {
	Depth int

	Range   int
	current int

	direction int
}

func NewLayer(Depth int, Range int) *Layer {
	return &Layer{Depth: Depth, Range: Range, direction: +1}
}

func (l *Layer) initWithTicks(ticks int) {
	l.current = 0
	opt := ticks % ((l.Range - 1) * 2)
	for j := 0; j < opt; j++ {
		l.move()
	}
}

func (l *Layer) move() {
	if l.current == 0 {
		l.direction = +1
	}

	if l.current == l.Range-1 {
		l.direction = -1
	}

	l.current = l.current + l.direction
}

func (l *Layer) severity() int {
	return l.Depth * l.Range
}

type Firewall struct {
	depth    map[int]*Layer
	maxDepth int
}

func NewFirewall() *Firewall {
	return &Firewall{depth: map[int]*Layer{}}
}

func (f *Firewall) initWithTicks(ticks int) {
	for k := range f.depth {
		f.depth[k].initWithTicks(ticks)
	}
}

func (f *Firewall) tick() {
	for k := range f.depth {
		f.depth[k].move()
	}
}

func (f *Firewall) reset() {
	for k := range f.depth {
		f.depth[k].current = 0
	}
}

func (f *Firewall) run(delay int) bool {
	f.initWithTicks(delay)

	for i := 0; i <= f.maxDepth; i++ {
		if f, ok := f.depth[i]; ok {
			isCaught := f.current == 0
			if isCaught {
				return false
			}
		}
		f.tick()
	}

	return true
}

// Solve solves the puzzle
func Solve(r io.Reader, part Part) int {
	scanner := bufio.NewScanner(r)

	firewall := NewFirewall()

	severity := 0

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		parts, err := utils.StringsToInts(strings.Split(line, ": "))
		if len(parts) != 2 || err != nil {
			panic(fmt.Errorf("Bad input: %s", line))
		}

		Depth := parts[0]
		Range := parts[1]

		firewall.depth[Depth] = NewLayer(Depth, Range)
		if Depth > firewall.maxDepth {
			firewall.maxDepth = Depth
		}
	}

	if part == Part1 {
		for i := 0; i <= firewall.maxDepth; i++ {
			if f, ok := firewall.depth[i]; ok {
				isCaught := f.current == 0
				if isCaught {
					fmt.Println("Caught", f.Depth)
					// caught
					severity += f.severity()
				}
			}
			firewall.tick()
		}

		return severity
	}

	for part == Part2 {
		for j := 0; j < 10000000; j++ {
			success := firewall.run(j)
			if success {
				return j
			}
		}
		panic("Can't find right delay")
	}

	panic("Unknown part")

}
