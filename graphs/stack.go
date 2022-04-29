package graphs

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) LasItemIndex() int {
	return s.Size() - 1
}
func (s *Stack[T]) Size() int {
	return len(s.data)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack[T]) Push(item T) {
	s.data = append(s.data, item)
}

func (s *Stack[T]) Pop() (T, bool) {

	if len(s.data) == 0 {
		return *new(T), false
	}

	lastItemIndex := s.Size() - 1

	item := s.data[lastItemIndex]

	s.data = s.data[:lastItemIndex]

	return item, true
}
