package main

import (
	"fmt"

	"github.com/yokowu/leetcode/tree"
)

func main() {
	t := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val:   2,
			Right: &tree.TreeNode{Val: 4},
		},
		Right: &tree.TreeNode{
			Val: 3,
		},
	}
	res := tree.AverageOfLevels(t)
	fmt.Println(res)
}
