package main

import (
	"fmt"

	"github.com/yokowu/leetcode/greedy"
)

func main() {
	i := greedy.FindContentChildren(
		[]int{3, 4, 5},
		[]int{2, 3, 4},
	)

	fmt.Println(i)
}
