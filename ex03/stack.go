package stack

type Stack struct {
	stack []int
	sp    int
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) Push(v int) {
	if s.sp >= cap(s.stack)-1 {
		s.stack = append(s.stack, v)
	} else {
		s.stack[s.sp] = v
	}
	s.sp += 1
}

func (s *Stack) Pop() int {
	s.sp -= 1
	return s.stack[s.sp]
}
