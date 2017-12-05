package Day05

import (
	"bufio"
	"io"
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

// Instruction represents instruction
type Instruction struct {
	offset int
}

// Instructions is a slice of instructions
type Instructions []Instruction

func (instructions Instructions) isInside(offset int) bool {
	if offset >= 0 && offset < len(instructions) {
		return true
	}
	return false
}

// OffsetFunc calculates offset
type OffsetFunc func(int) int

func decrementOffsetFunc(offset int) int {
	if offset >= 3 {
		return -1
	}

	return 1
}

func incrementOffsetFunc(offset int) int {
	return 1
}

// Solve solves the puzzle based on part
func Solve(r io.Reader, part Part) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	instructions := Instructions{}
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Bad input")
		}
		instructions = append(instructions, Instruction{offset: n})
	}

	var offsetFunc OffsetFunc
	if part == Part1 {
		offsetFunc = incrementOffsetFunc
	}
	if part == Part2 {
		offsetFunc = decrementOffsetFunc
	}

	currentPosition := 0
	stepsCount := 0

	for {
		currentOffset := instructions[currentPosition].offset
		newPosition := currentPosition + currentOffset

		if !instructions.isInside(newPosition) {
			return stepsCount + 1
		}

		instructions[currentPosition].offset += offsetFunc(currentOffset)

		currentPosition = newPosition
		stepsCount++
	}
}
