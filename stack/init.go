package stack

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Pop() {
	if s.size == 0 {
		return
	}

	s.data = append(s.data[:s.size-1])
	s.size--
	return
}

func (s *Stack) Push(data interface{}) {
	s.data = append(s.data, data)
	s.size++
}

func (s *Stack) Reset() {
	s.data = []interface{}{}
	s.size = 0
}
