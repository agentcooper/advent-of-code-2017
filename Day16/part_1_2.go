package Day16

import (
	"fmt"
	"strconv"
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

type Program struct {
	name rune
	pos  int
}

type Spin struct {
	amount int
}

func (s *Spin) run(input *string) {
	n := s.amount

	q := *input

	*input = q[len(q)-n:] + q[0:len(q)-n]
}

type Exchange struct {
	i1 int
	i2 int
}

func (e *Exchange) run(input *string) {
	out := []rune(*input)

	q := *input

	out[e.i1] = rune(q[e.i2])
	out[e.i2] = rune(q[e.i1])

	*input = string(out)
}

type Partner struct {
	p1 rune
	p2 rune
}

func (p *Partner) run(input *string) {
	i1 := strings.IndexRune(*input, p.p1)
	i2 := strings.IndexRune(*input, p.p2)

	out := []rune(*input)

	q := *input

	out[i1] = rune(q[i2])
	out[i2] = rune(q[i1])

	*input = string(out)
}

type Command interface {
	run(input *string)
}

func parseCommands(s string) []Command {
	cs := []Command{}

	moves := strings.Split(s, ",")

	for _, move := range moves {
		if strings.HasPrefix(move, "s") {
			n, err := strconv.Atoi(move[1:])
			if err != nil {
				panic(fmt.Errorf("Bad input: %s", err))
			}

			cs = append(cs, &Spin{amount: n})
		}

		if strings.HasPrefix(move, "x") {
			parts := strings.Split(move[1:], "/")

			i1, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(fmt.Errorf("Bad input: %s", move))
			}

			i2, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(fmt.Errorf("Bad input: %s", move))
			}

			cs = append(cs, &Exchange{i1: i1, i2: i2})
		}

		if strings.HasPrefix(move, "p") {
			parts := strings.Split(move[1:], "/")
			cs = append(cs, &Partner{p1: rune(parts[0][0]), p2: rune(parts[1][0])})
		}
	}

	return cs
}

// Solve solves the puzzle
func Solve(s string, order string, part Part) string {
	commands := parseCommands(s)

	pp := order

	var danceTimes int
	if part == Part1 {
		danceTimes = 1
	}
	if part == Part2 {
		danceTimes = 1000000000
	}

	seen := map[string]bool{}
	didJump := false

	for i := 0; i < danceTimes; i++ {
		for _, command := range commands {
			command.run(&pp)
		}
		if seen[pp] && !didJump {
			i = danceTimes - danceTimes%i
			didJump = true
		} else {
			seen[pp] = true
		}
	}

	return pp
}
