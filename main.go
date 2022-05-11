package main

import "github.com/yokowu/leetcode/linked"

func main() {
	l1 := linked.New(7)
	l1 = linked.Append(l1, 2)
	l1 = linked.Append(l1, 4)
	l1 = linked.Append(l1, 3)
	l2 := linked.New(5)
	l2 = linked.Append(l2, 6)
	l2 = linked.Append(l2, 4)

	l3 := linked.AddTwoNumbers(l1, l2)
	linked.Show(l3)
}
