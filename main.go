package main

import (
	"fmt"

	"github.com/yokowu/leetcode/mysort"
)

func main() {
	a := []int{6, 100, 50, 3, 10, 3, 2, 6, 3, 1, 100, 3, 6}
	// mysort.MergeSort2(a)
	mysort.Partition2(a, 6)
	fmt.Println(a)
}
