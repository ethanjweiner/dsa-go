package structures

type Stack struct {
	items []any
}

func (s *Stack) Push(value any) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() (popped any) {
	lastIndex := len(s.items) - 1

	if lastIndex == -1 {
		return nil
	}

	popped = s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return
}
