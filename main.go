package main

import (
	"fmt"

	"github.com/yokowu/leetcode/mysort"
)

func main() {
	a := []int{3, 2, 6, 10, 2, 4, 29, 34}
	mysort.HeapSort(a)
	fmt.Println(a[3-1])
}
