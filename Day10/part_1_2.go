package Day10

import (
	"fmt"
	"io"
)

// Part is either Part1 or Part2
type Part int

const (
	// Part1 is part 1
	Part1 Part = 1
	// Part2 is part 2
	Part2 Part = 2
)

type Node struct {
	value int
	next  *Node
}

func reverseList(list *Node, t int) *Node {
	current := list
	var top *Node
	for {
		if current == nil {
			break
		}
		temp := current.next
		current.next = top
		top = current
		current = temp
	}

	return top
}

func printList(pList *Node) {

	pCurr := pList
	for {
		fmt.Printf("%d ", pCurr.value)

		if pCurr.next != nil {
			pCurr = pCurr.next
		} else {
			break
		}
	}
	fmt.Println("")
}

// Solve solves the puzzle
func Solve(r io.Reader, part Part) int {

	var start, current *Node
	start = nil
	current = nil

	for i := 0; i < 5; i++ {
		if current == nil {
			start = &Node{value: i}
			current = start
		} else {
			node := Node{value: i}
			current.next = &node

			current = current.next
		}
	}

	// current.next = start

	// j := reverseList(start)

	// fmt.Println(j, j.next, j.next.next)

	printList(start)
	printList(reverseList(start, 3))

	// scanner := bufio.NewScanner(r)
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	fmt.Println(line)
	// }

	return -1
}
