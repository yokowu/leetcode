package main

import "github.com/yokowu/leetcode/linked"

func main() {
	node := linked.New(10)
	node = linked.Append(node, 5)
	node = linked.Append(node, 5)
	node = linked.Append(node, 5)
	node = linked.Append(node, 3)
	node = linked.Append(node, 2)
	node = linked.Append(node, 3)

	node = linked.RemoveElements(node, 5)
	linked.Show(node)
}
