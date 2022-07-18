package main

import (
	"fmt"

	"github.com/yokowu/leetcode/greedy"
)

func main() {
	fmt.Print(
		greedy.LemonadeChange(
			[]int{5, 5, 5, 10, 20},
		),
	)
}
