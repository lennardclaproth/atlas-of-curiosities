package datastruct

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(d T) {
	s.data = append(s.data, d)
}

func (s *Stack[T]) Pop() T {
	d := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return d
}
