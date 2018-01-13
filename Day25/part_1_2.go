package Day25

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

// Part is either Part1 or Part2
type Part int

const (
	// Part1 is part 1
	Part1 Part = 1
	// Part2 is part 2
	Part2 Part = 2
)

type Statement interface{}

type Program struct {
	begin string
	steps int
	state map[string]map[int][]Statement
}

type Write struct {
	value int
}

type Move struct {
	direction int
}

type Continue struct {
	state string
}

func (p *Program) run() int {
	state := p.begin
	tape := map[int]int{}
	cursor := 0

	for i := 0; i < p.steps; i++ {
		for _, s := range p.state[state][tape[cursor]] {
			switch v := s.(type) {
			case Write:
				tape[cursor] = v.value
				break
			case Move:
				cursor += v.direction
				break
			case Continue:
				state = v.state
			default:
				panic(fmt.Errorf("Unknown statement: %+v", v))
			}
		}

		// fmt.Println(tape)
	}

	count1 := 0
	for _, v := range tape {
		if v == 1 {
			count1++
		}
	}

	return count1
}

// Solve solves the puzzle
func Solve(r io.Reader, part Part) int {
	scanner := bufio.NewScanner(r)

	p := Program{state: map[string]map[int][]Statement{}}

	state := ""
	value := 0

	for scanner.Scan() {
		line := scanner.Text()

		var submatch []string

		begin := regexp.MustCompile(`Begin in state (\w+).`)
		submatch = begin.FindStringSubmatch(line)
		if len(submatch) > 0 {
			p.begin = submatch[1]
			continue
		}

		perform := regexp.MustCompile(`Perform a diagnostic checksum after (\d+) steps.`)
		submatch = perform.FindStringSubmatch(line)
		if len(submatch) > 0 {
			steps, err := strconv.Atoi(submatch[1])
			if err != nil {
				panic(fmt.Errorf("Bad input: %s", submatch[1]))
			}

			p.steps = steps
			continue
		}

		inState := regexp.MustCompile(`In state (\w+):`)
		submatch = inState.FindStringSubmatch(line)

		if len(submatch) > 0 {
			state = submatch[1]
			continue
		}

		valueEqual := regexp.MustCompile(`If the current value is (\d):`)
		submatch = valueEqual.FindStringSubmatch(line)
		if len(submatch) > 0 {
			v, err := strconv.Atoi(submatch[1])
			if err != nil {
				panic(fmt.Errorf("Bad input: %s", submatch[1]))
			}

			value = v
			continue
		}

		var statement Statement

		write := regexp.MustCompile(`Write the value (\d).`)
		submatch = write.FindStringSubmatch(line)
		if len(submatch) > 0 {
			v, err := strconv.Atoi(submatch[1])
			if err != nil {
				panic(fmt.Errorf("Bad input: %s", submatch[1]))
			}

			statement = Write{value: v}
		}

		move := regexp.MustCompile(`Move one slot to the (\w+).`)
		submatch = move.FindStringSubmatch(line)
		if len(submatch) > 0 {
			var direction int

			switch submatch[1] {
			case "left":
				direction = -1
				break
			case "right":
				direction = +1
				break
			}

			statement = Move{direction: direction}
		}

		continueCommand := regexp.MustCompile(`Continue with state (\w+).`)
		submatch = continueCommand.FindStringSubmatch(line)
		if len(submatch) > 0 {
			statement = Continue{state: submatch[1]}
		}

		if statement != nil {
			if p.state[state] == nil {
				p.state[state] = map[int][]Statement{}
			}

			if p.state[state][value] == nil {
				p.state[state][value] = []Statement{}
			}

			p.state[state][value] = append(p.state[state][value], statement)
		}
	}

	// fmt.Println(p)

	return p.run()
}
