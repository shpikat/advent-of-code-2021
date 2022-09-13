package day23

type Stack []Move

func NewStack() Stack {
	return make([]Move, 0, 100)
}
func (s *Stack) Push(moves []Move) {
	*s = append(*s, moves...)
}

func (s *Stack) MustPop() (next Move) {
	last := len(*s) - 1
	next = (*s)[last]
	*s = (*s)[:last]
	return
}
