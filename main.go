package main

import (
	"fmt"

	"github.com/yokowu/leetcode/greedy"
)

func main() {
	fmt.Print(
		greedy.CanCompleteCircuit(
			[]int{1, 2, 3, 4, 5},
			[]int{3, 4, 5, 1, 2},
		),
	)
}
