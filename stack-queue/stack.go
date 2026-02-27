package stackqueue

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(element T) {
	s.data = append(s.data, element)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}

	lastIndex := len(s.data) - 1
	element := s.data[lastIndex]

	var zero T
	s.data[lastIndex] = zero
	s.data = s.data[:lastIndex]

	return element, true
}

func (s *Stack[T]) Peek() (T, bool) {
	var element T
	var exists = false

	if len(s.data) > 0 {
		element = s.data[len(s.data)-1]
		exists = true
	}

	return element, exists
}

func (s *Stack[T]) Len() int {

	return len(s.data)
}

func (s *Stack[T]) IsEmpty() bool {

	return s.Len() == 0
}
