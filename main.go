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
	res := tree.LeverlOrderBottom(t)
	fmt.Println(res)
}
