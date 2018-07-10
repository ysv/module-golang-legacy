package stack

type Stack struct {
	stack []int
	sp    int
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) Push(v int) {
	s.sp += 1
	s.stack = append(s.stack, v)
}

func (s *Stack) Pop() int {
	s.sp -= 1
	return s.stack[s.sp]
}
