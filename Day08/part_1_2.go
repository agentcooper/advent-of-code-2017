package Day08

import (
	"bufio"
	"io"
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

// State represents current state
type State map[string]int

// Reference examples: 'a', 'xyz'
type Reference struct {
	name string
}

func (r Reference) get(state State) int {
	return state[r.name]
}

// Number examples: 42, -7
type Number struct {
	value int
}

func (n Number) get(state State) int {
	return n.value
}

// Value provides get method
type Value interface {
	get(state State) int
}

// Condition represents boolean expression of relation between 2 values
type Condition struct {
	left  Value
	op    string
	right Value
}

func (c Condition) evaluate(state State) bool {
	l := c.left.get(state)
	r := c.right.get(state)

	switch c.op {
	case "<":
		return l < r
	case ">":
		return l > r
	case ">=":
		return l >= r
	case "<=":
		return l <= r
	case "!=":
		return l != r
	case "==":
		return l == r

	default:
		panic("Unknown operator")
	}
}

// Command represents input command
type Command struct {
	register  string
	action    string
	value     Value
	condition Condition
}

func (c Command) evaluate(state State) (int, bool) {
	condition := c.condition.evaluate(state)

	if !condition {
		return -1, false
	}

	switch c.action {
	case "inc":
		state[c.register] += c.value.get(state)
	case "dec":
		state[c.register] -= c.value.get(state)
	default:
		panic("Unknown operator")
	}

	return state[c.register], true
}

func parseValue(s string) Value {
	n, err := strconv.Atoi(s)

	if err != nil {
		return Reference{name: s}
	}

	return Number{value: n}
}

func parseCommand(s string) Command {
	parts := strings.Fields(s)

	if len(parts) != 7 {
		panic("Bad command")
	}

	return Command{
		register: parts[0],
		action:   parts[1],
		value:    parseValue(parts[2]),
		condition: Condition{
			left:  parseValue(parts[4]),
			op:    parts[5],
			right: parseValue(parts[6]),
		},
	}
}

// Solve solves the puzzle
func Solve(r io.Reader, part Part) int {
	registers := State{}
	maxValue := 0

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		command := parseCommand(line)

		// fmt.Println(command)

		n, ok := command.evaluate(registers)

		if part == Part2 {
			if ok && n > maxValue {
				maxValue = n
			}
		}
	}

	if part == Part1 {
		for _, value := range registers {
			if value > maxValue {
				maxValue = value
			}
		}
	}

	return maxValue
}
