package Day17

import (
	"container/list"
	"fmt"
)

// Part is either Part1 or Part2
type Part int

const (
	// Part1 is part 1
	Part1 Part = 1
	// Part2 is part 2
	Part2 Part = 2
)

func Insert(p *[]int, index int, value int) {
	*p = append(*p, 0)

	pp := *p

	copy(pp[index+1:], pp[index:])
	pp[index] = value
}

type Node struct {
	value int
	next  *Node
}

func print(list *Node) {
	if list == nil {
		return
	}
	fmt.Println(list.value)
	print(list.next)
}

// Solve solves the puzzle
func Solve(times int, total int, part Part) int {

	if part == Part1 {
		spinlock := list.New()
		position := spinlock.PushBack(0)

		for i := 1; i <= total; i++ {
			for j := 0; j < times; j++ {
				// step
				if position.Next() == nil {
					position = spinlock.Front()
				} else {
					position = position.Next()
				}
			}

			spinlock.InsertAfter(i, position)

			// step
			if position.Next() == nil {
				position = spinlock.Front()
			} else {
				position = position.Next()
			}
		}

		return position.Next().Value.(int)
	}

	if part == Part2 {
		position := 0
		spinlockLength := 1

		afterZeroValue := -1

		for i := 1; i <= total; i++ {
			position = (position + times) % spinlockLength

			spinlockLength++
			if position == 0 {
				afterZeroValue = i
			}

			position++
		}

		return afterZeroValue
	}

	panic("Unknown part")
}
