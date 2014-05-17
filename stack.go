package expr

type nstack struct {
	top  *nelem
	size int
}

type nelem struct {
	val  Node
	next *nelem
}

func (s *nstack) Size() int {
	return s.size
}

func (s *nstack) Push(n Node) {
	s.top = &nelem{n, s.top}
	s.size++
}

func (s *nstack) Pop() (n Node) {
	if s.Size() > 0 {
		n, s.top = s.top.val, s.top.next
		s.size--
		return
	}
	return nil
}
