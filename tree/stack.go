package tree

type Stack struct {
	trees []*TreeNode
}

func (s *Stack) Push(n *TreeNode) {
	s.trees = append(s.trees, n)
}

func (s *Stack) Pop() *TreeNode {
	n := s.trees[len(s.trees)-1]
	s.trees = s.trees[:len(s.trees)-1]
	return n
}

func (s *Stack) Peek() *TreeNode {
	return s.trees[len(s.trees)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.trees) == 0
}
