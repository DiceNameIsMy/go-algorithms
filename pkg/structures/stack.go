package structures

type StackInt struct {
	items []int
}

func (s *StackInt) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
// Returns false if stack is empty
func (s *StackInt) Pop() (int, bool) {
	if s.IsEmpty() {
		return -1, false
	} else {
		lastItemIdx := len(s.items) - 1
		item := s.items[lastItemIdx]

		s.items = s.items[:lastItemIdx]

		return item, true
	}
}

func (s *StackInt) IsEmpty() bool {
	return len(s.items) == 0
}
