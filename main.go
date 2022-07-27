package main

import (
	"github.com/yokowu/leetcode/tree"
)

func main() {
	t := &tree.NodeTree{
		Left: &tree.NodeTree{
			Val:   10,
			Left:  &tree.NodeTree{Val: 12},
			Right: &tree.NodeTree{Val: 13},
		},
		Right: &tree.NodeTree{
			Val:   11,
			Left:  &tree.NodeTree{Val: 14},
			Right: &tree.NodeTree{Val: 15},
		},
		Val: 5,
	}
	tree.InOrderTraversal(t)
}
