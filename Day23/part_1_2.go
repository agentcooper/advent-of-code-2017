package Day23

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

	mulCounter int
}

type Program struct {
	ops   []BinaryOp
	state State
}

func (p *Program) isDone() bool {
	// if p.state.i > 1000000 {
	// 	panic("Infinite loop protection")
	// }

	return p.state.ic >= len(p.ops) || p.state.ready || p.state.i > 1000000
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

func (p *Program) run() {
	for !p.isDone() {
		p.step()
	}
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

type BinaryOp struct {
	name string
	op1  string
	op2  string
}

func (op BinaryOp) toString() string {
	return fmt.Sprintf("%s %s %s", op.name, op.op1, op.op2)
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
	case "sub":
		s.regs[b.op1] -= parseValue(b.op2).get(s)
		break
	case "mul":
		s.mulCounter++
		s.regs[b.op1] = parseValue(b.op1).get(s) * parseValue(b.op2).get(s)
		break
	case "jnz":
		value := parseValue(b.op1).get(s)

		if value != 0 {
			s.ic += parseValue(b.op2).get(s)
			didModifyIc = true
		}
	}

	return 0, didModifyIc
}

func parseOp(s string) BinaryOp {
	fields := strings.Fields(s)

	if len(fields) == 3 {
		return BinaryOp{name: fields[0], op1: fields[1], op2: fields[2]}
	}

	panic(fmt.Sprintf("Bad input: %s", s))
}

func (p *Program) toGoProgram(w io.Writer) {
	hasLabel := map[int]bool{}

	for i, op := range p.ops {
		absIndex := 0

		if op.name == "jnz" {
			absIndex = i + parseValue(op.op2).get(&p.state)
			hasLabel[absIndex] = true
		}
	}

	fmt.Fprintln(w, `
		// auto-generated
		package main

		import "fmt"

		func main() {
			a,b,c,d,e,f,g,h := 1,0,0,0,0,0,0,0
	`)

	for i, op := range p.ops {
		absIndex := 0

		prefix := ""
		if hasLabel[i] {
			prefix = fmt.Sprintf("L%d: ", i)
		}

		if op.name == "jnz" {
			absIndex = i + parseValue(op.op2).get(&p.state)

			if absIndex >= len(p.ops) {
				fmt.Fprintf(w, "%sif %s != 0 { goto END }\n", prefix, op.op1)
			} else {
				fmt.Fprintf(w, "%sif %s != 0 { goto L%d }\n", prefix, op.op1, absIndex)
			}

		} else if op.name == "set" {
			fmt.Fprintf(w, "%s%s = %s\n", prefix, op.op1, op.op2)
		} else if op.name == "mul" {
			fmt.Fprintf(w, "%s%s = %s * %s\n", prefix, op.op1, op.op1, op.op2)
		} else if op.name == "sub" {
			fmt.Fprintf(w, "%s%s = %s - %s\n", prefix, op.op1, op.op1, op.op2)
		}
	}
	fmt.Fprintln(w, `
		END: fmt.Println(h)
	}`)
}

func GoProgram(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)

	ops := []BinaryOp{}

	for scanner.Scan() {
		line := scanner.Text()
		ops = append(ops, parseOp(line))
	}

	p := Program{ops: ops, state: State{regs: map[string]int{"a": 1}}}

	p.toGoProgram(w)
}

// Solve solves the puzzle
func Solve(r io.Reader, part Part) int {
	scanner := bufio.NewScanner(r)

	ops := []BinaryOp{}

	for scanner.Scan() {
		line := scanner.Text()
		ops = append(ops, parseOp(line))
	}

	if part == Part1 {
		p := Program{ops: ops, state: State{regs: map[string]int{}}}

		p.run()

		return p.state.mulCounter
	}

	panic("Unknown part")
}
