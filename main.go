package main

import (
	"fmt"

	"github.com/yokowu/leetcode/greedy"
)

func main() {
	fmt.Print(
		greedy.EraseOverlapIntervals(
			[][]int{
				{1, 2},
				{1, 2},
				{1, 2},
			},
		),
	)
}
