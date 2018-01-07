package Day18

import (
	"bufio"
	"fmt"
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

type State struct {
	regs map[string]int
	ic   int

	queue []int
	ready bool

	i int

	wait       bool
	sndCounter int
	receiver   *Program
	naive      bool
}

type Program struct {
	ops   []Op
	state State
}

func (p *Program) isDone() bool {
	if p.state.i > 100000 {
		panic("Infinite loop protection")
	}

	return p.state.ic >= len(p.ops) || p.state.ready
}

func (p *Program) step() int {
	currentOp := p.ops[p.state.ic]

	out, didModifyIc := currentOp.eval(&p.state)

	if !didModifyIc {
		p.state.ic++
	}

	p.state.i++

	return out
}

func (p *Program) run() (int, error) {
	for !p.isDone() {
		out := p.step()

		if p.state.ready {
			return out, nil
		}
	}

	return -1, fmt.Errorf("rcv was not called")
}

type Register struct {
	name string
}

func (r Register) get(state *State) int {
	return state.regs[r.name]
}

type Number struct {
	value int
}

func (n Number) get(state *State) int {
	return n.value
}

type Value interface {
	get(state *State) int
}

func parseValue(s string) Value {
	n, err := strconv.Atoi(s)

	if err != nil {
		return Register{name: s}
	}

	return Number{value: n}
}

type Op interface {
	Name() string
	eval(state *State) (int, bool)
}

type UnaryOp struct {
	name string
	op   string
}

func (op UnaryOp) Name() string {
	return op.name
}

func (u UnaryOp) eval(s *State) (int, bool) {
	switch u.name {
	case "snd":
		value := parseValue(u.op).get(s)

		q := &s.receiver.state.queue
		*q = append(*q, value)

		s.sndCounter++

		return value, false
	case "rcv":
		if s.naive {
			value := parseValue(u.op).get(s)
			if value > 0 {
				s.ready = true

				q := s.receiver.state.queue

				var first int
				first, q = q[len(q)-1], q[:len(q)-1]

				return first, false
			}
		} else {
			if len(s.queue) == 0 {
				s.wait = true
				return -1, true
			}

			first := s.queue[0]
			s.queue = s.queue[1:]

			s.regs[u.op] = first

			return first, false
		}

	}

	return -1, false
}

type BinaryOp struct {
	name string
	op1  string
	op2  string
}

func (op BinaryOp) Name() string {
	return op.name
}

func (b BinaryOp) eval(s *State) (int, bool) {
	didModifyIc := false

	switch b.name {
	case "set":
		s.regs[b.op1] = parseValue(b.op2).get(s)
		break
	case "add":
		s.regs[b.op1] += parseValue(b.op2).get(s)
		break
	case "mul":
		s.regs[b.op1] *= parseValue(b.op2).get(s)
		break
	case "mod":
		s.regs[b.op1] %= parseValue(b.op2).get(s)
		break
	case "jgz":
		value := parseValue(b.op1).get(s)

		if value > 0 {
			s.ic += parseValue(b.op2).get(s)

			didModifyIc = true
		}
	}

	return 0, didModifyIc
}

func parseOp(s string) Op {
	fields := strings.Fields(s)

	if len(fields) == 2 {
		return UnaryOp{name: fields[0], op: fields[1]}
	}

	if len(fields) == 3 {
		return BinaryOp{name: fields[0], op1: fields[1], op2: fields[2]}
	}

	panic(fmt.Sprintf("Bad input: %s", s))
}

// Solve solves the puzzle
func Solve(r io.Reader, part Part) int {
	scanner := bufio.NewScanner(r)

	ops := []Op{}

	for scanner.Scan() {
		line := scanner.Text()
		ops = append(ops, parseOp(line))
	}

	if part == Part1 {
		p := Program{ops: ops, state: State{regs: map[string]int{}, naive: true}}
		p.state.receiver = &p

		res, err := p.run()
		if err != nil {
			panic(err)
		}

		return res
	}

	if part == Part2 {
		p0 := Program{ops: ops, state: State{regs: map[string]int{"p": 0}}}
		p1 := Program{ops: ops, state: State{regs: map[string]int{"p": 1}}}

		p0.state.receiver = &p1
		p1.state.receiver = &p0

		for !p0.isDone() && !p1.isDone() {
			p0.step()
			p1.step()

			isDeadlock := p0.state.wait && p1.state.wait
			if isDeadlock {
				break
			}
		}

		return p1.state.sndCounter
	}

	panic("Unknown part")
}
