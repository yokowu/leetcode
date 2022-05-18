package main

import (
	"fmt"

	"github.com/yokowu/leetcode/linked"
)

func main() {
	list1 := linked.New(1)
	list1 = linked.Append(list1, 3)
	list1 = linked.Append(list1, 2)
	list1 = linked.Append(list1, 2)
	list1 = linked.Append(list1, 3)
	list1 = linked.Append(list1, 2)
	list1 = linked.Append(list1, 2)
	list1 = linked.Append(list1, 2)
	list1 = linked.Append(list1, 7)

	list := linked.NodesBetweenCriticalPoints(list1)
	fmt.Println(list)
}
