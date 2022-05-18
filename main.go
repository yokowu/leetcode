package main

import "github.com/yokowu/leetcode/linked"

func main() {
	list1 := linked.New(1)
	list1 = linked.Append(list1, 2)
	list1 = linked.Append(list1, 3)
	list1 = linked.Append(list1, 4)
	list1 = linked.Append(list1, 5)

	list1 = linked.SwapNodes(list1, 1)
	linked.Show(list1)
}
