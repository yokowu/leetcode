package main

import (
	"fmt"

	"github.com/yokowu/leetcode/linked"
)

func main() {
	node := linked.New(1)
	node = linked.Append(node, 2)
	node = linked.Append(node, 2)
	node = linked.Append(node, 1)

	fmt.Println(linked.IsPalindrome(node))
}
