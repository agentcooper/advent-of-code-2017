package Day10

import (
	"fmt"
	"io"
	"io/ioutil"
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

// Grab takes `length` elements starting from `current`, using `arr` as circular slice
func Grab(arr []int, current int, length int) []int {
	out := []int{}
	i := 0
	for i < length {
		out = append(out, arr[(current+i)%len(arr)])
		i++
	}
	return out
}

func Perform(arr []int, current int, length int) {
	i := current

	gr := Grab(arr, current, length)
	for j := range gr {
		arr[i] = gr[len(gr)-1-j]
		i = (i + 1) % len(arr)
	}
}

func Partition(arr []int, partitionSize int, callback func(int, []int)) {
	l := len(arr)

	k := l / partitionSize
	if l%partitionSize > 0 {
		k++
	}

	for i := 0; i < k; i++ {
		start := i * partitionSize
		end := (i + 1) * partitionSize

		if end > l {
			callback(i, arr[start:l])
		} else {
			callback(i, arr[start:end])
		}
	}
}

func Run(lengthSequence []int, times int) []int {
	arr := []int{}
	for i := 0; i < 256; i++ {
		arr = append(arr, i)
	}

	current := 0
	skip := 0

	for k := 0; k < times; k++ {
		for _, length := range lengthSequence {
			Perform(arr, current, length)

			current = (current + length + skip) % len(arr)
			skip++
		}
	}

	return arr
}

// Solve solves the puzzle
func SolvePart1(r io.Reader) int {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	ns, err := utils.StringsToInts(strings.Split(string(b), ","))
	if err != nil {
		panic(err)
	}

	out := Run(ns, 1)
	return out[0] * out[1]
}

func KnotHash(input string) string {
	lengthSequence := []int{}
	for _, r := range input {
		lengthSequence = append(lengthSequence, int(r))
	}
	lengthSequence = append(lengthSequence, []int{17, 31, 73, 47, 23}...)

	out := ""
	Partition(Run(lengthSequence, 64), 16, func(i int, part []int) {
		sum := 0
		for _, p := range part {
			sum = sum ^ p
		}

		out += fmt.Sprintf("%02x", sum)
	})

	return out
}

func SolvePart2(r io.Reader) string {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	return KnotHash(string(b))
}
