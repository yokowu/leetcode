package main

import (
	"fmt"

	"github.com/yokowu/leetcode/greedy"
)

func main() {
	fmt.Print(
		greedy.FindMinArrowShots(
			[][]int{
				{10, 16},
				{2, 8},
				{1, 6},
				{7, 12},
			},
		),
	)
}
