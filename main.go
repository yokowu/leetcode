package main

import (
	"fmt"

	"github.com/yokowu/leetcode/tree"
)

func main() {
	t := &tree.TreeNode{
		Left: &tree.TreeNode{
			Val:   2,
			Right: &tree.TreeNode{Val: 4},
		},
		Right: &tree.TreeNode{
			Val: 3,
		},
		Val: 1,
	}
	s := tree.Serialize(t)
	fmt.Println(s)

	t = tree.Deserialize(s)
	tree.PreOrder(t)
}
