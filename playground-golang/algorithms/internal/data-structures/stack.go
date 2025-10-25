package datastruct

type Stack[T any] struct {
	data []T
}

// push adds a new item to the stack data structure. It does this by
// appending it to the end of the data of the stack.
func (s *Stack[T]) Push(d T) {
	s.data = append(s.data, d)
}

// Pop gets the last added item of the stack and returns it.
// the stack data structure rests on the LIFO principle (Last
// In First Out)
func (s *Stack[T]) Pop() T {
	d := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return d
}
