package main

import (
	"fmt"

	"github.com/yokowu/leetcode/greedy"
)

func main() {
	r := greedy.Merge([][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	})
	fmt.Println(r)
}
