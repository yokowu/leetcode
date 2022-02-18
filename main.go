package main

import (
	"fmt"

	"github.com/yokowu/leetcode/mysort"
)

func main() {
	a := []int{3, 2, 6, 3, 1, 100}
	mysort.MergeSort2(a)
	fmt.Println(a)
}
