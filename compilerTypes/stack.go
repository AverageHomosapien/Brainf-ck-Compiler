package compilerTypes

type Stack struct {
	elements [][]int
}

func (s *Stack) Push(memoryPointer int, codePointer int) {
	s.elements = append(s.elements, []int{memoryPointer, codePointer})
}

func (s *Stack) Pop() (memoryPointer int, codePointer int, elementPopped bool) {
	if len(s.elements) == 0 {
		return 0, 0, false
	}

	elem := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return elem[0], elem[1], true
}

func (s *Stack) Peek() (memoryPointer int, codePointer int, elementPopped bool) {
	if len(s.elements) == 0 {
		return 0, 0, false
	}

	elem := s.elements[len(s.elements)-1]
	return elem[0], elem[1], true
}
