package compiler_types

type Stack struct {
	elements [][]int
}

func (s *Stack) Push(memory_pointer int, code_pointer int) {
	s.elements = append(s.elements, []int{memory_pointer, code_pointer})
}

func (s *Stack) Pop() (memory_pointer int, code_pointer int, element_popped bool) {
	if len(s.elements) == 0 {
		return 0, 0, false
	}

	elem := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return elem[0], elem[1], true
}

func (s *Stack) Peek() (memory_pointer int, code_pointer int, element_popped bool) {
	if len(s.elements) == 0 {
		return 0, 0, false
	}

	elem := s.elements[len(s.elements)-1]
	return elem[0], elem[1], true
}
