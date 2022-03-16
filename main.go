package main

import (
	"fmt"

	"github.com/yokowu/leetcode/linked"
)

func main() {
	node := linked.New(1)
	node = linked.Append(node, 2)
	node = linked.Append(node, 3)
	node = linked.Append(node, 2)
	node = linked.Append(node, 1)

	b := linked.IsPalindrome2(node)
	fmt.Println(b)
}
