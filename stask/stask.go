package stask

type Stask struct {
	len int
	*node
}

type node struct {
	val  int
	next *node
}

func New() *Stask {
	return &Stask{
		len: 0,
	}
}

func (s *Stask) Push(val int) {
	n := &node{val: val}
	if s.node == nil {
		s.node = n
		s.len++
		return
	}

	n.next = s.node
	s.node = n
	s.len++
}

func (s *Stask) Pop() int {
	if s.IsEmpty() {
		return 0
	}

	val := s.node.val
	s.node = s.node.next
	s.len--
	return val
}

func (s *Stask) IsEmpty() bool {
	return s.len == 0
}
